package common

import (
	"database/sql"
	"fmt"
)

const (
	FormatLog = `{"message": "%v", "data": "%#v"}`
)

type BaseError struct {
	Message     string        `json:"message,omitempty"`
	Data        []interface{} `json:"data, omitempty"`
	Description string
}

func (be *BaseError) Error() string {
	return be.Message
}

func (be *BaseError) Log() string {
	return fmt.Sprintf(FormatLog, be.Message, be.Data)
}

func ResolveDB(err error, data ...interface{}) error {
	if err == sql.ErrNoRows {
		return NewNotFound(err.Error(), data)
	}

	return NewDB(err.Error(), data)
}

type NotFound struct {
	BaseError
}

func NewNotFound(message string, data ...interface{}) *NotFound {
	return &NotFound{BaseError{
		Message: message,
		Data:    data,
	}}
}

type DB struct {
	BaseError
}

func NewDB(message string, data ...interface{}) *DB {
	return &DB{BaseError{
		Message: message,
		Data:    data,
	}}
}

type Unauthorized struct {
	BaseError
}

func NewUnauthorized(message string, data ...interface{}) *Unauthorized {
	return &Unauthorized{BaseError{
		Message: message,
		Data:    data,
	}}
}

type Forbidden struct {
	BaseError
}

func NewForbidden(message string, data ...interface{}) *Forbidden {
	return &Forbidden{BaseError{
		Message: message,
		Data:    data,
	}}
}

type BadRequest struct {
	BaseError
}

func NewBadRequest(message string, data ...interface{}) *BadRequest {
	if message == "" {
		message = "bad request"
	}
	return &BadRequest{BaseError{
		Message: message,
		Data:    data,
	}}
}
