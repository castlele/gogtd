package commands

import (
	"bytes"
	"testing"
)

func TestExecute(t *testing.T) {
	expMsg := "test message"
	var actualMsgBuf bytes.Buffer
	sut := newHelpCommand(expMsg, &actualMsgBuf)

	sut.Execute()

	actualMsg := actualMsgBuf.String()

	if actualMsg != expMsg {
		t.Fatalf("got %q, want %q", actualMsg, expMsg)
	}
}
