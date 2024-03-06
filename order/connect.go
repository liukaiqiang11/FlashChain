package order

import (
	"flag"

	"fchain/config"
	"fchain/order/internal"
	pb "fchain/proto"

	"github.com/pkg/errors"
)

var orderAddress = flag.String("order.address", config.OrderAddress, "order address")

type OrderChaincodeStream interface {
	Send(*pb.ChaincodeMessage) error
	Recv() (*pb.ChaincodeMessage, error)
}

// OrderStream supports the (original) chaincode-as-client interaction pattern
type OrderStream interface {
	OrderChaincodeStream
	CloseSend() error
}

func userChaincodeStreamGetter() (OrderStream, error) {

	if *orderAddress == "" {
		return nil, errors.New("flag 'order.address' must be set")
	}

	conf, err := internal.LoadConfig()
	if err != nil {
		return nil, err
	}

	conn, err := internal.NewClientConn(*orderAddress, conf)
	if err != nil {
		return nil, err
	}

	return internal.NewRegisterClient(conn)
}

type peerStreamGetter func() (OrderStream, error)

var streamGetter peerStreamGetter

func GetOrderStream() (OrderStream, error) {
	if streamGetter == nil {
		streamGetter = userChaincodeStreamGetter
	}

	stream, err := streamGetter()
	if err != nil {
		return nil, err
	}
	return stream, nil
}
