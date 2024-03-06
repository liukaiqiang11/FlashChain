package database

import (
	log "github.com/corgi-kx/logcustom"
	"github.com/pkg/errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/iterator"
	goleveldbutil "github.com/syndtr/goleveldb/leveldb/util"
)

const maxBatchSize = 1000000

// DBProvider is a core to a named db
type DBProvider struct {
	dbName string
	db     *DB
}

// Iterator extends actual leveldb iterator
type Iterator struct {
	dbName string
	iterator.Iterator
}

// NewDBProvider constructs a `DBProvider`
func NewDBProvider(dp, dn string) *DBProvider {
	db := CreateDB(dp)
	db.Open()

	return &DBProvider{
		dbName: dn,
		db:     db,
	}
}

// Get returns the value for the given key
func (p *DBProvider) Get(key []byte) ([]byte, error) {
	return p.db.Get(constructLevelKey(p.dbName, key))
}

// Put saves the key/value
func (p *DBProvider) Put(key []byte, value []byte, sync bool) error {
	return p.db.Put(constructLevelKey(p.dbName, key), value, sync)
}

// Delete deletes the given key
func (p *DBProvider) Delete(key []byte, sync bool) error {
	return p.db.Delete(constructLevelKey(p.dbName, key), sync)
}

// DeleteAll deletes all the keys that belong to the channel (dbName).
func (p *DBProvider) DeleteAll() error {
	iter, err := p.GetIterator(nil, nil)
	if err != nil {
		return err
	}
	defer iter.Release()

	// use leveldb iterator directly to be more efficient
	dbIter := iter.Iterator

	// This is common code shared by all the leveldb instances. Because each leveldb has its own key size pattern,
	// each batch is limited by memory usage instead of number of keys. Once the batch memory usage reaches maxBatchSize,
	// the batch will be committed.
	numKeys := 0
	batchSize := 0
	batch := &leveldb.Batch{}
	for dbIter.Next() {
		if err := dbIter.Error(); err != nil {
			return errors.Wrap(err, "internal leveldb error while retrieving data from db iterator")
		}
		key := dbIter.Key()
		numKeys++
		batchSize = batchSize + len(key)
		batch.Delete(key)
		if batchSize >= maxBatchSize {
			if err := p.db.WriteBatch(batch, true); err != nil {
				return err
			}
			log.Infof("Have removed %d entries for channel %s in leveldb %s", numKeys, p.dbName, p.db.dbPath)
			batchSize = 0
			batch.Reset()
		}
	}
	if batch.Len() > 0 {
		return p.db.WriteBatch(batch, true)
	}
	return nil
}

// IsEmpty returns true if no data exists for the DBProvider
func (p *DBProvider) IsEmpty() (bool, error) {
	itr, err := p.GetIterator(nil, nil)
	if err != nil {
		return false, err
	}
	defer itr.Release()

	if err := itr.Error(); err != nil {
		return false, errors.WithMessagef(itr.Error(), "internal leveldb error while obtaining next entry from iterator")
	}

	return !itr.Next(), nil
}

// NewUpdateBatch returns a new UpdateBatch that can be used to update the db
func (p *DBProvider) NewUpdateBatch() *UpdateBatch {
	return &UpdateBatch{
		dbName:       p.dbName,
		leveldbBatch: &leveldb.Batch{},
	}
}

// WriteBatch writes a batch in an atomic way
func (p *DBProvider) WriteBatch(batch *UpdateBatch, sync bool) error {
	if batch == nil || batch.leveldbBatch.Len() == 0 {
		return nil
	}
	if err := p.db.WriteBatch(batch.leveldbBatch, sync); err != nil {
		return err
	}
	return nil
}

