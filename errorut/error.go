package errorut

type myErr struct {
	name string
	msg  string
}

func New(name, msg string) error {
	return myErr{name: name, msg: msg}
}

func (m myErr) Error() string {
	return m.msg
}

func IsErrByName(err error, name string) bool {
	if err == nil {
		return false
	}

	if realErr, ok := err.(myErr); ok {
		return realErr.name == name
	}
	return false
}
