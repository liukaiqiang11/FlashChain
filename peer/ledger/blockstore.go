package ledger

import (
	"fchain/database"
	pb "fchain/proto"

	log "github.com/corgi-kx/logcustom"
)

type BlockStore struct {
	blockStorageDir string
	fileMgr         *blockfileMgr
}

func NewBlockStore(path, dbName string) (*BlockStore, error) {

	dbProvider := database.NewDBProvider(path, dbName)
	err := dbProvider.DeleteAll()
	if err != nil {
		return nil, err
	}
	fileMgr, err := newBlockfileMgr(path, dbProvider)
	if err != nil {
		return nil, err
	}

	return &BlockStore{path, fileMgr}, nil
}

func (store *BlockStore) AddBlock(block *pb.Block) error {
	result := store.fileMgr.AddBlock(block)

	return result
}

func (store *BlockStore) GetBlockchainInfo() (*pb.BlockchainInfo, error) {
	return store.fileMgr.getBlockchainInfo(), nil
}

func (store *BlockStore) RetrieveBlocks(startNum uint64) (*database.Iterator, error) {
	return store.fileMgr.retrieveBlocks(startNum)
}

func (store *BlockStore) RetrieveBlockByHash(blockHash []byte) (*pb.Block, error) {
	return store.fileMgr.retrieveBlockByHash(blockHash)
}

func (store *BlockStore) RetrieveBlockByNumber(blockNum uint64) (*pb.Block, error) {
	return store.fileMgr.retrieveBlockByNumber(blockNum)
}

func (store *BlockStore) Shutdown() {
	log.Debug("closing fs blockStore")
	store.fileMgr.close()
}
