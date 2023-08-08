package codewars

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTcpFsmState1(t *testing.T) {
	input := []string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN"}
	expected := "CLOSE_WAIT"

	res := TraverseTCPStates(input)

	assert.Equal(t, expected, res)
}

func TestTcpFsmState2(t *testing.T) {
	input := []string{"APP_PASSIVE_OPEN", "RCV_SYN", "RCV_ACK"}
	expected := "ESTABLISHED"

	res := TraverseTCPStates(input)

	assert.Equal(t, expected, res)
}

func TestTcpFsmState3(t *testing.T) {
	input := []string{"APP_ACTIVE_OPEN", "RCV_SYN_ACK", "RCV_FIN", "APP_CLOSE"}
	expected := "LAST_ACK"

	res := TraverseTCPStates(input)

	assert.Equal(t, expected, res)
}

func TestTcpFsmState4(t *testing.T) {
	input := []string{"APP_ACTIVE_OPEN"}
	expected := "SYN_SENT"

	res := TraverseTCPStates(input)

	assert.Equal(t, expected, res)
}

func TestTcpFsmState5(t *testing.T) {
	input := []string{"APP_PASSIVE_OPEN", "RCV_SYN", "RCV_ACK", "APP_CLOSE", "APP_SEND"}
	expected := "ERROR"

	res := TraverseTCPStates(input)

	assert.Equal(t, expected, res)
}

func TestTcpFsmState6(t *testing.T) {
	input := []string{"RCV_SYN", "RCV_ACK"}
	expected := "ESTABLISHED"

	res := TraverseTCPStates(input)

	assert.Equal(t, expected, res)
}

func TestTcpFsmState7(t *testing.T) {
	input := []string{"RCV_SYN", "RCV_ACK", "APP_CLOSE"}
	expected := "FIN_WAIT_1"

	res := TraverseTCPStates(input)

	assert.Equal(t, expected, res)
}

func TestTcpFsmState8(t *testing.T) {
	input := []string{"APP_PASSIVE_OPEN", "RCV_SYN", "RCV_ACK", "APP_CLOSE", "APP_SEND"}
	expected := "ERROR"

	res := TraverseTCPStates(input)

	assert.Equal(t, expected, res)
}
