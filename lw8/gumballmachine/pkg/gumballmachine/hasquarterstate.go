package gumballmachine

import "fmt"

type hasQuarterState struct {
	machine *GumballMachine
}

func (s *hasQuarterState) insertQuarter() {
	s.machine.coinsCount++
	fmt.Fprintln(s.machine.writer, "You inserted a quarter")

}
func (s *hasQuarterState) ejectQuarter() {
	s.machine.coinsCount = 0
	s.machine.setState(s.machine.noQuarterState)
	fmt.Fprintln(s.machine.writer, "Quarter returned")
}
func (s *hasQuarterState) turnCrank() {
	fmt.Fprintln(s.machine.writer, "You turned...")
	s.machine.setState(s.machine.soldState)
}
func (s *hasQuarterState) dispense() {
	fmt.Fprintln(s.machine.writer, "No gumball dispensed")
}
func (s *hasQuarterState) String() string { return "waiting for turn of crank" }
