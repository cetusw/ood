package naivegumballmachine

import (
	"fmt"
	"io"
)

type State int

const (
	SoldOut State = iota
	NoQuarter
	HasQuarter
	Sold
)

const (
	maxCoinsCount = 5
)

func (s State) String() string {
	switch s {
	case SoldOut:
		return "sold out"
	case NoQuarter:
		return "waiting for quarter"
	case HasQuarter:
		return "waiting for turn of crank"
	case Sold:
		return "delivering a gumball"
	default:
		return "unknown state"
	}
}

type GumballMachine struct {
	ballsCount int
	coinsCount int
	state      State

	writer io.Writer
}

func NewGumballMachine(count int, w io.Writer) *GumballMachine {
	if w == nil {
		w = io.Discard
	}

	initialState := SoldOut
	if count > 0 {
		initialState = NoQuarter
	}

	return &GumballMachine{
		ballsCount: count,
		coinsCount: 0,
		state:      initialState,
		writer:     w,
	}
}

func (m *GumballMachine) InsertQuarter() {
	if m.coinsCount >= maxCoinsCount {
		fmt.Fprintln(m.writer, "The maximum number of coins has been reached")
		return
	}
	switch m.state {
	case SoldOut:
		fmt.Fprintln(m.writer, "You can't insert a quarter, the machine is sold out")
	case NoQuarter:
		fmt.Fprintln(m.writer, "You inserted a quarter")
		m.coinsCount++
		m.state = HasQuarter
	case HasQuarter:
		fmt.Fprintln(m.writer, "You inserted a quarter")
		m.coinsCount++
	case Sold:
		fmt.Fprintln(m.writer, "Please wait, we're already giving you a gumball")
	}
}

func (m *GumballMachine) EjectQuarter() {
	switch m.state {
	case HasQuarter:
		fmt.Fprintln(m.writer, "Quarter returned")
		m.coinsCount = 0
		m.state = NoQuarter
	case NoQuarter:
		fmt.Fprintln(m.writer, "You haven't inserted a quarter")
	case Sold:
		fmt.Fprintln(m.writer, "Sorry you already turned the crank")
	case SoldOut:
		if m.coinsCount == 0 {
			fmt.Fprintln(m.writer, "You can't eject, you haven't inserted a quarter yet")
			return
		}
		m.coinsCount = 0
		fmt.Fprintln(m.writer, "Quarter returned")
	}
}

func (m *GumballMachine) TurnCrank() {
	switch m.state {
	case SoldOut:
		fmt.Fprintln(m.writer, "You turned but there's no gumballs")
	case NoQuarter:
		fmt.Fprintln(m.writer, "You turned but there's no quarter")
	case HasQuarter:
		fmt.Fprintln(m.writer, "You turned...")
		m.state = Sold
		m.dispense()
	case Sold:
		fmt.Fprintln(m.writer, "Turning twice doesn't get you another gumball")
	}
}

func (m *GumballMachine) Refill(numBalls int) {
	if m.state == Sold {
		fmt.Fprintln(m.writer, "Cannot refill while dispensing a gumball")
		return
	}

	m.ballsCount += numBalls

	if m.state == SoldOut && m.ballsCount > 0 {
		if m.coinsCount > 0 {
			m.state = HasQuarter
		} else {
			m.state = NoQuarter
		}
	}
}

func (m *GumballMachine) String() string {
	suffix := "s"
	if m.ballsCount == 1 {
		suffix = ""
	}

	return fmt.Sprintf(`
Mighty Gumball, Inc.
Go-enabled Standing Gumball Model #2016 (Naive)
Inventory: %d gumball%s
Machine is %s
`, m.ballsCount, suffix, m.state.String())
}

func (m *GumballMachine) dispense() {
	switch m.state {
	case Sold:
		fmt.Fprintln(m.writer, "A gumball comes rolling out the slot")
		m.ballsCount--
		m.coinsCount--

		if m.ballsCount == 0 {
			fmt.Fprintln(m.writer, "Oops, out of gumballs")
			m.state = SoldOut
			return
		}

		if m.coinsCount == 0 {
			m.state = NoQuarter
		} else {
			m.state = HasQuarter
		}
	case NoQuarter:
		fmt.Fprintln(m.writer, "You need to pay first")
	case SoldOut, HasQuarter:
		fmt.Fprintln(m.writer, "No gumball dispensed")
	}
}
