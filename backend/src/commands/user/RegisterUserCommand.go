package user

import (
	"github.com/gin-gonic/gin"
	"net/http"

	helper "spapp/src/common/helpers"
	apimodels "spapp/src/models/apimodels/user"
	"spapp/src/models/domain"
	"spapp/src/persistence"
)

func RegisterUserCommand (context * gin.Context){
	var input apimodels.RegisterUserInput
	var err error
	// 1
	if err = context.BindJSON(&input) ; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg": "Model isn't null",
		})
		return
	}

	// 2
	if len(input.Username) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg": "Username isn't null or empty",
		})
		return
	}

	// 3
	if !helper.IsEmail(input.Username) {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg": "Username isn't valid email address",
		})
		return
	}

	// 5
	var users []domain.UserDomain
	_, err = persistence.DbContext.Select(users,"select Id, Username From User Where Username=?", input.Username)

	if len(users) > 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg": "Username is existed",
		})
		return
	}

	user := &domain.UserDomain{ 0, input.Username }
	persistence.DbContext.Insert(user)
	output := &apimodels.RegisterUserOutput{user.Id, user.Username}
	
	context.JSON(http.StatusCreated, gin.H{
		"success" : true,
		"data": output,
	})
	return

}
