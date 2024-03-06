package peer

import (
	"fchain/config"
	"fchain/peer/shim"
	log "github.com/corgi-kx/logcustom"
)

func Start() {

	cs := shim.NewChaincodeServer(config.PeerAddress)
	err := cs.Start()
	if err != nil {
		log.Debugf("Peer节点开启失败！", err)
	}

}
