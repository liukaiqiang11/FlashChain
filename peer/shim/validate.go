package shim

import (
	"bytes"
	"fchain/common/signer"
	pb "fchain/proto"
	"fchain/proto/util"
	"github.com/pkg/errors"
	"log"
	"sync"
	"sync/atomic"
	"time"
)

var SuccessNum int32

var EndChannel = make(chan bool, 1)

type blockValidationRequest struct {
	blockNum uint64
	envelope *pb.Envelope
	gIdx     int
	tIdx     int
}

func validateSignatureHeader(sHdr *pb.SignatureHeader) error {
	if sHdr == nil {
		return errors.New("nil SignatureHeader provided")
	}

	// ensure that there is a nonce
	if sHdr.Nonce == nil || len(sHdr.Nonce) == 0 {
		return errors.New("invalid nonce specified in the header")
	}

	// ensure that there is a creator
	if sHdr.Creator == nil || len(sHdr.Creator) == 0 {
		return errors.New("invalid creator specified in the header")
	}

	return nil
}

func validateEndorserTransaction(payload *pb.Payload, hdr *pb.SignatureHeader) (*pb.KVRWSet, error) {

	tx, err := util.UnmarshalTransaction(payload.Data)
	if err != nil {
		return nil, errors.New("transaction Unmarshal fail")
	}

	if tx == nil {
		return nil, errors.New("nil transaction")
	}

	prp, err := util.UnmarshalProposalResponsePayload(tx.Payload.ProposalResponsePayload)
	if err != nil {
		return nil, err
	}

	ca, err := util.UnmarshalChaincodeAction(prp.Extension)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting chaincodeAction from envelop")
	}

	csBytes, err := util.GetBytesChaincodeSpec(tx.Payload.Input)
	if err != nil {
		return nil, err
	}
	pHash, err := util.GetProposalHash(hdr, csBytes)
	if err != nil {
		return nil, err
	}
	// ensure that the proposal hash matches
	if !bytes.Equal(pHash, prp.ProposalHash) {
		return nil, errors.New("proposal hash does not match")
	}

	prpBytes, err := util.GetBytesProposalResponsePayload(pHash, ca.Response, ca.Results)
	if err != nil {
		return nil, err
	}

	for _, endorsement := range tx.Payload.Endorsements {

		if err = signer.Verify(endorsement.Endorser, endorsement.Signature, append(prpBytes, endorsement.Endorser...)); err != nil {
			return nil, err
		}
	}

	kVRWSet, err := util.UnmarshalKVRWSet(ca.Results)
	if err != nil {
		return nil, errors.WithMessage(err, "error getting txID kVRWSet ChaincodeAction")
	}

	return kVRWSet, nil
}

func validateRWSet(blockNum uint64, gIdx, seq int, txRWSet *pb.KVRWSet) error {

	for _, write := range txRWSet.Writes {
		detection.Remove(write.Key)
	}

	// 验证读写集，如果发现脏读，则把这个事务中止
	for _, read := range txRWSet.Reads {
		stateRead, _ := State.Get(read.Key)

		if stateRead == nil {
			return errors.New("Failed to read a transaction from the state database")
		}

		if read.Version.TxNum != stateRead.Version.TxNum || read.Version.GroupNum != stateRead.Version.GroupNum ||
			read.Version.BlockNum != stateRead.Version.BlockNum {

			return errors.New("transaction rwSet validate failure")
		}
	}

	for _, write := range txRWSet.Writes {
		if write.IsDelete {
			State.Remove(write.Key)
		} else {
			value := &StateDB{Value: write.Value, Version: &pb.Version{BlockNum: blockNum, GroupNum: uint64(gIdx), TxNum: uint64(seq)}}
			State.Set(write.Key, value)

		}
	}

	atomic.AddInt32(&SuccessNum, 1)

	tl, ok := TxLatency.Get(txRWSet.TxID)
	if ok {
		tl.endTime = uint64(time.Now().UnixMilli())
		TxLatency.Set(txRWSet.TxID, tl)
	}

	return nil
}

func validateTx(req *blockValidationRequest, vCodes *pb.ValidationCodes) {
	// 验证事务，获取事务验证结果
	validateCode := validateTransaction(req)
	vCodes.Transaction_State = append(vCodes.Transaction_State, validateCode)
}

