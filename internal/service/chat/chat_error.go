package chat

type RoomError struct {
	Text string
}

func (stringError *RoomError) Error() string {
	return stringError.Text
}