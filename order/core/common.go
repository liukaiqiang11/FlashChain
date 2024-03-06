package core

import (
	pb "fchain/proto"
)

func DeleteTx(arr []string, id string) []string {
	for i := 0; i < len(arr); i++ {
		if id == arr[i] {
			arr = append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}

func GetEnvelope(txID string, envelopes []*pb.Envelope) (*pb.Envelope, []*pb.Envelope) {
	for index, envelope := range envelopes {
		if envelope.TxID == txID {
			envelopes = append(envelopes[:index], envelopes[index+1:]...)
			return envelope, envelopes
		}
	}
	return nil, envelopes
}
