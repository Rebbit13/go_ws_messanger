package authorization

type AuthError struct {
	Text string
}

func (stringError *AuthError) Error() string {
	return stringError.Text
}
