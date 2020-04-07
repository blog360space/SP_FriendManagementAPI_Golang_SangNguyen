package user

import (
	"github.com/gin-gonic/gin"
	"net/http"

	helper "spapp/src/common/helpers"
	apimodels "spapp/src/models/apimodels/user"
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
	//user := persistence.DbContext.SelectOne("")
	if !helper.IsEmail(input.Username) {
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg": "Username isn't valid email address",
		})
		return
	}


	//user := &domain.UserDomain{ 0, input.Username }
	//db.DbContext.Insert(user)
	context.JSON(http.StatusOK, gin.H{
		//"Id": user.ID,
		//"Username": user.Username,
	})
	return

}
