package communication

import (
	."fmt"
	."string"
	."network"
	)
type elevatorNr struct {
	ip string
}

func getMessage() {
	for {
		message := networkChan.Incomming
		go extractMessage(message)
	}
}

func extractMessage(message network.Message) {}
	tempMessage := message.IncommingMessage
	tempMessage = json.Decode(tempMessage)
	if hasPrefix(tempMessage) == "ord" {
		//remove prefix and run order

	}
	else if hasPrefix(tempMessage) == "iah" {
		//if timeLimit has run out -> call chooseMaster
	}
}