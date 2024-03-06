package shim

import (
	"context"
	"fchain/config"
	"fchain/peer/shim/internal"
	pb "fchain/proto"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"time"
)

func userChaincodeStreamGetter(address string) (pb.ChaincodeSupportClient, error) {

	if address == "" {
		return nil, errors.New("flag 'peer.address' must be set")
	}

	conf, err := internal.LoadConfig()
	if err != nil {
		return nil, err
	}

	conn, err := internal.NewClientConn(address, conf)
	if err != nil {
		return nil, err
	}

	return internal.NewRegisterClient(conn)
}

type peerStreamGetter func(address string) (pb.ChaincodeSupportClient, error)

var streamGetter peerStreamGetter

func GetPeerClient() ([]pb.ChaincodeSupportClient, error) {
	if streamGetter == nil {
		streamGetter = userChaincodeStreamGetter
	}

	var peerClients []pb.ChaincodeSupportClient

	for _, peerAddress := range config.PeerAddressArr {
		stream, err := streamGetter(peerAddress)
		if err != nil {
			return nil, err
		}
		peerClients = append(peerClients, stream)
	}

	return peerClients, nil
}

func GetPeerClientStream() ([]ClientStream, error) {
	var peerClientStreams []ClientStream

	peerClients, err := GetPeerClient()
	if err != nil {
		return nil, err
	}
	for _, client := range peerClients {
		peerClientStream, err := client.Register(context.Background())
		if err != nil {
			return nil, err
		}
		peerClientStreams = append(peerClientStreams, peerClientStream)
	}

	return peerClientStreams, nil
}

func chatWithPeer(stream PeerChaincodeStream) error {

	handler := newChaincodeHandler(stream)

	type recvMsg struct {
		msg *pb.ChaincodeMessage
		err error
	}
	msgAvail := make(chan *recvMsg, 1)
	errc := make(chan error)

	receiveMessage := func() {
		in, err := stream.Recv()
		msgAvail <- &recvMsg{in, err}
	}

	go receiveMessage()
	for {
		select {
		case rmsg := <-msgAvail:
			switch {
			case rmsg.err == io.EOF:
				time.Sleep(10 * time.Second)
				return errors.New("received EOF, ending chaincode stream")
			case rmsg.err != nil:
				return fmt.Errorf("receive failed: %s", rmsg.err)
			case rmsg.msg == nil:
				return errors.New("received nil message, ending chaincode stream")
			default:
				err := handler.handleMessage(rmsg.msg, errc)
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
