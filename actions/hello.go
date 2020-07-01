package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

// HelloFullCycleHandler is a default handler to say Hello
// a simlpe test page.
func HelloFullCycleHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.HTML("hello.html"))
}
