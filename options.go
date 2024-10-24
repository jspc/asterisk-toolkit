package asterisktoolkit

import (
	"net/url"

	"github.com/CyCoreSystems/ari/v6/client/native"
)

// Options returns a configured ari Options struct based on sensible
// defaults and a Websocket URL inferred from the ARI URL.
//
// It exists as a convenience wrapper only
func Options(username, password, application, ari string) *native.Options {
	return &native.Options{
		Application:  application,
		Username:     username,
		Password:     password,
		URL:          ari,
		WebsocketURL: ariUrlToWS(ari),
		SubscribeAll: true,
	}
}

func ariUrlToWS(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return ""
	}

	u.Scheme = "ws"
	u.Path = "/ari/events"

	return u.String()
}
