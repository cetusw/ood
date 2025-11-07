package robot

import (
	"fmt"
)

type WalkDirection int

const (
	North WalkDirection = iota
	South
	West
	East
)

var directionToString = map[WalkDirection]string{
	North: "north",
	South: "south",
	West:  "west",
	East:  "east",
}

type Robot struct {
	turnedOn  bool
	direction *WalkDirection
}

func NewRobot() *Robot {
	return &Robot{}
}

func (r *Robot) TurnOn() {
	if !r.turnedOn {
		r.turnedOn = true
		fmt.Println("It am waiting for your commands")
	}
}

func (r *Robot) TurnOff() {
	if r.turnedOn {
		r.turnedOn = false
		r.direction = nil
		fmt.Println("It is a pleasure to serve you")
	}
}

func (r *Robot) Walk(direction WalkDirection) {
	if r.turnedOn {
		r.direction = &direction
		fmt.Printf("Walking %s\n", directionToString[direction])
	} else {
		fmt.Println("The robot should be turned on first")
	}
}

func (r *Robot) Stop() {
	if r.turnedOn {
		if r.direction != nil {
			r.direction = nil
			fmt.Println("Stopped")
		} else {
			fmt.Println("I am staying still")
		}
	} else {
		fmt.Println("The robot should be turned on first")
	}
}
