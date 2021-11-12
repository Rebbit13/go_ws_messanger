package password_hash

type HashServiceError struct {
	Text string
}

func (stringError *HashServiceError) Error() string {
	return stringError.Text
}
