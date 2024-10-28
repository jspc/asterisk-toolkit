# Asterisk Toolkit

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/github.com/jspc/asterisk-toolkit)
[![Go Report Card](https://goreportcard.com/badge/github.com/jspc/asterisk-toolkit)](https://goreportcard.com/report/github.com/jspc/asterisk-toolkit)

package asterisktoolkit provides an opinionated wrapper around ARI connections
that fills a very specific need I have for my daft little apps.

It can be used as per:

```go
package main

import (
    "context"

    "github.com/CyCoreSystems/ari/v6"
    "github.com/jspc/asterisk-toolkit"
)

func handleCall(ctx context.Context, h *ari.ChannelHandle) error {
    // Do a thing
    return nil
}

func main() {
    c, err := asterisktoolkit.New(asterisktoolkit.Options("my-user", "supersecurepassword", "my-app", "[http://localhost/ari](http://localhost/ari)", handleCall)
    if err != nil {
	panic(err)
    }

    panic(c.Run())
}
```

Where `handleCall` is your ARI app; perhaps playing obscene and bawdy songs, or reading out the pools results, or whatever.

## Functions

### func [GetArguments](/context.go#L12)

`func GetArguments(ctx context.Context) []string`

GetArguments can be run on the context passed to a SessionFunc
in order to get the arguments the StasisApp was started with
(if any).

This can be useful for when an app is used in many places.

### func [Options](/options.go#L13)

`func Options(username, password, application, ari string) *native.Options`

Options returns a configured ari Options struct based on sensible
defaults and a Websocket URL inferred from the ARI URL.

It exists as a convenience wrapper only

## Types

### type [Controller](/controller.go#L20)

`type Controller struct { ... }`

Controller is an ARI App Controller

#### func (*Controller) [DontAnswer](/controller.go#L46)

`func (c *Controller) DontAnswer()`

DontAnswer can be used to stop the Controller trying to answer a call, such as when
it has already answered a call (like when one stasis app has called another, or if
this controller is being called as a Hangup Handler)

#### func (*Controller) [DontHangup](/controller.go#L39)

`func (c *Controller) DontHangup()`

DontHangup can be used to stop the Controller hanging up after running, such as
when a Controller needs to redirect to another extension

#### func (*Controller) [Run](/controller.go#L52)

`func (c *Controller) Run() (err error)`

Run connects to ARI and handles new Sessions, ultimately passing into the
developer specified SessionFunc

### type [SessionFunc](/controller.go#L17)

`type SessionFunc func(context.Context, *ari.CallerID, *ari.ChannelHandle) error`

SessionFunc is called on each incoming phone call in a goroutine
and typically manages the state of a specific session.

It accepts three arguments:

```go
1. context.Context - contains a series of values that can be extracted in this package
2. *ari.CallerID - the Caller ID of the incoming call
3. *ari.ChannelHandle - operations on the incoming channel
```

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
