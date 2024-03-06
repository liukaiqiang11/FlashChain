package order

import (
	"fmt"
	"io"

	"fchain/order/internal"
	pb "fchain/proto"

	"github.com/pkg/errors"
)

type OrderServer struct {
	Address string
	pb.UnimplementedChaincodeSupportServer
}

type recvMsg struct {
	msg *pb.ChaincodeMessage
	err error
}

func NewOrderServer(addr string) *OrderServer {

	return &OrderServer{Address: addr}
}

func (os *OrderServer) Register(ServerStream pb.ChaincodeSupport_RegisterServer) error {

	hdl := NewOrderHandler(ServerStream)

	msgAvail := make(chan *recvMsg, 1)
	errc := make(chan error)

	receiveMessage := func() {
		in, err := ServerStream.Recv()
		msgAvail <- &recvMsg{in, err}
	}

	go receiveMessage()

	for {
		select {
		case rmsg := <-msgAvail:
			switch {
			case rmsg.err == io.EOF:
				return errors.New("received EOF, ending chaincode stream")
			case rmsg.err != nil:
				err := fmt.Errorf("receive failed: %s", rmsg.err)
				return err
			case rmsg.msg == nil:
				err := errors.New("received nil message, ending chaincode stream")
				return err
			default:
				err := hdl.HandleMessage(rmsg.msg, errc)
				if err != nil {
					err = fmt.Errorf("error handling message: %s", err)
					return err
				}

				go receiveMessage()
			}

		case sendErr := <-errc:
			if sendErr != nil {
				err := fmt.Errorf("error sending: %s", sendErr)
				return err
			}
		}

	}
}

func (os *OrderServer) Start() error {

	if os.Address == "" {
		return errors.New("address must be specified")
	}

	var err error

	tlsCfg, err := internal.LoadTLSConfig()
	if err != nil {
		return err
	}

	server, err := internal.NewServer(os.Address, tlsCfg)
	if err != nil {
		return err
	}

	defer server.Stop()

	pb.RegisterChaincodeSupportServer(server.Server, os)

	return server.Start()
}
