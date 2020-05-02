package x52

type x52Error struct {
	msg string
	err error
}

// x52Error satisfies the error interface
func (err *x52Error) Error() string {
	output := err.msg
	if err.err != nil {
		output += ": " + err.err.Error()
	}

	return output
}

// Unwrap returns the wrapped error
func (err *x52Error) Unwrap() error {
	return err.err
}

func errNotSupported(reason string) *x52Error {
	msg := "x52: not supported"
	if len(reason) > 0 {
		msg += ": " + reason
	}

	return &x52Error{
		msg: msg,
	}
}

func errInvalidParam(reason string) *x52Error {
	msg := "x52: invalid parameter"
	if len(reason) > 0 {
		msg += ": " + reason
	}

	return &x52Error{
		msg: msg,
	}
}

func errNotConnected(err error) *x52Error {
	return &x52Error{
		msg: "x52: not connected",
		err: err,
	}
}
