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

	// 1
	if error := context.BindJSON(&input) ; error != nil {
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

	// 4
	var users []domain.UserDomain
	if _, error := persistence.DbContext.Select(&users, "select Id, Username From User Where Username=?", input.Username); error != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg": "Db Connection refused",
		})
		return
	}
	// 5
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
	
	context.JSON(http.StatusOK, gin.H{
		"success" : true,
		"data": output,
	})
	return

}
