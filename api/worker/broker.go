package worker

import (
	"auction/api/utils"
	e "auction/pkg/entity"
)

//Broker somewhat mimics a message broker by treating go channels as queues
//Every we need a new queue update the Broker Struct
//And then change the instantiated broker variable
//This requires restarting the server
//"on" flag suggest that broker is up and running

type Broker struct {
	BidQueue chan *e.BidChannelMessage
	on       bool
}

//Intantiated Broker, Its reference will move around everywhere
var broker = Broker{
	BidQueue: make(chan *e.BidChannelMessage),
}

//start: It starts th broker and run the queue consumers as goroutines
func (broker *Broker) start(services *utils.Services) {
	broker.on = true
	for {
		select {
		case bidMessage := <-broker.BidQueue:
			consumeBids(bidMessage, services)
		}
	}
}

func GetOrCreateBroker(services *utils.Services) *Broker {
	if broker.on {
		return &broker
	}
	go broker.start(services)
	return &broker
}
