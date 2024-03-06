package signer

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"math/big"
	"os"
	"strings"

	pb "fchain/proto"
	"fchain/proto/util"

	"github.com/pkg/errors"
)

// Config 用户签名的配置（里面包含MSPID 用户属于那个组织，IdentityPath 用户证书的地址，KeyPath 用户密钥的地址）
type Config struct {
	MSPID        string
	IdentityPath string
	KeyPath      string
}

// Signer 用户的签名，包含公钥、私钥和用户的身份
type Signer struct {
	privateKey   *ecdsa.PrivateKey
	publicKeyKey *ecdsa.PublicKey
	creator      []byte
}

type ECDSASignature struct {
	R, S *big.Int
}

// NewSigner 根据配置创建一个用户签名
func NewSigner(conf *Config) (*Signer, error) {
	sId, err := serializeIdentity(conf.IdentityPath, conf.MSPID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	privateKey, err := loadPrivateKey(conf.KeyPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	publicKey, err := loadPublicKey(conf.IdentityPath)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return &Signer{
		creator:      sId,
		privateKey:   privateKey,
		publicKeyKey: publicKey,
	}, nil
}

func serializeIdentity(clientCert string, mspID string) ([]byte, error) {
	b, err := os.ReadFile(clientCert)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if err := validateEnrollmentCertificate(b); err != nil {
		return nil, err
	}

	bl, _ := pem.Decode(b)
	if bl == nil {
		return nil, err
	}

	key, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return nil, err
	}

	publicKeyBytes, err := x509.MarshalPKIXPublicKey(key.PublicKey)
	if err != nil {
		return nil, err
	}

	sId := &pb.Creator{
		Mspid:   mspID,
		IdBytes: publicKeyBytes,
	}
	return util.MarshalOrPanic(sId), nil
}

func validateEnrollmentCertificate(b []byte) error {
	bl, _ := pem.Decode(b)
	if bl == nil {
		return errors.Errorf("enrollment certificate isn't a valid PEM block")
	}

	if bl.Type != "CERTIFICATE" {
		return errors.Errorf("enrollment certificate should be a certificate, got a %s instead", strings.ToLower(bl.Type))
	}

	if _, err := x509.ParseCertificate(bl.Bytes); err != nil {
		return errors.Errorf("enrollment certificate is not a valid x509 certificate: %v", err)
	}
	return nil
}

func (si *Signer) Serialize() ([]byte, error) {
	return si.creator, nil
}

func (si *Signer) Sign(msg []byte) ([]byte, error) {
	digest := sha256.Sum256(msg)
	return signECDSA(si.privateKey, digest[:])
}

func (si *Signer) Verify(signature, msg []byte) bool {
	digest := sha256.Sum256(msg)
	return verifyECDSA(si.publicKeyKey, signature, digest[:])
}

// Verify 验证签名是否正确
func Verify(creatorBytes []byte, signature, msg []byte) error {
	creator, err := util.UnmarshalCreator(creatorBytes)
	if err != nil {
		return err
	}

	publicKey, err := x509.ParsePKIXPublicKey(creator.IdBytes)
	if err != nil {
		return err
	}
	publicKeyKey := publicKey.(*ecdsa.PublicKey)
	digest := sha256.Sum256(msg)
	if !verifyECDSA(publicKeyKey, signature, digest[:]) {
		return errors.New("verifySignature error")
	}

	return nil
}

// 根据文件加载用户的私钥
func loadPrivateKey(file string) (*ecdsa.PrivateKey, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	bl, _ := pem.Decode(b)
	if bl == nil {
		return nil, errors.Errorf("failed to decode PEM block from %s", file)
	}
	key, err := parsePrivateKey(bl.Bytes)
	if err != nil {
		return nil, err
	}
	return key.(*ecdsa.PrivateKey), nil
}

// 根据文件加载用户的公钥
func loadPublicKey(file string) (*ecdsa.PublicKey, error) {
	b, err := os.ReadFile(file)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	bl, _ := pem.Decode(b)
	if bl == nil {
		return nil, errors.Errorf("failed to decode PEM block from %s", file)
	}
	key, err := x509.ParseCertificate(bl.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey := key.PublicKey
	return publicKey.(*ecdsa.PublicKey), nil
}

func parsePrivateKey(der []byte) (crypto.PrivateKey, error) {
	// OpenSSL 1.0.0 generates PKCS#8 keys.
	if key, err := x509.ParsePKCS8PrivateKey(der); err == nil {
		switch key := key.(type) {
		// Fabric only supports ECDSA at the moment.
		case *ecdsa.PrivateKey:
			return key, nil
		default:
			return nil, errors.Errorf("found unknown private key type (%T) in PKCS#8 wrapping", key)
		}
	}

	// OpenSSL ecparam generates SEC1 EC private keys for ECDSA.
	key, err := x509.ParseECPrivateKey(der)
	if err != nil {
		return nil, errors.Errorf("failed to parse private key: %v", err)
	}

	return key, nil
}

func signECDSA(k *ecdsa.PrivateKey, digest []byte) (signature []byte, err error) {
	r, s, err := ecdsa.Sign(rand.Reader, k, digest)
	if err != nil {
		return nil, err
	}

	return asn1.Marshal(ECDSASignature{R: r, S: s})
}

func verifyECDSA(pubKey *ecdsa.PublicKey, signature, digest []byte) bool {
	var sig ECDSASignature
	_, err := asn1.Unmarshal(signature, &sig)
	if err != nil {
		return false
	}

	return ecdsa.Verify(pubKey, digest, sig.R, sig.S)
}
