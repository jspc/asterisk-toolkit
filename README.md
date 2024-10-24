# Asterisk Toolkit

[![GoDoc](https://img.shields.io/badge/pkg.go.dev-doc-blue)](http://pkg.go.dev/.)
[![Go Report Card](https://goreportcard.com/badge/.)](https://goreportcard.com/report/.)

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

### func [Options](/options.go#L13)

`func Options(username, password, application, ari string) *native.Options`

Options returns a configured ari Options struct based on sensible
defaults and a Websocket URL inferred from the ARI URL.

It exists as a convenience wrapper only

## Types

### type [Controller](/controller.go#L15)

`type Controller struct { ... }`

Controller is an ARI App Controller

#### func (*Controller) [Run](/controller.go#L32)

`func (c *Controller) Run() (err error)`

Run connects to ARI and handles new Sessions, ultimately passing into the
developer specified SessionFunc

### type [SessionFunc](/controller.go#L12)

`type SessionFunc func(context.Context, *ari.ChannelHandle) error`

SessionFunc is called on each incoming phone call in a goroutine
and typically manages the state of a specific session

---
Readme created from Go doc with [goreadme](https://github.com/posener/goreadme)
