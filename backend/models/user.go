package models

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ManuelFresnedaLlamas/GymTonic/common"
	"github.com/google/uuid"
	"golang.org/x/crypto/argon2"
)

const (
	UserTable = "users"
	AuthTable = "auth"
	secretKey = "4lOyNjwwM6dHJjT"

	UserID        = "id"
	UserFirstName = "firstName"
	UserEmail     = "email"

	AuthEmail  = "login"
	AuthToken  = "passwordResetToken"
	AuthUserID = "userID"
	AuthRole   = "role"
)

type User struct {
	ID        uuid.UUID       `json:"id" db:"pk,id"`
	FirstName string          `json:"firstName" db:"firstName"`
	LastName  string          `json:"lastName" db:"lastName"`
	CreatedAt time.Time       `json:"createdAt" db:"createdAt"`
	Email     string          `json:"email" db:"email"`
	Phone     string          `json:"phone" db:"phone"`
	Language  string          `json:"language" db:"language"`
	Role      common.RoleType `json:"role" db:"-"`
	Disabled  bool            `json:"disabled" db:"disabled"`
	InitPass  bool            `json:"initPass" db:"initPass"`
	Clinics   []uuid.UUID     `db:"-" json:"clinics,omitempty"`
}

func (User) TableName() string {
	return UserTable
}

type AuthLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Auth struct {
	ID                 uuid.UUID       `json:"id" db:"id"`
	Login              string          `json:"login" db:"login"`
	PasswordHash       string          `json:"passwordHash" db:"passwordHash"`
	PasswordSalt       string          `json:"passwordSalt" db:"passwordSalt"`
	PasswordResetToken sql.NullString  `json:"passwordResetToken" db:"passwordResetToken"`
	Role               common.RoleType `json:"role" db:"role"`
	UserID             uuid.UUID       `json:"userID" db:"userID"`
}

type LoggedUser struct {
	ID             uuid.UUID   `json:"id"`
	FirstName      string      `json:"firstName"`
	LastName       string      `json:"lastName"`
	CreatedAt      time.Time   `json:"createdAt"`
	Email          string      `json:"email"`
	Phone          string      `json:"phone"`
	Language       string      `json:"language"`
	Disabled       bool        `json:"disabled"`
	SessionExpires time.Time   `json:"sessionExpires"`
	Clinics        []uuid.UUID `json:"clinics"`
	Permissions    []string    `json:"permissions"`
}

type RegisterUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type ChangePassword struct {
	UserID      string `json:"id"`
	Email       string `json:"email"`
	NewPassword string `json:"newPassword"`
	Token       string `json:"token"`
}

func (Auth) TableName() string {
	return AuthTable
}

func (u *Auth) SetPassword(password string) {
	u.PasswordSalt = uuid.New().String()
	u.PasswordHash = fmt.Sprintf("%x", argon2.Key([]byte(password), []byte(secretKey+u.PasswordSalt), 3, 32*1024, 4, 32))
}

func (u *Auth) ValidatePassword(password string) bool {
	hash := argon2.Key([]byte(password), []byte(secretKey+u.PasswordSalt), 3, 32*1024, 4, 32)
	return fmt.Sprintf("%x", hash) == u.PasswordHash
}

type UserQuery struct {
	ID      string          `json:"id"`
	IDs     []string        `json:"ids"`
	Email   string          `json:"email"`
	Clinics []string        `json:"clinics"`
	Role    common.RoleType `json:"role"`

	common.BaseQuery
}

type AuthQuery struct {
	ID       uuid.UUID
	Login    string
	Password string
	Role     common.RoleType
	Roles    []common.RoleType
	Token    string
	UserIDs  []string
	common.BaseQuery
}

type UserList struct {
	Items []User `json:"items"`
	Count int64  `json:"count"`
}
