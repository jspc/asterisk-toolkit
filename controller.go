package asterisktoolkit

import (
	"context"

	"github.com/CyCoreSystems/ari/v6"
	"github.com/CyCoreSystems/ari/v6/client/native"
)

// SessionFunc is called on each incoming phone call in a goroutine
// and typically manages the state of a specific session
type SessionFunc func(context.Context, *ari.ChannelHandle) error

// Controller is an ARI App Controller
type Controller struct {
	opts *native.Options
	sf   SessionFunc
}

// New creates a new Controller instance
func New(opts *native.Options, sf SessionFunc) (c *Controller, err error) {
	c = &Controller{
		opts: opts,
		sf:   sf,
	}

	return
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

	sub := cl.Bus().Subscribe(nil, "StasisStart")

	for {
		select {
		case e := <-sub.Events():
			v := e.(*ari.StasisStart)

			go c.handle(ctx, cl.Channel().Get(v.Key(ari.ChannelKey, v.Channel.ID)))
		case <-ctx.Done():
			return
		}
	}
}

func (c *Controller) handle(ctx context.Context, h *ari.ChannelHandle) (err error) {
	defer h.Hangup() //nolint:errcheck

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	end := h.Subscribe(ari.Events.StasisEnd)
	defer end.Cancel()

	if err = h.Answer(); err != nil {
		return
	}

	select {
	case <-end.Events():
		return

	default:
		err = c.sf(ctx, h)
		if err != nil {
			return
		}
	}

	return
}
