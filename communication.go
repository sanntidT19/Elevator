package communication

import (
	."fmt"
	."string"
	."network"
	)
type elevatorNr struct {
	ip string
}

func getMessage(networkChan <- chan []byte]) { //gets chan from network module
	for {
		message := <-networkChan
		extractMessage(message)
	}
}
func packMessage(outgoingChan chan string, networkChan chan []byte) {
	tempMessage := <- outgoingChan
	codedMessage := json.Marshall(tempMessage)
	networkChan <- codedMessage
}


func extractMessage(messageB []byte, incomingChan chan string) {}

	decodedMessage := json.Decode(messageB)
	if HasPrefix(decodedMessage) == "ord" {
		decodedMessage = TrimPrefix(decodedMessage, "ord") //remove prefix and run order
	}
	else if HasPrefix(decodedMessage) == "iah" { //if timeLimit has run out -> call chooseMaster
		decodedMessage = TrimPrefix(decodedMessage, "iah")

	}
	incomingChan <- decodedMessage
}