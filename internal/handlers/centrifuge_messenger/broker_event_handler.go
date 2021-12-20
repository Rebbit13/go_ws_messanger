package centrifuge_messenger

import (
	"context"
	"github.com/centrifugal/centrifuge"
)

func onConnecting(context context.Context, event centrifuge.ConnectEvent) (centrifuge.ConnectReply, error) {
	cred, _ := centrifuge.GetCredentials(context)
	return centrifuge.ConnectReply{
		Data: []byte(`{}`),
		Subscriptions: map[string]centrifuge.SubscribeOptions{
			"#" + cred.UserID: {Recover: true, Presence: true, JoinLeave: true},
		},
	}, nil
}
