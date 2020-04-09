package user

import (
	"github.com/gin-gonic/gin"
	commands "spapp/src/commands/user"
	helper "spapp/src/common/helpers"
	usermodels "spapp/src/models/apimodels/user"
)

// Register an User docs
// @Summary Register an User
// @Description As a user, I need an API to create an user by email address
// @Tags User
// @Accept  json
// @Produce  json
// @Param input body user.RegisterUserInput true "Email"
// @Success 201 {object} user.RegisterUserOutput
// @Failure 400 {object} user.RegisterUserOutput
// @Router /user/register-user [post]
func RegisterUserHandle (context * gin.Context){
	var input usermodels.RegisterUserInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output usermodels.RegisterUserOutput
		output.Success = false
		output.Msgs = []string{"Input isn't null"}
		helper.BadRequest(context, output)
		return
	}

	var data = commands.RegisterUserCommand(input)
	helper.ApiReturn(context, data)
}
