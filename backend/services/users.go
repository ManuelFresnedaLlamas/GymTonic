package services

import (
	"database/sql"
	"github.com/ManuelFresnedaLlamas/GymTonic/common"
	"github.com/ManuelFresnedaLlamas/GymTonic/models"
	"github.com/ManuelFresnedaLlamas/GymTonic/repositories"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func NewUsers(ur *repositories.Users, ar *repositories.Auth, log *zap.Logger) *Users {
	return &Users{ur: ur, ar: ar, log: log}
}

type Users struct {
	ur  *repositories.Users
	ar  *repositories.Auth
	log *zap.Logger
}

func (us Users) ByID(ctx *common.AppContext, q *models.UserQuery) (*models.User, error) {
	u, err := us.ur.ByID(q)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (us Users) Create(ctx *common.AppContext, q *models.UserQuery, u *models.User) error {

	if err := us.ur.Create(q, u); err != nil {
		return err
	}

	auth := &models.Auth{
		ID:                 uuid.New(),
		Login:              u.Email,
		PasswordResetToken: sql.NullString{},
		Role:               u.Role,
		UserID:             u.ID,
	}

	auth.SetPassword("12345678Aa")

	if err := us.ar.Create(&models.AuthQuery{}, auth); err != nil {
		return err
	}

	return nil
}

func (us Users) Update(ctx *common.AppContext, q *models.UserQuery, u *models.User) error {

	return us.ur.Update(q, u)
}

func (us Users) Delete(ctx *common.AppContext, q *models.UserQuery) error {
	return us.ur.Delete(q)
}

func NewAuth(ar *repositories.Auth, ur *repositories.Users, log *zap.Logger) *Auth {
	return &Auth{ar: ar, ur: ur, log: log}
}

type Auth struct {
	ar  *repositories.Auth
	ur  *repositories.Users
	log *zap.Logger
}

func (as Auth) ByID(ctx *common.AppContext, q *models.AuthQuery) (*models.Auth, error) {
	return as.ar.ByID(q)
}

func (as Auth) Create(ctx *common.AppContext, q *models.AuthQuery, u *models.Auth) error {
	return as.ar.Create(q, u)
}

func (as Auth) Update(ctx *common.AppContext, q *models.AuthQuery, u *models.Auth) error {
	return as.ar.Update(q, u)
}

func (as Auth) Delete(ctx *common.AppContext, q *models.AuthQuery) error {
	return as.ar.Delete(q)
}
