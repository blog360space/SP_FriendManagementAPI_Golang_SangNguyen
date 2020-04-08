package user

import (
	"github.com/gin-gonic/gin"
	"net/http"

	helper "spapp/src/common/helpers"
	apimodels "spapp/src/models/apimodels/user"
	"spapp/src/models/domain"
	"spapp/src/persistence"
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
func RegisterUserCommand (context * gin.Context){
	var input apimodels.RegisterUserInput
	var err error
	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output apimodels.RegisterUserOutput
		output.Success = false
		output.Msgs = []string{"Input isn't null"}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 2
	if len(input.Username) == 0 {
		var output apimodels.RegisterUserOutput
		output.Success = false
		output.Msgs = []string{"Username isn't null or empty"}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 3
	if !helper.IsEmail(input.Username) {
		var output apimodels.RegisterUserOutput
		output.Success = false
		output.Msgs = []string{"Username isn't valid email address"}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 5
	var users []domain.UserDomain
	_, err = persistence.DbContext.Select(users,"select Id, Username From User Where Username=?", input.Username)

	if len(users) > 0 {
		var output apimodels.RegisterUserOutput
		output.Success = false
		output.Msgs = []string{"Username is existed"}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	user := &domain.UserDomain{ 0, input.Username }
	persistence.DbContext.Insert(user)
	var output apimodels.RegisterUserOutput
	output.Success = true
	output.Data.Id = user.Id
	output.Data.Username =  user.Username
	context.JSON(http.StatusCreated, output)
	return

}
