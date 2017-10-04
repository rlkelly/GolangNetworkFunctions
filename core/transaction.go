package core


type Transaction struct {
	Header    TransactionHeader
	Signature []byte
	Payload   []byte
}

type TransactionHeader struct {
	From          []byte
	To            []byte
	Timestamp     uint32
	PayloadHash   []byte
	PayloadLength uint32
	Nonce         uint32
}
