package gumballmachine

import "fmt"

type hasQuarterState struct {
	machine *GumballMachine
}

func (s *hasQuarterState) insertQuarter() {
	fmt.Fprintln(s.machine.writer, "You can't insert another quarter")
}
func (s *hasQuarterState) ejectQuarter() {
	fmt.Fprintln(s.machine.writer, "Quarter returned")
	s.machine.setState(s.machine.noQuarterState)
}
func (s *hasQuarterState) turnCrank() {
	fmt.Fprintln(s.machine.writer, "You turned...")
	s.machine.setState(s.machine.soldState)
}
func (s *hasQuarterState) dispense() {
	fmt.Fprintln(s.machine.writer, "No gumball dispensed")
}
func (s *hasQuarterState) String() string { return "waiting for turn of crank" }
