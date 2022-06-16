package common

import (
	routing "github.com/go-ozzo/ozzo-routing"
)

const (
	UserID   = "app-user-id"
	ActionID = "app-action-id"
	Role     = "app-role"
)

type AppContext struct {
	UserID   string
	ActionID string
	Role     RoleType
}

func NewAppContext(ctx *routing.Context) *AppContext {
	var uID string
	var rID string
	var role RoleType

	if id, ok := ctx.Get(UserID).(string); ok {
		uID = id
	}

	if id, ok := ctx.Get(ActionID).(string); ok {
		rID = id
	}

	if r, ok := ctx.Get(Role).(RoleType); ok {
		role = r
	}

	return &AppContext{UserID: uID, ActionID: rID, Role: role}
}
