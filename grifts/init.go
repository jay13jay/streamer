package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/jay13jay/streamer/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
