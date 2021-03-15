package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/jay13jay/livepeer/create"
	)

// HomeHandler is a default handler to serve up
// a home page.
func CreateHandler(c buffalo.Context) error {
	resp := create.createStream()
	return c.Render(http.StatusOK, r.HTML("index.html"))
}

