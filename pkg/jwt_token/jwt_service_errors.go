package jwt_token

type JWTServiceError struct {
	Text string
}

func (stringError *JWTServiceError) Error() string {
	return stringError.Text
}
