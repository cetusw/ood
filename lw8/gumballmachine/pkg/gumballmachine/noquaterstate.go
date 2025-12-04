package gumballmachine

import "fmt"

type noQuarterState struct {
	machine *GumballMachine
}

func (s *noQuarterState) insertQuarter() {
	fmt.Fprintln(s.machine.writer, "You inserted a quarter")
	s.machine.coinsCount++
	s.machine.setState(s.machine.hasQuarterState)
}
func (s *noQuarterState) ejectQuarter() {
	fmt.Fprintln(s.machine.writer, "You haven't inserted a quarter")
}
func (s *noQuarterState) turnCrank() {
	fmt.Fprintln(s.machine.writer, "You turned but there's no quarter")
}
func (s *noQuarterState) dispense() {
	fmt.Fprintln(s.machine.writer, "You need to pay first")
}
func (s *noQuarterState) String() string { return "waiting for quarter" }
