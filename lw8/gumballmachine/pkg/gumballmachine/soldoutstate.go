package gumballmachine

import "fmt"

type soldOutState struct {
	machine *GumballMachine
}

func (s *soldOutState) insertQuarter() {
	fmt.Fprintln(s.machine.writer, "You can't insert a quarter, the machine is sold out")
}
func (s *soldOutState) ejectQuarter() {
	if s.machine.coinsCount > 0 {
		s.machine.coinsCount = 0
		fmt.Fprintln(s.machine.writer, "Quarter returned")
		return
	}
	fmt.Fprintln(s.machine.writer, "You haven't inserted a quarter")
}
func (s *soldOutState) turnCrank() {
	fmt.Fprintln(s.machine.writer, "You turned but there's no gumballs")
}
func (s *soldOutState) dispense() {
	fmt.Fprintln(s.machine.writer, "No gumball dispensed")
}
func (s *soldOutState) String() string { return "sold out" }
