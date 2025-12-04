package gumballmachine

import (
	"bytes"
	"strings"
	"testing"
)

func TestGumBallMachine_NoQuarterStateInsertQuarter(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.setState(m.noQuarterState)
	buf.Reset()

	m.InsertQuarter()

	assertOutput(t, buf, "You inserted a quarter")
	assertState(t, m, m.hasQuarterState)
}

func TestGumBallMachine_NoQuarterStateEjectQuarter(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.setState(m.noQuarterState)
	buf.Reset()

	m.EjectQuarter()

	assertOutput(t, buf, "You haven't inserted a quarter")
	assertState(t, m, m.noQuarterState)
}

func TestGumBallMachine_NoQuarterStateTurnCrank(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.setState(m.noQuarterState)
	buf.Reset()

	m.TurnCrank()

	output := buf.String()
	if !strings.Contains(output, "You turned but there's no quarter") {
		t.Errorf("Missing expected turn message. Got: %s", output)
	}
	if !strings.Contains(output, "You need to pay first") {
		t.Errorf("Missing expected dispense message. Got: %s", output)
	}
	assertState(t, m, m.noQuarterState)
}

func TestGumBallMachine_NoQuarterStateDispense(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.setState(m.noQuarterState)
	buf.Reset()

	m.currentState.dispense()

	assertOutput(t, buf, "You need to pay first")
	assertState(t, m, m.noQuarterState)
}

func TestGumBallMachine_HasQuarterStateInsertQuarter(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.setState(m.hasQuarterState)
	buf.Reset()

	m.InsertQuarter()

	assertOutput(t, buf, "You inserted a quarter")
	assertState(t, m, m.hasQuarterState)
}

func TestGumBallMachine_HasQuarterStateEjectQuarter(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.coinsCount = 1
	m.setState(m.hasQuarterState)
	buf.Reset()

	m.EjectQuarter()

	assertOutput(t, buf, "Quarter returned")
	assertState(t, m, m.noQuarterState)
}

func TestGumBallMachine_HasQuarterStateTurnCrank(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.coinsCount = 1
	m.setState(m.hasQuarterState)
	buf.Reset()

	m.TurnCrank()

	output := buf.String()
	if !strings.Contains(output, "You turned...") {
		t.Errorf("Expected 'You turned...', got: %s", output)
	}
	if !strings.Contains(output, "A gumball comes rolling out") {
		t.Errorf("Expected rolling out message, got: %s", output)
	}

	assertState(t, m, m.noQuarterState)
}

func TestGumBallMachine_HasQuarterStateDispense(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.setState(m.hasQuarterState)
	buf.Reset()

	m.currentState.dispense()

	assertOutput(t, buf, "No gumball dispensed")
	assertState(t, m, m.hasQuarterState)
}

func TestGumBallMachine_SoldStateInsertQuarter(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.setState(m.soldState)
	buf.Reset()

	m.InsertQuarter()

	assertOutput(t, buf, "Please wait, we're already giving you a gumball")
	assertState(t, m, m.soldState)
}

func TestGumBallMachine_SoldStateEjectQuarter(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.coinsCount = 1
	m.setState(m.soldState)
	buf.Reset()

	m.EjectQuarter()

	assertOutput(t, buf, "Sorry you already turned the crank")
	assertState(t, m, m.soldState)
}

func TestGumBallMachine_SoldStateTurnCrank(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(5, buf)
	m.setState(m.soldState)
	buf.Reset()

	m.currentState.turnCrank()

	assertOutput(t, buf, "Turning twice doesn't get you another gumball")
	assertState(t, m, m.soldState)
}

func TestGumBallMachine_SoldStateDispenseWithGumballs(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(2, buf)
	m.coinsCount = 1
	m.setState(m.soldState)
	buf.Reset()

	m.currentState.dispense()

	output := buf.String()
	if !strings.Contains(output, "A gumball comes rolling out") {
		t.Errorf("Expected gumball rolling out, got: %s", output)
	}
	assertState(t, m, m.noQuarterState)
	if m.ballsCount != 1 {
		t.Errorf("Expected 1 ball left, got %d", m.ballsCount)
	}
}

func TestGumBallMachine_SoldStateDispenseLastGumball(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(1, buf)
	m.setState(m.soldState)
	buf.Reset()

	m.currentState.dispense()

	output := buf.String()
	if !strings.Contains(output, "Oops, out of gumballs") {
		t.Errorf("Expected out of gumballs message, got: %s", output)
	}
	assertState(t, m, m.soldOutState)
	if m.ballsCount != 0 {
		t.Errorf("Expected 0 balls left, got %d", m.ballsCount)
	}
}

func TestGumBallMachine_SoldOutStateInsertQuarter(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(0, buf)
	m.setState(m.soldOutState)
	buf.Reset()

	m.InsertQuarter()

	assertOutput(t, buf, "You can't insert a quarter, the machine is sold out")
	assertState(t, m, m.soldOutState)
}

func TestGumBallMachine_SoldOutStateEjectQuarter(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(0, buf)
	m.setState(m.soldOutState)
	buf.Reset()

	m.EjectQuarter()

	assertOutput(t, buf, "You haven't inserted a quarter")
	assertState(t, m, m.soldOutState)
}

func TestGumBallMachine_SoldOutStateTurnCrank(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(0, buf)
	m.setState(m.soldOutState)
	buf.Reset()

	m.TurnCrank()

	output := buf.String()
	if !strings.Contains(output, "You turned but there's no gumballs") {
		t.Errorf("Missing turn message. Got: %s", output)
	}
	if !strings.Contains(output, "No gumball dispensed") {
		t.Errorf("Missing dispense message. Got: %s", output)
	}
	assertState(t, m, m.soldOutState)
}

func TestGumBallMachine_SoldOutStateDispense(t *testing.T) {
	buf := new(bytes.Buffer)
	m := NewGumballMachine(0, buf)
	m.setState(m.soldOutState)
	buf.Reset()

	m.currentState.dispense()

	assertOutput(t, buf, "No gumball dispensed")
	assertState(t, m, m.soldOutState)
}

func assertOutput(t *testing.T, buffer *bytes.Buffer, expectedPart string) {
	t.Helper()
	got := buffer.String()
	if !strings.Contains(got, expectedPart) {
		t.Errorf("Expected output containing: %q\nGot: %q", expectedPart, got)
	}
	buffer.Reset()
}

func assertState(t *testing.T, m *GumballMachine, expectedState state) {
	t.Helper()
	if m.currentState != expectedState {
		t.Errorf("Expected state to be %T, but got %T", expectedState, m.currentState)
	}
}
