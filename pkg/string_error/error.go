package string_error

type StringError struct {
	Text string
}

func (stringError *StringError) Error() string {
	return stringError.Text
}
