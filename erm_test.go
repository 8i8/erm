package erm

import (
	"errors"
	"fmt"
	"testing"
)

var errType = errors.New("main error type")
var myErrMsg = NewErrMsg(errType, "Ops, an error occured!")

func errGen() error {
	return myErrMsg.Write("This is the key message!")
}

func errWrapper() error {
	err := errGen()
	return fmt.Errorf("outer shell: %w", err)
}

func TestErm(t *testing.T) {
	const fname = "TestErm"
	err := errWrapper()

	//if !errors.Is(err, myErrMsg) {
	if !errors.Is(err, err) {
		t.Errorf("%s: want %T got %T", fname, myErrMsg, errors.Unwrap(err))
	}
	errCore, ok := errors.Unwrap(err).(ErrMsg)
	if !ok {
		t.Errorf("%s: want %T got %T", fname, errType, errCore)
	}
}
