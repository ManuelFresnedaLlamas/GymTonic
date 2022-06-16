package common

type RoleType int

const (
	RoleDisabled RoleType = iota
	RoleGymAdmin
	RoleGymUser
	RoleSystem = 100
)

const (
	ViewUsers    = "ViewUsers"
	EditAllUsers = "EditAllUsers"
	EditUsers    = "EditUsers"
	DeleteUsers  = "DeleteUsers"
)

type GenderType int

const (
	GenderMale GenderType = iota + 1
	GenderFemale
)
