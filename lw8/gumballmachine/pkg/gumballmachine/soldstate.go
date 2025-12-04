package gumballmachine

import "fmt"

type soldState struct {
	machine *GumballMachine
}

func (s *soldState) insertQuarter() {
	fmt.Fprintln(s.machine.writer, "Please wait, we're already giving you a gumball")
}

func (s *soldState) ejectQuarter() {
	fmt.Fprintln(s.machine.writer, "Sorry you already turned the crank")
}

func (s *soldState) turnCrank() {
	fmt.Fprintln(s.machine.writer, "Turning twice doesn't get you another gumball")
}

func (s *soldState) dispense() {
	s.machine.coinsCount--
	s.machine.releaseBall()
	if s.machine.ballsCount <= 0 {
		fmt.Fprintln(s.machine.writer, "Oops, out of gumballs")
		s.machine.setState(s.machine.soldOutState)
		return
	}
	if s.machine.coinsCount <= 0 {
		s.machine.setState(s.machine.noQuarterState)
		return
	}
	s.machine.setState(s.machine.hasQuarterState)
}

func (s *soldState) String() string { return "delivering a gumball" }
