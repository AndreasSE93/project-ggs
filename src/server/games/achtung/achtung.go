package achtung

import (
	"server/connection"
)

type Player struct {
	Conn connection.Connector
	Color string
	Number int
}

type Achtung struct {
	POne, PTwo, PThree, PFour Player
	Field [10000]int
}

func GenPlayer(a *Achtung, clientList []connection.Connector) {
	a.POne.Conn = clientList[0]
	a.POne.Color = "Black"
	a.POne.Number = 1
	a.PTwo.Conn = clientList[1]
	a.PTwo.Color = "Red"
	a.PTwo.Number = 2
	a.PThree.Conn = clientList[2]
	a.PThree.Color = "Green"
	a.PThree.Number = 3
	a.PFour.Conn = clientList[3]
	a.PFour.Color = "Yellow"
	a.PFour.Number = 4
}