func validateTransaction(req *blockValidationRequest) pb.TxValidationCode {
	env := req.envelope

	if env == nil {
		return pb.TxValidationCode_NIL_ENVELOPE
	}

	// 验证交易负载
	payload, err := util.UnmarshalPayload(env.Payload)
	if err != nil {
		return pb.TxValidationCode_BAD_PAYLOAD
	}

	sHdr := payload.SignatureHeader

	// 验证签名头
	err = validateSignatureHeader(sHdr)
	if err != nil {
		return pb.TxValidationCode_BAD_SIGNATURE_HEADER
	}

	// 验证交易签名
	err = signer.Verify(sHdr.Creator, env.Signature, env.Payload)
	if err != nil {
		return pb.TxValidationCode_BAD_CREATOR_SIGNATURE
	}

	//获取交易ID
	txId, err := util.GetOrComputeTxIDFromEnvelope(env)
	if err != nil {
		return pb.TxValidationCode_BAD_TXID
	}

	// 验证交易ID
	err = util.CheckTxID(
		txId,
		sHdr.Nonce,
		sHdr.Creator)
	if err != nil {
		return pb.TxValidationCode_BAD_TXID
	}

	if payload.Data == nil {
		return pb.TxValidationCode_BAD_PAYLOAD
	}

	//验证交易背书
	rwSet, err := validateEndorserTransaction(payload, sHdr)
	if err != nil {
		return pb.TxValidationCode_INVALID_ENDORSER_TRANSACTION
	}

	//验证交易读写集
	err = validateRWSet(req.blockNum, req.gIdx, req.tIdx, rwSet)
	if err != nil {
		return pb.TxValidationCode_BAD_RWSET
	}

	return pb.TxValidationCode_VALID
}

// validateBlockMataData 验证order签名是否正确
func validateBlockMataData(block *pb.Block) error {
	blockHeaderBytes := util.BlockHeaderBytes(block.Header)
	creator := block.Metadata.SignatureHeader.Creator
	signature := block.Metadata.Signature
	if err := signer.Verify(creator, signature, blockHeaderBytes); err != nil {
		return err
	}
	return nil
}

// validateDataHash 验证区块的交易hash是否被修改
func validateDataHash(block *pb.Block) error {
	dataHash := util.BlockDataHash(block.Data)
	if !bytes.Equal(block.Header.DataHash, dataHash) {
		return errors.New("DataHash is error")
	}
	return nil
}

// Validate 验证区块
func Validate(block *pb.Block) error {

	var err error
	var wg sync.WaitGroup

	// 验证DataHash查看交易是否被修改
	err = validateDataHash(block)
	if err != nil {
		return err
	}

	// 验证MateDate查看order节点签名是否正确
	err = validateBlockMataData(block)
	if err != nil {
		return err
	}

	// 并发的对每个组的事务进行验证
	for gIdx, data := range block.Data.Data {
		wg.Add(1)
		var vCodes = &pb.ValidationCodes{Transaction_State: []pb.TxValidationCode{}}
		go func(gIndex int, data []byte) {
			defer wg.Done()
			envs, err := util.UnmarshalEnvelopes(data)
			if err != nil {
				log.Fatal(err)
			}
			// 验证每个组中的事务
			for txIndex, env := range envs.Envelope {
				validateTx(&blockValidationRequest{
					envelope: env,
					blockNum: block.Header.Number,
					gIdx:     gIndex,
					tIdx:     txIndex,
				}, vCodes)
			}
		}(gIdx, data)
		block.Metadata.Transaction_State = append(block.Metadata.Transaction_State, vCodes)
	}
	wg.Wait()

	return nil
}

// ValidateBlock 验证区块中的交易
func ValidateBlock(msg *pb.ValidateMessage) error {

	if msg.Payload != nil {
		block, err := util.UnmarshalBlock(msg.Payload)
		if err != nil {
			return err
		}

		// 对区块进行验证
		err = Validate(block)
		if err != nil {
			return err
		}

		// 如果区块验证成功，则把该区块写进账本
		err = Blockchain.AddBlock(block)
		if err != nil {
			return err
		}
	}

	if msg.Type == pb.ValidateMessage_Type_VALIDATE_COMPLETED {
		EndChannel <- true
	}

	return nil
}
