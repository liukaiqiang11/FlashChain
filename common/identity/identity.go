package identity

type Signer interface {
	Sign(message []byte) ([]byte, error)
}

type Serializer interface {
	Serialize() ([]byte, error)
}
