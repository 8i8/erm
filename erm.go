package erm

type errmsg struct {
	err string
	msg string
}

type ErrMsg interface {
	Error() string
	Read() string
}

func (e errmsg) Error() string           { return e.err }
func (e errmsg) Read() string            { return e.msg }
func (e errmsg) Write(msg string) errmsg { e.msg = msg; return e }

// NewErrMsg returns and error of types err that returns a message when
// calle.
func NewErrMsg(err interface{}, m string) errmsg {
	if e, ok := err.(error); ok {
		return errmsg{err: e.Error(), msg: m}
	}
	if e, ok := err.(string); ok {
		return errmsg{err: e, msg: m}
	}
	panic("MewErrMsg requirese either an error or a string")
}
