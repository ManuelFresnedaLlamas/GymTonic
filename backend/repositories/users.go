package repositories

import (
	"github.com/ManuelFresnedaLlamas/GymTonic/common"
	"github.com/ManuelFresnedaLlamas/GymTonic/models"
	dbx "github.com/go-ozzo/ozzo-dbx"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"strconv"
)

func NewUsers(db *dbx.DB, log *zap.Logger) *Users {
	return &Users{
		db:  db,
		log: log,
	}
}

type Users struct {
	db  *dbx.DB
	log *zap.Logger
}

func (ur Users) ByID(q *models.UserQuery) (*models.User, error) {
	u := &models.User{}

	if err := ur.db.Select().Model(q.ID, u); err != nil {
		return nil, common.ResolveDB(err, q)
	}

	return u, nil
}

func (ur Users) ByEmail(email string) (*models.User, error) {
	u := &models.User{}

	if err := ur.db.Select().From(models.UserTable).Where(dbx.HashExp{models.UserEmail: email}).One(u); err != nil {
		return nil, common.ResolveDB(err, email)
	}

	return u, nil
}

func (ur Users) Users(q *models.UserQuery) ([]models.User, error) {
	var us []models.User

	if err := ur.db.Select().
		From(models.UserTable).
		Where(ur.query(q)).
		Limit(q.Pager.Limit).
		Offset(q.Pager.Offset).
		OrderBy(q.OrderBy()).
		All(&us); err != nil {
		return nil, common.ResolveDB(err, q)
	}

	return us, nil
}

func (ur Users) Count(q *models.UserQuery) (int64, error) {
	count := int64(0)

	if err := ur.db.Select("COUNT(*) as count").
		From(models.UserTable).
		Where(ur.query(q)).
		Row(&count); err != nil {
		return 0, common.ResolveDB(err, "", q)
	}

	return count, nil
}

func (ur Users) Create(q *models.UserQuery, u *models.User) error {
	if err := ur.db.Model(u).Insert(q.Fields...); err != nil {
		return common.ResolveDB(err, q)
	}

	return nil
}

func (ur Users) Update(q *models.UserQuery, u *models.User) error {
	if err := ur.db.Model(u).Update(q.Fields...); err != nil {
		return common.ResolveDB(err, q)
	}

	return nil
}

func (ur Users) Delete(q *models.UserQuery) error {
	if err := ur.db.Model(&models.User{ID: uuid.MustParse(q.ID)}).Delete(); err != nil {
		return common.ResolveDB(err, q)
	}

	return nil
}

func (ur Users) query(q *models.UserQuery) dbx.Expression {
	exps := make([]dbx.Expression, 0, 2)

	if len(q.IDs) > 0 {
		exps = append(exps, dbx.In(models.UserID, common.SliceStringToInterfaces(q.IDs)...))
	}

	return dbx.And(exps...)
}

func NewAuth(db *dbx.DB, log *zap.Logger) *Auth {
	return &Auth{
		db:  db,
		log: log,
	}
}

type Auth struct {
	db  *dbx.DB
	log *zap.Logger
}

func (ar Auth) ByID(q *models.AuthQuery) (*models.Auth, error) {
	u := &models.Auth{}

	if err := ar.db.Select().Model(q.ID, u); err != nil {
		return nil, common.ResolveDB(err, q)
	}

	return u, nil
}

func (ar Auth) Create(q *models.AuthQuery, u *models.Auth) error {
	if err := ar.db.Model(u).Insert(q.Fields...); err != nil {
		return common.ResolveDB(err, q)
	}

	return nil
}

func (ar Auth) Update(q *models.AuthQuery, u *models.Auth) error {
	if err := ar.db.Model(u).Update(); err != nil {
		return common.ResolveDB(err, q)
	}

	return nil
}

func (ar Auth) Delete(q *models.AuthQuery) error {
	if err := ar.db.Model(&models.Auth{UserID: q.ID}).Delete(); err != nil {
		return common.ResolveDB(err, q)
	}

	return nil
}

func (ar Auth) query(q *models.AuthQuery) dbx.Expression {
	exps := make([]dbx.Expression, 0, 5)

	if q.Token != "" {
		exps = append(exps, dbx.HashExp{models.AuthToken: q.Token})
	}

	if q.Login != "" {
		exps = append(exps, dbx.HashExp{models.AuthEmail: q.Login})
	}

	if len(q.UserIDs) > 0 {
		exps = append(exps, dbx.In(models.AuthUserID, common.SliceStringToInterfaces(q.UserIDs)...))
	}

	if q.Role > 0 {
		exps = append(exps, dbx.HashExp{models.AuthRole: q.Role})
	}

	if len(q.Roles) > 0 {
		r := make([]string, 0, len(q.Roles))
		for i := range q.Roles {
			r = append(r, strconv.Itoa(int(q.Roles[i])))
		}
		exps = append(exps, dbx.In(models.AuthRole, common.SliceStringToInterfaces(r)...))
	}

	return dbx.And(exps...)
}
