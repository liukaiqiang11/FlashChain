package util

import (
	"bytes"
	"crypto/sha256"
	pb "fchain/proto"
	"fmt"
	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
	"strings"
)

// NewBlock 创建一个新的区块
func NewBlock(seqNum uint64, previousHash []byte) *pb.Block {
	block := &pb.Block{}
	block.Header = &pb.BlockHeader{}
	block.Header.Number = seqNum
	block.Header.Timestamp = timestamppb.Now()
	block.Header.PreviousHash = previousHash
	block.Header.DataHash = []byte{}
	block.Data = &pb.BlockData{}
	block.Metadata = &pb.BlockMetadata{}
	block.Metadata.Transaction_State = []*pb.ValidationCodes{}

	return block
}

func GenesisBlock() *pb.Block {
	block := NewBlock(0, []byte{})
	block.Header.Hash = BlockHeaderHash(block.Header)
	return block
}

func CreateNewBlock(seqNum uint64, previousHash []byte, messages []*pb.Envelopes) *pb.Block {
	var err error
	block := NewBlock(seqNum, previousHash)
	data := &pb.BlockData{
		Data: make([][]byte, len(messages)),
	}
	for i, msg := range messages {
		data.Data[i], err = proto.Marshal(msg)
		if err != nil {
			panic(err)
		}
	}
	block.Data = data
	block.Header.DataHash = BlockDataHash(block.Data)
	block.Header.Hash = BlockHeaderHash(block.Header)
	return block
}

func NewBlockSign(block *pb.Block, signatureHeader *pb.SignatureHeader, signature []byte) *pb.Block {
	block.Metadata.SignatureHeader = signatureHeader
	block.Metadata.Signature = signature
	return block
}

// NewGenesisBlock 生成创世区块
func NewGenesisBlock() *pb.Block {
	genesisBlock := NewBlock(0, []byte{})

	return genesisBlock
}

func BlockBytes(b *pb.Block) []byte {
	result, err := proto.Marshal(b)
	if err != nil {
		panic(err)
	}
	return result
}

// BlockHeaderBytes 获取区块头的字节数组
func BlockHeaderBytes(b *pb.BlockHeader) []byte {
	result, err := proto.Marshal(b)
	if err != nil {
		panic(err)
	}
	return result
}

// BlockMetadataBytes 获取区块元数据的字节数组
func BlockMetadataBytes(b *pb.BlockMetadata) []byte {
	result, err := proto.Marshal(b)
	if err != nil {
		panic(err)
	}
	return result
}

// BlockHeaderHash 计算区块头的hash
func BlockHeaderHash(b *pb.BlockHeader) []byte {
	sum := sha256.Sum256(BlockHeaderBytes(b))
	return sum[:]
}

// BlockDataHash 计算区块体的hash
func BlockDataHash(b *pb.BlockData) []byte {
	sum := sha256.Sum256(bytes.Join(b.Data, nil))
	return sum[:]
}

// GetMetadataFromBlock 从区块中获取 BlockMetadata
func GetMetadataFromBlock(block *pb.Block) (*pb.BlockMetadata, error) {
	if block.Metadata == nil {
		return nil, errors.New("no metadata in block")
	}
	return block.Metadata, nil
}

// GetMetadataFromBlockOrPanic 从区块中获取 BlockMetadata,如果出错则抛出 panic
func GetMetadataFromBlockOrPanic(block *pb.Block) *pb.BlockMetadata {
	md, err := GetMetadataFromBlock(block)
	if err != nil {
		panic(err)
	}
	return md
}

func StringBlockByte(blockByte []byte) {
	var lines []string

	block, err := UnmarshalBlock(blockByte)
	if err != nil {
		log.Fatal(err)
	}

	lines = append(lines, fmt.Sprintf("==================================== block Number is: %d ====================================", block.Header.Number))

	lines = append(lines, fmt.Sprintf("--- block PreviousHash is: %v", block.Header.PreviousHash))

	lines = append(lines, fmt.Sprintf("--- block Hash is: %v", block.Header.Hash))

	lines = append(lines, fmt.Sprintf("--- block DataHash is: %v", block.Header.DataHash))

	lines = append(lines, fmt.Sprintf("--- block create Timestamp is: %s", block.Header.Timestamp))

	for gIdx, data := range block.Data.Data {
		envs, err := UnmarshalEnvelopes(data)
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, "")
		lines = append(lines, fmt.Sprintf("--- Transaction Group is: %d", gIdx))

		for txIdx, envelope := range envs.Envelope {
			lines = append(lines, fmt.Sprintf("======================== Transaction ID is: %s ========================", envelope.TxID))

			ap, err := GetActionPayloadFromEnvelopeMsg(envelope)
			if err != nil {
				log.Fatal(err)
			}
			lines = append(lines, fmt.Sprintf("--- Transaction  Input is: %v", ap.Input))
			rwset, err := GetKVRWSetKwFromEnvelope(envelope)
			if err != nil {
				log.Fatal(err)
			}
			if rwset.Reads != nil {
				lines = append(lines, fmt.Sprintf("--- Transaction  ReadSet is: %v", rwset.Reads))
			}
			if rwset.Writes != nil {
				lines = append(lines, fmt.Sprintf("--- Transaction  WriteSet is: %v", rwset.Writes))
			}

			validateCode := block.Metadata.Transaction_State[gIdx].Transaction_State[txIdx]

			lines = append(lines, fmt.Sprintf("--- Transaction  ValidationCode is: %v", validateCode))

			lines = append(lines, "")
		}
	}

	fmt.Printf("%s\n\n\n\n\n\n", strings.Join(lines, "\n"))

}
