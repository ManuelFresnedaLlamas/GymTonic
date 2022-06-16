package handlers

import (
	"github.com/ManuelFresnedaLlamas/GymTonic/common"
	"github.com/ManuelFresnedaLlamas/GymTonic/models"
	"github.com/ManuelFresnedaLlamas/GymTonic/services"
	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

func NewUsers(us *services.Users, log *zap.Logger) *Users {
	return &Users{
		us:  us,
		log: log,
	}
}

type Users struct {
	us  *services.Users
	log *zap.Logger
}

func (uh Users) ByID(ctx *routing.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return common.NewBadRequest("", "")
	}

	u, err := uh.us.ByID(common.NewAppContext(ctx), &models.UserQuery{ID: id.String()})
	if err != nil {
		return err
	}

	if err = ctx.Write(u); err != nil {
		uh.log.Error(err.Error())
	}

	return nil
}

func (uh Users) Create(ctx *routing.Context) error {
	q, err := uh.parseQuery(ctx)
	if err != nil {
		return err
	}

	u := &models.User{}
	if err := ctx.Read(u); err != nil {
		return err
	}

	if err := uh.us.Create(common.NewAppContext(ctx), q, u); err != nil {
		return err
	}

	if err := ctx.Write(u); err != nil {
		uh.log.Error(err.Error())
	}

	return nil
}

func (uh Users) Update(ctx *routing.Context) error {
	q, err := uh.parseQuery(ctx)
	if err != nil {
		return err
	}

	u := &models.User{}
	if err := ctx.Read(u); err != nil {
		return err
	}

	if err := uh.us.Update(common.NewAppContext(ctx), q, u); err != nil {
		return err
	}

	if err := ctx.Write(u); err != nil {
		uh.log.Error(err.Error())
	}

	return nil
}

func (uh Users) Delete(ctx *routing.Context) error {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		return common.NewBadRequest("", "")
	}

	if err := uh.us.Delete(common.NewAppContext(ctx), &models.UserQuery{ID: id.String()}); err != nil {
		return err
	}

	ctx.Response.WriteHeader(http.StatusNoContent)

	return nil
}

func (uh Users) parseQuery(ctx *routing.Context) (*models.UserQuery, error) {
	q := models.UserQuery{}
	if err := q.ParseBase(ctx); err != nil {
		return nil, err
	}
	if err := common.ParseQuery(ctx, &q); err != nil {
		return nil, err
	}

	if len(q.Sorts) == 0 {
		q.Sorts = append(q.Sorts, common.Sort{
			Field: models.UserFirstName,
		})
	}

	return &q, nil
}

const (
	sid                 = "SID"
	loginPage           = "/login"
	registerPage        = "/register"
	recoverPasswordPage = "/recover-password"
	changePasswordPage  = "/change-password"
	filesPage           = "/files"
)

func NewAuth(as *services.Auth, us *services.Users, log *zap.Logger) *Auth {
	return &Auth{
		as:  as,
		us:  us,
		log: log,
	}
}

type Auth struct {
	as  *services.Auth
	us  *services.Users
	log *zap.Logger
}

/*
func (ah Auth) Update(ctx *routing.Context) error {
	c, err := ctx.Request.Cookie(sid)
	if err != nil {
		return nil
	}
	s, err := ah.as.UserIDBySessionID(common.NewAppContext(ctx), c.Value)
	if err != nil {
		return err
	}

	u, err := ah.us.ByID(common.NewAppContext(ctx), &models.UserQuery{ID: s.UserID.String()})
	if err != nil {
		return err
	}

	lu := models.LoggedUser{
		ID:             u.ID,
		FirstName:      u.FirstName,
		LastName:       u.LastName,
		Email:          u.Email,
		Phone:          u.Phone,
		Disabled:       u.Disabled,
		Language:       u.Language,
		SessionExpires: s.ExpiredAt,
		Clinics:        u.Clinics,
		Permissions:    services.GetPermissionsByRole(s.Role),
	}

	if err = ctx.Write(lu); err != nil {
		ah.log.Error(err.Error())
	}

	return nil
}



func (ah Auth) AuthMiddleware(ctx *routing.Context) error {
	c, err := ctx.Request.Cookie(sid)
	url := ctx.Request.URL.Path
	s := &models.Session{
		ID:        "",
		UserID:    uuid.UUID{},
		ExpiredAt: time.Time{},
		Role:      common.RoleDisabled,
	}

	if err == nil {
		s, err = ah.as.UserIDBySessionID(common.NewAppContext(ctx), c.Value)
		if _, ok := err.(*common.NotFound); err != nil && !ok {
			ah.log.Error(err.Error())
			return err
		}
	}

	switch true {
	case ah.isLogged(s.UserID, url):
		ctx.Set(common.UserID, s.UserID.String())
		ctx.Set(common.Role, s.Role)
		return ctx.Next()
	case ah.wantLogin(s.UserID, url):
		return ctx.Next()
	case ah.wantRegister(s.UserID, url):
		return ctx.Next()
	case ah.wantRecoverPassword(s.UserID, url):
		return ctx.Next()
	case ah.wantChangePassword(s.UserID, url):
		return ctx.Next()
	case ah.wantLoadFile(s.UserID, url):
		return ctx.Next()
	case ah.isLoggedAndWantLogin(s.UserID, url):
		ctx.Abort()
		return nil
	default:
		ctx.Abort()
		return common.NewUnauthorized("", "", s.UserID, url)
	}

}
*/

func (ah Auth) isLogged(uID uuid.UUID, url string) bool {
	return uID != uuid.Nil && url != loginPage
}

func (ah Auth) wantLogin(uID uuid.UUID, url string) bool {
	return uID == uuid.Nil && url == loginPage
}

func (ah Auth) isLoggedAndWantLogin(uID uuid.UUID, url string) bool {
	return uID != uuid.Nil && url == loginPage
}

func (ah Auth) wantRegister(uID uuid.UUID, url string) bool {
	return uID == uuid.Nil && url == registerPage
}

func (ah Auth) wantRecoverPassword(uID uuid.UUID, url string) bool {
	return uID == uuid.Nil && url == recoverPasswordPage
}

func (ah Auth) wantChangePassword(uID uuid.UUID, url string) bool {
	return uID == uuid.Nil && strings.Contains(url, changePasswordPage)
}

func (ah Auth) wantLoadFile(uID uuid.UUID, url string) bool {
	return uID == uuid.Nil && url == filesPage
}
