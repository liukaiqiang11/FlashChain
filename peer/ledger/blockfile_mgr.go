package ledger

import (
	"bytes"
	"math"
	"sync"
	"sync/atomic"

	"fchain/database"
	pb "fchain/proto"
	"fchain/proto/util"

	log "github.com/corgi-kx/logcustom"
	"github.com/pkg/errors"
)

const (
	blockNumIdxKeyPrefix        = 'n'
	blockHashIdxKeyPrefix       = 'h'
	txIDIdxKeyPrefix            = 't'
	blockNumTranNumIdxKeyPrefix = 'a'
)

type blockfileMgr struct {
	blockStorageDir           string
	db                        *database.DBProvider
	bootstrappingSnapshotInfo *pb.LastBlockInfo
	blkMutex                  *sync.Mutex
	bcInfo                    atomic.Value
}

func newBlockfileMgr(path string, indexStore *database.DBProvider) (*blockfileMgr, error) {

	mgr := &blockfileMgr{blockStorageDir: path, db: indexStore}

	bcInfo := &pb.BlockchainInfo{}

	if mgr.bootstrappingSnapshotInfo != nil {
		bcInfo.Height = mgr.bootstrappingSnapshotInfo.LastBlockNum + 1
		bcInfo.CurrentBlockHash = mgr.bootstrappingSnapshotInfo.LastBlockHash
		bcInfo.PreviousBlockHash = mgr.bootstrappingSnapshotInfo.PreviousBlockHash
	}

	mgr.bcInfo.Store(bcInfo)

	return mgr, nil
}

func (mgr *blockfileMgr) getBlockchainInfo() *pb.BlockchainInfo {
	return mgr.bcInfo.Load().(*pb.BlockchainInfo)
}

func (mgr *blockfileMgr) updateBlockchainInfo(latestBlockHash []byte, latestBlock *pb.Block) {
	currentBCInfo := mgr.getBlockchainInfo()
	newBCInfo := &pb.BlockchainInfo{
		Height:            currentBCInfo.Height + 1,
		CurrentBlockHash:  latestBlockHash,
		PreviousBlockHash: latestBlock.Header.PreviousHash,
	}

	mgr.bcInfo.Store(newBCInfo)
}

func (mgr *blockfileMgr) AddBlock(block *pb.Block) error {
	bcInfo := mgr.getBlockchainInfo()
	if block.Header.Number != bcInfo.Height {
		return errors.Errorf(
			"block number should have been %d but was %d",
			mgr.getBlockchainInfo().Height, block.Header.Number,
		)
	}

	if !bytes.Equal(block.Header.PreviousHash, bcInfo.CurrentBlockHash) {
		return errors.Errorf(
			"unexpected Previous block hash. Expected PreviousHash = [%x], PreviousHash referred in the latest block= [%x]",
			bcInfo.CurrentBlockHash, block.Header.PreviousHash,
		)
	}

	blockBytes := util.BlockBytes(block)
	blockNumBytes := constructBlockNumKey(block.Header.Number)
	if err := mgr.db.Put(blockNumBytes, blockBytes, true); err != nil {
		return errors.WithMessage(err, "error saving block to db")
	}
	blockHashBytes := constructBlockHashKey(block.Header.Hash)
	if err := mgr.db.Put(blockHashBytes, blockNumBytes, true); err != nil {
		return errors.WithMessage(err, "error saving block to db")
	}

	blockHash := block.Header.Hash

	mgr.updateBlockchainInfo(blockHash, block)
	return nil
}

func (mgr *blockfileMgr) retrieveBlockByHash(blockHash []byte) (*pb.Block, error) {
	log.Debugf("retrieveBlockByHash() - blockHash = [%#v]", blockHash)

	return mgr.getBlockByBlockHash(blockHash)
}

func (mgr *blockfileMgr) retrieveBlockByNumber(blockNum uint64) (*pb.Block, error) {

	log.Debugf("retrieveBlockByNumber() - blockNum = [%d]", blockNum)

	// interpret math.MaxUint64 as a request for last block
	if blockNum == math.MaxUint64 {
		blockNum = mgr.getBlockchainInfo().Height - 1
	}
	if blockNum < mgr.firstPossibleBlockNumberInBlockFiles() {
		return nil, errors.Errorf(
			"cannot serve block [%d]. The ledger is bootstrapped from a snapshot. First available block = [%d]",
			blockNum, mgr.firstPossibleBlockNumberInBlockFiles(),
		)
	}

	return mgr.getBlockByBlockNum(blockNum)
}

func (mgr *blockfileMgr) getBlockByBlockNum(blockNum uint64) (*pb.Block, error) {
	blockNumHash := constructBlockNumKey(blockNum)
	blockByte, err := mgr.db.Get(blockNumHash)
	if err != nil {
		return nil, err
	}

	return util.UnmarshalBlock(blockByte)
}

func (mgr *blockfileMgr) getBlockByBlockHash(blockHash []byte) (*pb.Block, error) {
	blockHashKey := constructBlockHashKey(blockHash)
	blockNumHash, err := mgr.db.Get(blockHashKey)
	if err != nil {
		return nil, err
	}
	blockByte, err := mgr.db.Get(blockNumHash)
	if err != nil {
		return nil, err
	}
	return util.UnmarshalBlock(blockByte)
}

func (mgr *blockfileMgr) firstPossibleBlockNumberInBlockFiles() uint64 {
	if mgr.bootstrappingSnapshotInfo == nil {
		return 0
	}
	return mgr.bootstrappingSnapshotInfo.LastBlockNum + 1
}

func (mgr *blockfileMgr) retrieveBlocks(startNum uint64) (*database.Iterator, error) {
	if startNum < mgr.firstPossibleBlockNumberInBlockFiles() {
		return nil, errors.Errorf(
			"cannot serve block [%d]. The ledger is bootstrapped from a snapshot. First available block = [%d]",
			startNum, mgr.firstPossibleBlockNumberInBlockFiles(),
		)
	}
	sNum := constructBlockNumKey(startNum)
	return mgr.db.GetIterator(sNum, nil)
}

func constructBlockNumKey(blockNum uint64) []byte {
	blkNumBytes := EncodeUint64(blockNum)
	return append([]byte{blockNumIdxKeyPrefix}, blkNumBytes...)
}

func constructBlockHashKey(blockHash []byte) []byte {
	return append([]byte{blockHashIdxKeyPrefix}, blockHash...)
}

func (mgr *blockfileMgr) close() {
	mgr.db.Close()
}
