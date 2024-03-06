package order

import (
	"log"

	"fchain/config"
)

func Start() {

	cs := NewOrderServer(config.OrderAddress)
	err := cs.Start()
	if err != nil {
		log.Println(err)
	}
}
