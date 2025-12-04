package gumballmachine

import (
	"bytes"
	"testing"
)

func TestMultiCoin_Insert(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(10, buf)

	m.InsertQuarter()
	m.InsertQuarter()
	m.InsertQuarter()

	if m.coinsCount != 3 {
		t.Errorf("Expected 3 coins, got %d", m.coinsCount)
	}

	assertOutput(t, buf, "You inserted a quarter")
	assertState(t, m, m.hasQuarterState)
}

func TestMultiCoin_Limit(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(10, buf)

	for i := 0; i < 5; i++ {
		m.InsertQuarter()
	}
	buf.Reset()

	m.InsertQuarter()

	assertOutput(t, buf, "The maximum number of coins has been reached")
	if m.coinsCount != 5 {
		t.Errorf("Expected cap at 5 coins, got %d", m.coinsCount)
	}
}

func TestMultiCoin_EjectAll(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(10, buf)

	m.InsertQuarter()
	m.InsertQuarter()

	buf.Reset()
	m.EjectQuarter()

	if m.coinsCount != 0 {
		t.Errorf("Expected 0 coins after eject, got %d", m.coinsCount)
	}
	assertState(t, m, m.noQuarterState)
}

func TestMultiCoin_CrankSequence(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(10, buf)

	m.InsertQuarter()
	m.InsertQuarter()

	m.TurnCrank()
	if m.ballsCount != 9 {
		t.Errorf("Expected 9 balls, got %d", m.ballsCount)
	}
	if m.coinsCount != 1 {
		t.Errorf("Expected 1 coin left, got %d", m.coinsCount)
	}
	assertState(t, m, m.hasQuarterState)

	m.TurnCrank()
	if m.ballsCount != 8 {
		t.Errorf("Expected 8 balls, got %d", m.ballsCount)
	}
	if m.coinsCount != 0 {
		t.Errorf("Expected 0 coins left, got %d", m.coinsCount)
	}
	assertState(t, m, m.noQuarterState)
}

func TestMultiCoin_CoinsMoreThanBalls(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(1, buf)

	m.InsertQuarter()
	m.InsertQuarter()

	m.TurnCrank()

	if m.ballsCount != 0 {
		t.Errorf("Expected 0 balls, got %d", m.ballsCount)
	}
	if m.coinsCount != 1 {
		t.Errorf("Expected 1 coin remaining (credit), got %d", m.coinsCount)
	}
	assertState(t, m, m.soldOutState)

	buf.Reset()
	m.TurnCrank()

	assertOutput(t, buf, "You turned but there's no gumballs")
	if m.coinsCount != 1 {
		t.Errorf("Expected 1 coin after TurnCrank from SoldOut, got %d", m.coinsCount)
	}

	buf.Reset()
	m.EjectQuarter()

	assertOutput(t, buf, "Quarter returned")
	if m.coinsCount != 0 {
		t.Errorf("Expected 0 coins after eject from SoldOut, got %d", m.coinsCount)
	}
}

func TestMultiCoin_BallsMoreThanCoins(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)

	m.InsertQuarter()

	m.TurnCrank()

	if m.ballsCount != 4 {
		t.Errorf("Expected 4 balls, got %d", m.ballsCount)
	}
	if m.coinsCount != 0 {
		t.Errorf("Expected 0 coin remaining (credit), got %d", m.coinsCount)
	}
	assertState(t, m, m.noQuarterState)

	buf.Reset()
	m.TurnCrank()

	assertOutput(t, buf, "You turned but there's no quarter")
	if m.ballsCount != 4 {
		t.Errorf("Expected 4 balls, got %d", m.ballsCount)
	}
	if m.coinsCount != 0 {
		t.Errorf("Expected 0 coin after TurnCrank from SoldOut, got %d", m.coinsCount)
	}

	buf.Reset()
	m.EjectQuarter()

	assertOutput(t, buf, "You haven't inserted a quarter")
	if m.coinsCount != 0 {
		t.Errorf("Expected 0 coins after eject from SoldOut, got %d", m.coinsCount)
	}
}
