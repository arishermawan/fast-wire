package web

import (
	"fast-wire/app/handler/web/home"

	"github.com/fate-lovely/phi"
)

func Handler() *phi.Mux {
	r := phi.NewRouter()

	r.Get("/", home.Index)
	return r
}
