package main

import (
	"github.com/ManuelFresnedaLlamas/GymTonic/appctr"
	"github.com/ManuelFresnedaLlamas/GymTonic/handlers"
	"github.com/ManuelFresnedaLlamas/GymTonic/repositories"
	"github.com/ManuelFresnedaLlamas/GymTonic/services"
)

func prepareIoC() {
	rs := prepareRepositories()
	ss := prepareServices(rs)
	hs := prepareHandlers(ss)
	prepareRouter(hs)
}

type rprs struct {
	users *repositories.Users
	auth  *repositories.Auth
}

type srvs struct {
	users *services.Users
	auth  *services.Auth
}

type hdls struct {
	users *handlers.Users
	auth  *handlers.Auth
}

func prepareRepositories() *rprs {
	db := appctr.DB()
	log := appctr.Log()

	return &rprs{
		users: repositories.NewUsers(db, log),
		auth:  repositories.NewAuth(db, log),
	}
}

func prepareServices(rs *rprs) *srvs {
	log := appctr.Log()

	return &srvs{
		users: services.NewUsers(rs.users, rs.auth, log),
		auth:  services.NewAuth(rs.auth, rs.users, log),
	}
}

func prepareHandlers(ss *srvs) *hdls {
	log := appctr.Log()

	return &hdls{
		users: handlers.NewUsers(ss.users, log),
		auth:  handlers.NewAuth(ss.auth, ss.users, log),
	}
}

func prepareRouter(hs *hdls) {
	//appctr.AddAuthMiddleware(hs.auth.AuthMiddleware)

	r := appctr.RouteGroup("/users")
	r.Get("/<id>", hs.users.ByID)
	r.Post("", hs.users.Create)
	r.Put("", hs.users.Update)

}
