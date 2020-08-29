package errs

type ProcessError struct {
	Code string
	Msg  string
}

func (error ProcessError) Error() string {
	return error.Code + "---" + error.Msg
}
