package x52

import ()

type Error struct {
	// Msg is the error message
	Msg string

	// Err is the underlying wrapped error
	Err error
}

func (err *Error) Error() string {
	output := err.Msg
	if err.Err != nil {
		output += ": " + err.Err.Error()
	}

	return output
}

func (err *Error) Unwrap() error {
	return err.Err
}

func ErrNotSupported(reason string) *Error {
	msg := "x52: not supported"
	if len(reason) > 0 {
		msg += ": " + reason
	}

	return &Error{
		Msg: msg,
	}
}

func ErrInvalidParam(reason string) *Error {
	msg := "x52: invalid parameter"
	if len(reason) > 0 {
		msg += ": " + reason
	}

	return &Error{
		Msg: msg,
	}
}

func ErrNotConnected(err error) *Error {
	return &Error{
		Msg: "x52: not connected",
		Err: err,
	}
}
