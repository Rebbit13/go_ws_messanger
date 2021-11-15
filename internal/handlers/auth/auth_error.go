package auth

type AuthHandlerError struct {
	Text string
}

func (stringError *AuthHandlerError) Error() string {
	return stringError.Text
}