// GetIterator gets a core to iterator. The iterator should be released after the use.
// The resultset contains all the keys that are present in the db between the startKey (inclusive) and the endKey (exclusive).
// A nil startKey represents the first available key and a nil endKey represent a logical key after the last available key
func (p *DBProvider) GetIterator(startKey []byte, endKey []byte) (*Iterator, error) {
	var itr iterator.Iterator
	if startKey == nil && endKey == nil {
		itr = p.db.GetIterator(startKey, startKey)
	} else if endKey == nil {
		sKey := constructLevelKey(p.dbName, startKey)
		itr = p.db.GetIterator(sKey, startKey)
	} else {
		sKey := constructLevelKey(p.dbName, startKey)
		eKey := constructLevelKey(p.dbName, endKey)
		itr = p.db.GetIterator(sKey, eKey)
	}
	if err := itr.Error(); err != nil {
		itr.Release()
		return nil, errors.Wrapf(err, "internal leveldb error while obtaining db iterator")
	}
	return &Iterator{p.dbName, itr}, nil
}

func (p *DBProvider) PrefixQuery(key []byte) (*Iterator, error) {
	k := constructLevelKey(p.dbName, key)
	//log.Debugf("Getting iterator for range [%#v] - [%#v]", sKey, eKey)
	itr := p.db.PrefixQuery(k)
	if err := itr.Error(); err != nil {
		itr.Release()
		return nil, errors.Wrapf(err, "internal leveldb error while obtaining db iterator")
	}
	return &Iterator{p.dbName, itr}, nil
}
func (p *DBProvider) GetSnapshot() (*Snapshot, error) {
	itr, err := p.db.db.GetSnapshot()
	if err != nil {
		return nil, errors.Wrap(err, "error writing batch to leveldb")
	}

	return &Snapshot{
		dbName:   p.dbName,
		snapshot: itr,
	}, nil
}

func (sn *Snapshot) Get(key []byte) ([]byte, error) {
	k := constructLevelKey(sn.dbName, key)
	value, err := sn.snapshot.Get(k, sn.readOpts)
	if err != nil {
		return nil, errors.Wrap(err, "error writing batch to leveldb")
	}
	return value, nil
}
func (sn *Snapshot) Release() {
	sn.snapshot.Release()
}

func (sn *Snapshot) NewIterator(startKey []byte, endKey []byte) (*Iterator, error) {
	var itr iterator.Iterator
	if startKey == nil && endKey == nil {
		itr = sn.snapshot.NewIterator(&goleveldbutil.Range{Start: startKey, Limit: endKey}, sn.readOpts)
	} else {
		sKey := constructLevelKey(sn.dbName, startKey)
		eKey := constructLevelKey(sn.dbName, endKey)
		itr = sn.snapshot.NewIterator(&goleveldbutil.Range{Start: sKey, Limit: eKey}, sn.readOpts)
	}

	if err := itr.Error(); err != nil {
		itr.Release()
		return nil, errors.Wrapf(err, "internal leveldb error while obtaining db iterator")
	}

	return &Iterator{sn.dbName, itr}, nil
}

// Close closes the DBProvider after its db data have been deleted
func (p *DBProvider) Close() {
	p.db.Close()
}

// Put adds a KV
func (b *UpdateBatch) Put(key []byte, value []byte) {
	if value == nil {
		panic("Nil value not allowed")
	}
	k := constructLevelKey(b.dbName, key)
	b.leveldbBatch.Put(k, value)
	b.size += len(k) + len(value)
}

// Delete deletes a Key and associated value
func (b *UpdateBatch) Delete(key []byte) {
	k := constructLevelKey(b.dbName, key)
	b.size += len(k)
	b.leveldbBatch.Delete(k)
}

// Size returns the current size of the batch
func (b *UpdateBatch) Size() int {
	return b.size
}

// Len returns number of records in the batch
func (b *UpdateBatch) Len() int {
	return b.leveldbBatch.Len()
}

// Reset resets the batch
func (b *UpdateBatch) Reset() {
	b.leveldbBatch.Reset()
	b.size = 0
}

// Seek moves the iterator to the first key/value pair
// whose key is greater than or equal to the given key.
// It returns whether such pair exist.
func (itr *Iterator) Seek(key []byte) bool {
	levelKey := constructLevelKey(itr.dbName, key)
	return itr.Iterator.Seek(levelKey)
}

func constructLevelKey(dbName string, key []byte) []byte {
	return append([]byte(dbName), key...)
}
