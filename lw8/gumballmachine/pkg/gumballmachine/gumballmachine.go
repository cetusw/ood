package gumballmachine

import (
	"fmt"
	"io"
	"sync"
)

const (
	maxCoinsCount = 5
)

type state interface {
	insertQuarter()
	ejectQuarter()
	turnCrank()
	dispense()
	String() string
}

type GumballMachine struct {
	mu sync.Mutex

	ballsCount   uint
	coinsCount   uint
	currentState state

	soldOutState    state
	noQuarterState  state
	hasQuarterState state
	soldState       state

	writer io.Writer
}

func NewGumballMachine(numBalls uint, w io.Writer) *GumballMachine {
	if w == nil {
		w = io.Discard
	}

	m := &GumballMachine{
		ballsCount: numBalls,
		writer:     w,
	}

	m.soldOutState = &soldOutState{machine: m}
	m.noQuarterState = &noQuarterState{machine: m}
	m.hasQuarterState = &hasQuarterState{machine: m}
	m.soldState = &soldState{machine: m}

	if m.ballsCount > 0 {
		m.currentState = m.noQuarterState
	} else {
		m.currentState = m.soldOutState
	}

	return m
}

func (m *GumballMachine) InsertQuarter() {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.coinsCount < maxCoinsCount {
		m.currentState.insertQuarter()
	} else {
		fmt.Fprintln(m.writer, "The maximum number of coins has been reached")
	}
}

func (m *GumballMachine) EjectQuarter() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.currentState.ejectQuarter()
}

func (m *GumballMachine) TurnCrank() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.currentState.turnCrank()
	m.currentState.dispense()
}

func (m *GumballMachine) Refill(numBalls uint) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if m.currentState == m.soldState {
		fmt.Fprintln(m.writer, "Cannot refill while dispensing a gumball")
		return
	}

	m.ballsCount += numBalls
	if m.currentState == m.soldOutState {
		if m.coinsCount > 0 {
			m.currentState = m.hasQuarterState
		} else {
			m.currentState = m.noQuarterState
		}
	}
}

func (m *GumballMachine) String() string {
	m.mu.Lock()
	defer m.mu.Unlock()

	suffix := "s"
	if m.ballsCount == 1 {
		suffix = ""
	}

	return fmt.Sprintf(`
Mighty Gumball, Inc.
Go-enabled Standing Gumball Model #2025
Inventory: %d gumball%s
Machine is %s
`, m.ballsCount, suffix, m.currentState.String())
}

func (m *GumballMachine) GetBallCount() uint {
	return m.ballsCount
}

func (m *GumballMachine) releaseBall() {
	fmt.Fprintln(m.writer, "A gumball comes rolling out the slot...")
	if m.ballsCount > 0 {
		m.ballsCount--
	}
}

func (m *GumballMachine) setState(s state) {
	m.currentState = s
}
