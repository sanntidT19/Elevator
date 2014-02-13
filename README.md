Elevator
========
***Press edit so see formating***
-Our main elevator repository

-Please note conventions here:


-Questions for Student assistant:


All network packages is sendt by the form of a struct. Is this a good idea?



**MODULES**
*Network 
	passes information over the network
	(maybe put selfcheck into this?)

*Master Communication:
	recives external input:
	-to slave Com:
		total order list

*Slave Communication:
	-to master Com:
		external list	- to update total list
		internal list 	- to calculate cost functon
		position 		
		direction

	-needs to manupulate working list when floor reached

*Slave Work(IO unit):
	-to IO to slave Communication
	
	-Utilizes the elev.go to operate the elevator

*Setup Module
	-choose who is master and slaves.
	-obtain ip adresses, and numbering of slaves.

*Selfcheck Module
	-Working with master and slaves, with detecting change in number of elevators.
	(if no message recived in 0.5 sec assume connecton is lost. Assumed 10 messages per secound is sent).
	
	
**What exceptions can happen and how do we handle it?**

- If one elevator with orders loses connection to the others:
	- Need to assume that the elevator went down, and distribute the external orders that was queued for it.
	- The isolated elevator should assume that all the other elevators have gone down and take over all
	  the queued orders for the other elevators. A little overkill, but its so far the only way.

- The lost elevator connects back with the others after some time:
	Both the isolated elevator and one of the elevators in the group are now masters, but only one master 
	is required. Need to pick a master, random master is probably ok. 
	Two elevators should be more efficient than one, so if an elevator filled with orders is connected to 
	an elevator without orders, the external orders should be distributed between the elevators.
	Here it has a lot to say what way the elevator is going and which floor it is in.




