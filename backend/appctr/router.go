package appctr

import (
	"time"

	"github.com/go-ozzo/ozzo-routing/content"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/cors"
)

func Router() *routing.Router {
	return r
}

func RouteGroup(module string) *routing.RouteGroup {
	return r.Group(module)
}

var (
	r                           = routing.New()
	authHandler routing.Handler = nil
)

func AddAuthMiddleware(h routing.Handler) {
	authHandler = h
}

func UseMiddlewares() {
	r.Use(
		logMiddleware,
		cors.Handler(cors.Options{
			AllowOrigins:     "*",
			AllowHeaders:     "*",
			AllowMethods:     "*",
			AllowCredentials: true,
		}),
		authMiddleware,
		content.TypeNegotiator(content.JSON),
	)
}

func logMiddleware(ctx *routing.Context) error {
	t1 := time.Now()
	err := ctx.Next()
	lg.Debug(time.Since(t1).String())

	if err != nil {
		lg.Debug(err.Error())
		lg.Error(err.Error())
	}

	return err
}

func authMiddleware(ctx *routing.Context) error {
	return authHandler(ctx)
}
