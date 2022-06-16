package fixtures

import (
	"github.com/ManuelFresnedaLlamas/GymTonic/appctr"
	"github.com/ManuelFresnedaLlamas/GymTonic/common"
)

var ctx = common.SystemContext("faker", common.RoleSystem)

func Make() {
	makeUsers()
	appctr.Log().Debug("finishing fixtures")
}
