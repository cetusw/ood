package gumballmachine

import (
	"bytes"
	"testing"
)

func TestMultiCoin_NoQuarterStateRefill(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.coinsCount = 0
	m.setState(m.noQuarterState)

	m.Refill(5)

	if m.ballsCount != 10 {
		t.Errorf("Expected 10 balls after Refill, got %d", m.ballsCount)
	}

	assertState(t, m, m.noQuarterState)
}

func TestMultiCoin_HasQuarterStateRefill(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.coinsCount = 1
	m.setState(m.hasQuarterState)

	m.Refill(5)

	if m.ballsCount != 10 {
		t.Errorf("Expected 10 balls after Refill, got %d", m.ballsCount)
	}

	buf.Reset()
	m.TurnCrank()

	assertOutput(t, buf, "You turned...")
	if m.ballsCount != 9 {
		t.Errorf("Expected 9 balls after Refill and TurnCrank, got %d", m.ballsCount)
	}
	if m.coinsCount != 0 {
		t.Errorf("Expected 0 coins after Refill and TurnCrank, got %d", m.coinsCount)
	}

	assertState(t, m, m.noQuarterState)
}

func TestMultiCoin_SoldOutStateRefill(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(0, buf)
	m.coinsCount = 1
	m.setState(m.soldOutState)

	m.Refill(5)

	if m.ballsCount != 5 {
		t.Errorf("Expected 5 balls after Refill, got %d", m.ballsCount)
	}

	buf.Reset()
	m.TurnCrank()

	assertOutput(t, buf, "You turned...")
	if m.ballsCount != 4 {
		t.Errorf("Expected 4 balls after Refill and TurnCrank, got %d", m.ballsCount)
	}
	if m.coinsCount != 0 {
		t.Errorf("Expected 0 coins after Refill and TurnCrank, got %d", m.coinsCount)
	}

	assertState(t, m, m.noQuarterState)
}

func TestMultiCoin_SoldOutStateRefillWithoutCoins(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(0, buf)
	m.coinsCount = 0
	m.setState(m.soldOutState)

	m.Refill(5)

	if m.ballsCount != 5 {
		t.Errorf("Expected 5 balls after Refill, got %d", m.ballsCount)
	}

	assertState(t, m, m.noQuarterState)
}
