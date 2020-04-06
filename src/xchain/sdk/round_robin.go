package sdk

import (
	"math"
	"sync"
)

var orderIndex = 0
var orderMutex sync.Mutex
var peerIndex = 0
var peerMutex sync.Mutex

func PickOrderClient(clients []*OrderClient) *OrderClient {
	if clients == nil || len(clients) == 0 {
		return nil
	}
	orderMutex.Lock()
	defer orderMutex.Unlock()
	if orderIndex == math.MaxInt32 {
		orderIndex = 0
	} else {
		orderIndex++
	}
	return clients[orderIndex%len(clients)]
}

func PickPeerClient(clients []*PeerClient) *PeerClient {
	if clients == nil || len(clients) == 0 {
		return nil
	}
	peerMutex.Lock()
	defer peerMutex.Unlock()
	if peerIndex == math.MaxInt32 {
		peerIndex = 0
	} else {
		peerIndex++
	}
	return clients[peerIndex%len(clients)]
}
