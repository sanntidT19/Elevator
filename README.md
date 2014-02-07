Elevator
========

-Our main elevator repository

-Please note conventions here:


-Questions for Student assistant:


All network packages is sendt by the form of a struct.



**MODULES**:

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

*Slave Work(IO unit):\\
	-to IO to slave Communication\\

*Setup Module\\
	-choose who is master and slaves.\\
	-obtain ip adresses, and numbering of slaves.\\

*Selfcheck Module\\
	-Working with master and slaves, with detecting change in number of elevators.\\
	(if no message recived in 0.5 sec assume connecton is lost. Assumed 10 messages per secound is sendt).\\

