package asterisktoolkit

import (
	"context"

	"github.com/CyCoreSystems/ari/v6"
	"github.com/CyCoreSystems/ari/v6/client/native"
)

// SessionFunc is called on each incoming phone call in a goroutine
// and typically manages the state of a specific session.
//
// It accepts three arguments:
//  1. context.Context - contains a series of values that can be extracted in this package
//  2. *ari.CallerID - the Caller ID of the incoming call
//  3. *ari.ChannelHandle - operations on the incoming channel
type SessionFunc func(context.Context, *ari.CallerID, *ari.ChannelHandle) error

// Controller is an ARI App Controller
type Controller struct {
	answer, hangup bool
	opts           *native.Options
	sf             SessionFunc
}

// New creates a new Controller instance
func New(opts *native.Options, sf SessionFunc) (c *Controller, err error) {
	c = &Controller{
		hangup: true,
		answer: true,
		opts:   opts,
		sf:     sf,
	}

	return
}

// DontHangup can be used to stop the Controller hanging up after running, such as
// when a Controller needs to redirect to another extension
func (c *Controller) DontHangup() {
	c.hangup = false
}

// DontAnswer can be used to stop the Controller trying to answer a call, such as when
// it has already answered a call (like when one stasis app has called another, or if
// this controller is being called as a Hangup Handler)
func (c *Controller) DontAnswer() {
	c.answer = false
}

// Run connects to ARI and handles new Sessions, ultimately passing into the
// developer specified SessionFunc
func (c *Controller) Run() (err error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cl, err := native.Connect(c.opts)
	if err != nil {
		return
	}

	sub := cl.Bus().Subscribe(nil, ari.Events.StasisStart)

	for {
		select {
		case e := <-sub.Events():
			switch v := e.(type) {
			case *ari.StasisStart:
				ctx = context.WithValue(ctx, sessionContextArgsValue{}, v.Args)

				go c.handle(
					ctx,
					v.Channel.Caller,
					cl.Channel().Get(v.Key(ari.ChannelKey, v.Channel.ID)),
				)
			}

		case <-ctx.Done():
			return
		}
	}
}

func (c *Controller) handle(ctx context.Context, cid *ari.CallerID, h *ari.ChannelHandle) (err error) {
	if c.hangup {
		defer h.Hangup() //nolint:errcheck
	}

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	end := h.Subscribe(ari.Events.StasisEnd)
	defer end.Cancel()

	// Sometimes the dialplan answers, sometimes this stasis app is called
	// from another, and sometimes we're running in a hangup handler; in which
	// case trying to error again will blow up in our faces
	if c.answer {
		if err = h.Answer(); err != nil {
			return
		}
	}

	select {
	case <-end.Events():
		return

	default:
		err = c.sf(ctx, cid, h)
		if err != nil {
			return
		}
	}

	return
}
