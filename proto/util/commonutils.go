package util

import (
	"crypto/rand"
	"fmt"

	"fchain/common/identity"
	pb "fchain/proto"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"
)

// MarshalOrPanic 序列化一个proto消息，如果错误则抛出panic
func MarshalOrPanic(pb proto.Message) []byte {
	data, err := proto.Marshal(pb)
	if err != nil {
		panic(err)
	}
	return data
}

// Marshal 序列化一个proto消息
func Marshal(pb proto.Message) ([]byte, error) {
	return proto.Marshal(pb)
}

// CreateNonceOrPanic 创建一个nonce，如果错误则抛出panic
func CreateNonceOrPanic() []byte {
	nonce, err := CreateNonce()
	if err != nil {
		panic(err)
	}
	return nonce
}

// CreateNonce 创建一个nonce
func CreateNonce() ([]byte, error) {
	nonce, err := getRandomNonce()
	return nonce, errors.WithMessage(err, "error generating random nonce")
}

// getRandomNonce 获取一个随机的nonce
func getRandomNonce() ([]byte, error) {
	key := make([]byte, 24)

	_, err := rand.Read(key)
	if err != nil {
		return nil, errors.Wrap(err, "error getting random bytes")
	}
	return key, nil
}

// MakeSignatureHeader 创建一个SignatureHeader
func MakeSignatureHeader(serializedCreatorCertChain []byte, nonce []byte) *pb.SignatureHeader {
	return &pb.SignatureHeader{
		Creator: serializedCreatorCertChain,
		Nonce:   nonce,
	}
}

// NewSignatureHeader 创建一个SignatureHeader
func NewSignatureHeader(id identity.Serializer) (*pb.SignatureHeader, error) {
	creator, err := id.Serialize()
	if err != nil {
		return nil, err
	}
	nonce, err := CreateNonce()
	if err != nil {
		return nil, err
	}

	return &pb.SignatureHeader{
		Creator: creator,
		Nonce:   nonce,
	}, nil
}

// NewSignatureHeaderOrPanic 创建一个SignatureHeader，如果错误则抛出panic
func NewSignatureHeaderOrPanic(id identity.Serializer) *pb.SignatureHeader {
	if id == nil {
		panic(errors.New("invalid signer. cannot be nil"))
	}

	signatureHeader, err := NewSignatureHeader(id)
	if err != nil {
		panic(fmt.Errorf("failed generating a new SignatureHeader: %s", err))
	}

	return signatureHeader
}

// SignOrPanic 给消息签名，如果错误则抛出panic
func SignOrPanic(signer identity.Signer, msg []byte) []byte {
	if signer == nil {
		panic(errors.New("invalid signer. cannot be nil"))
	}

	sigma, err := signer.Sign(msg)
	if err != nil {
		panic(fmt.Errorf("failed generating signature: %s", err))
	}
	return sigma
}
