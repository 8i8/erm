package erm

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

var errType = errors.New("main error type")
var myErrMsg = New(errType, "Ops, an error occured!")

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

	if _, ok := errors.Unwrap(err).(ErrMsg); !ok {
		t.Errorf("%s: want %T got %T", fname, myErrMsg, err)
	}
	str := fmt.Sprint(errors.Unwrap(err).Error())
	str2 := fmt.Sprint(errType.Error())
	if strings.Compare(str, str2) != 0 {
		t.Errorf("%s: want %q got %q", fname, str2, str)
	}
}
