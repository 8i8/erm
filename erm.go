package erm

type errmsg struct {
	err string
	msg string
}

type ErrMsg interface {
	Error() string
	Read() string
}

type ErrMsgWriter interface {
	Error() string
	Read() string
	Write(string) errmsg
}

func (e errmsg) Error() string           { return e.err }
func (e errmsg) Read() string            { return e.msg }
func (e errmsg) Write(msg string) errmsg { e.msg = msg; return e }
func (e errmsg) Message() string         { return e.msg }

// New returns an error of types errmsg that conforms to the ErrMsg
// interface, which adds the Read{} functionality for passing messages.
func New(err interface{}, m string) errmsg {
	if e, ok := err.(error); ok {
		return errmsg{err: e.Error(), msg: m}
	}
	if e, ok := err.(string); ok {
		return errmsg{err: e, msg: m}
	}
	panic("erm.ErrMsg requires either an error or a string")
}
