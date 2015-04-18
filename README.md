# TTK4145-Simulator
Simulator of the elevator model at "Sanntidssalen" at NTNU

How to use: 
The simulator core have the same interface as the driver in this rep. The driver wraps Go-calls to C-calls and imports Comedi.

Recomended way to implement:
Set up the elevator program on the lab with the driver/wraper. Then edit the import statement from "driver" to "simulator". The simulator should start and self initialise. When you need to button press commands, you need to start the simulatorInterface.go. This program sends buttoncommands from the keyboard to the simulator core. The buttons are mapped from Z->V on internal orders, Q->E for external orders Up, and S->F for external orders Down.