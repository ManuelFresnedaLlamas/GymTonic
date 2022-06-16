package fixtures

import (
	"github.com/ManuelFresnedaLlamas/GymTonic/appctr"
	"github.com/ManuelFresnedaLlamas/GymTonic/common"
	"github.com/ManuelFresnedaLlamas/GymTonic/models"
	"github.com/ManuelFresnedaLlamas/GymTonic/repositories"
	"github.com/ManuelFresnedaLlamas/GymTonic/services"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

const (
	gymAdmin = "cdb94efa-506e-4a94-86de-0bdb8609e733"
	gymUser  = "4bd50739-9612-463b-bce7-17ade0fc35be"

	gymAdminEmail = "admin@gymtonic.es"
	gymUserEmail  = "user@gymtonic.es"
)

type usersFixtures struct {
	us *services.Users
	as *services.Auth

	log *zap.Logger

	users map[string]models.User
}

var usrFxtr = usersFixtures{}

func makeUsers() {
	ur := repositories.NewUsers(appctr.DB(), appctr.Log())
	ar := repositories.NewAuth(appctr.DB(), appctr.Log())
	us := services.NewUsers(ur, ar, appctr.Log())
	as := services.NewAuth(ar, ur, appctr.Log())

	usrFxtr.us = us
	usrFxtr.as = as

	usrFxtr.make()

}

func (uf *usersFixtures) make() {
	uf.create()
}

func (uf *usersFixtures) create() {
	type user struct {
		id        string
		firstName string
		lastName  string
		email     string
		role      common.RoleType
	}

	ul := []user{
		{
			id:        gymAdmin,
			firstName: "Gym",
			lastName:  "Admin",
			email:     gymAdminEmail,
			role:      common.RoleGymAdmin,
		},
		{
			id:        gymUser,
			firstName: "Gym",
			lastName:  "User",
			email:     gymUserEmail,
			role:      common.RoleGymUser,
		},
	}

	for i := range ul {
		u := &models.User{
			ID:        uuid.MustParse(ul[i].id),
			FirstName: ul[i].firstName,
			LastName:  ul[i].lastName,
			CreatedAt: time.Now(),
			Email:     ul[i].email,
			Phone:     "+34111111111",
			Language:  models.Spanish,
			Role:      ul[i].role,
			InitPass:  false,
		}

		if err := uf.us.Create(&ctx, &models.UserQuery{}, u); err != nil {
			uf.log.Fatal(err.Error())
		}
	}
}
