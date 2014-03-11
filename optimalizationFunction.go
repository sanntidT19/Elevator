//PSUDOCODE
/*
	Get all directions and runningToCurrentFloor

	-> iterate over all elervators

	if nextExternalListElement is upward realtive to current but not higher than internal
		Elevator to run = X
	else if nextExternalListElement is downward realtive to current but not lower 
		Elevator to run = X
	
	if X != nil
		assign task to relevant elevator


	go to orderd floor


	stop at internal orders if there are some.

	recive(hopefully) the confirmation on floor reached.

	use priority at internal orders if there exists internal and external
	
	gives output: elevator N -> go to floor X


	/////QUESTIONS//////
	- DIEMMA: should the slave or the master take care of the internal list(?);
	1. if the internal list is in the slave i think it does not correspond to our level of abstaction
	2. the master should not be concerned with the internal list, creates a lot of uneccessary sending.

	- is it possible to just lay the orderqueue on a chan who is used only by the master and when the master dies, 
	the next master takes over based on the existing informatin on that chan?
	
*/

//needs access to an array containing the elevatorsExternalOrders and their internal lists
func OptimalizationAlgorithm() string { //incoming External is an int
	var assignToElevatorNr int
	assignToElevatorNr := N_ELEV+1

	var incomingExternalOrder int 
	var externalOrders [][]string 
	var internalOrders []string
	var currentFloor []int
	var dir []int
	


	for i:=0; i< N_ELEV; i++ {
		switch dir[i] {
			case 1://going up 

				if externalOrders[i][0] == 1 || incomingExternalOrder > currentFloor[i] {//external ordered upwards
					//assignOrder to elevator nr i
					return
				} 

			case -1://going down
				if externalOrders[i][0] == -1 || incomingExternalOrder < currentFloor[i] {//external ordered downwards
					//assignOrder to elevator nr i
					return
				}
			case 0://idle
				for j:=0; j<N_FLOORS;j++ {
					checkForOrders := false
					if externalOrders[j] == 1 {
						checkForOrders = true
						break
					}
				////assignOrder to elevator nr i
					return
				}

		}


	}
	//assignOrder to elevator N_ELEV-1 
}	


func reciveFloorReached(executed <- chan, elevator int) {

}