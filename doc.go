/*
package asterisktoolkit provides an opinionated wrapper around ARI connections
that fills a very specific need I have for my daft little apps.

It can be used as per:

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
	    c, err := asterisktoolkit.New(asterisktoolkit.Options("my-user", "supersecurepassword", "my-app", "http://localhost/ari", handleCall)
	    if err != nil {
		panic(err)
	    }

	    panic(c.Run())
	}

Where `handleCall` is your ARI app; perhaps playing obscene and bawdy songs, or reading out the pools results, or whatever.
*/
package asterisktoolkit
