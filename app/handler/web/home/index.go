package home

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	ctx.SetContentType("text/html")
	fmt.Fprint(ctx, "Welcome to Home Page!\n")
}
