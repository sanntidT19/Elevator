package workingChannels

import (
	"fmt"

)
type CommunicationChannels struct {
	order chan string //sends order og i am here
	ioCommand chan int //give floor
	floorReached chan bool

}

var communicationChannels CommunicationChannels
var elevatorChannels elevatorChannels
var networkChannels network.networkChannels

func initComunicationChannels() {
	communicationChannels.order = make(chan string)
	communicationChannels.ioCommand = make(chan int)
	communicationChannels.floorReached = make(chan bool)
}

func initChannels() {
	initComunicationChannels()
}