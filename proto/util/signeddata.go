package util

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"strings"

	pb "fchain/proto"

	"google.golang.org/protobuf/proto"
)

type SignedData struct {
	Data      []byte
	Identity  []byte
	Signature []byte
}

func EnvelopeAsSignedData(env *pb.Envelope) ([]*SignedData, error) {
	if env == nil {
		return nil, fmt.Errorf("No signatures for nil Envelope")
	}

	payload := &pb.Payload{}
	err := proto.Unmarshal(env.Payload, payload)
	if err != nil {
		return nil, err
	}

	if payload.SignatureHeader == nil {
		return nil, fmt.Errorf("Missing SignatureHeader")
	}

	shdr := payload.SignatureHeader

	return []*SignedData{{
		Data:      env.Payload,
		Identity:  shdr.Creator,
		Signature: env.Signature,
	}}, nil
}

func LogMessageForSerializedIdentity(serializedIdentity []byte) string {
	id := &pb.Creator{}
	err := proto.Unmarshal(serializedIdentity, id)
	if err != nil {
		return fmt.Sprintf("Could not unmarshal serialized identity: %s", err)
	}
	pemBlock, _ := pem.Decode(id.IdBytes)
	if pemBlock == nil {
		// not all identities are certificates so simply log the serialized
		// identity bytes
		return fmt.Sprintf("serialized-identity=%x", serializedIdentity)
	}
	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return fmt.Sprintf("Could not parse certificate: %s", err)
	}
	return fmt.Sprintf("(mspid=%s subject=%s issuer=%s serialnumber=%d)", id.Mspid, cert.Subject, cert.Issuer, cert.SerialNumber)
}

func LogMessageForSerializedIdentities(signedData []*SignedData) (logMsg string) {
	var identityMessages []string
	for _, sd := range signedData {
		identityMessages = append(identityMessages, LogMessageForSerializedIdentity(sd.Identity))
	}
	return strings.Join(identityMessages, ", ")
}
