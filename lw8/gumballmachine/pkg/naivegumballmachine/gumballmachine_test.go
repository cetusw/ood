package naivegumballmachine

import (
	"bytes"
	"strings"
	"testing"
)

func TestNaiveMultiCoin_Limit(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(10, buf)

	for i := 0; i < 5; i++ {
		m.InsertQuarter()
	}

	buf.Reset()
	m.InsertQuarter()

	assertOutput(t, buf, "The maximum number of coins has been reached")
	if m.coinsCount != 5 {
		t.Errorf("Expected 5 coins, got %d", m.coinsCount)
	}
}

func TestNaiveMultiCoin_CrankSequence(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)

	m.InsertQuarter()
	m.InsertQuarter()

	m.TurnCrank()
	if m.coinsCount != 1 {
		t.Errorf("Expected 1 coin left, got %d", m.coinsCount)
	}
	assertState(t, m, HasQuarter)

	m.TurnCrank()
	if m.coinsCount != 0 {
		t.Errorf("Expected 0 coins left, got %d", m.coinsCount)
	}
	assertState(t, m, NoQuarter)
}

func TestNaiveMultiCoin_CoinsMoreThanBalls(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(1, buf)

	m.InsertQuarter()
	m.InsertQuarter()

	m.TurnCrank()

	if m.ballsCount != 0 {
		t.Errorf("Balls not decremented")
	}
	if m.coinsCount != 1 {
		t.Errorf("Expected 1 coin left (credit), got %d", m.coinsCount)
	}
	assertState(t, m, SoldOut)

	m.EjectQuarter()
	if m.coinsCount != 0 {
		t.Errorf("Coins not returned from SoldOut")
	}
}

func TestNaive_RefillFromSoldOutWithNoCoins(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(0, buf)
	m.Refill(5)

	if m.ballsCount != 5 {
		t.Errorf("Expected 5 balls, got %d", m.ballsCount)
	}
	assertState(t, m, NoQuarter)
}

func TestNaive_RefillFromSoldOutWithCoins(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(0, buf)

	m.InsertQuarter()
	if m.coinsCount != 1 {
		t.Errorf("Expected 1 coin, got %d", m.coinsCount)
	}

	m.Refill(5)

	if m.ballsCount != 5 {
		t.Errorf("Expected 5 balls, got %d", m.ballsCount)
	}
	assertState(t, m, HasQuarter)
}

func TestNaive_RefillIncrements(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.Refill(5)

	if m.ballsCount != 10 {
		t.Errorf("Expected 10 balls (5+5), got %d", m.ballsCount)
	}
}

func assertOutput(t *testing.T, buffer *bytes.Buffer, expectedPart string) {
	t.Helper()
	got := buffer.String()
	if !strings.Contains(got, expectedPart) {
		t.Errorf("Expected output containing: %q\nGot: %q", expectedPart, got)
	}
	buffer.Reset()
}

func assertState(t *testing.T, m *GumballMachine, expectedState State) {
	t.Helper()
	if m.state != expectedState {
		t.Errorf("Expected state %s, got %s", expectedState, m.state)
	}
}
