package friend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"

	apimodels "spapp/src/models/apimodels/friend"
	helper "spapp/src/common/helpers"

	"spapp/src/models/domain"
	"spapp/src/persistence"
)

func MakeFriendCommand (context * gin.Context){
	var input apimodels.MakeFriendInput

	// 1
	if error := context.BindJSON(&input) ; error != nil {
		var output = &apimodels.MakeFriendOutput{false, []string {"Input isn't null"}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 2
	if helper.IsNull(input.Friends) || len(input.Friends) < 2 {
		var output = &apimodels.MakeFriendOutput{false, []string {"Input isn't valid"} }
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 3
	var count = len(input.Friends)
	var output = &apimodels.MakeFriendOutput{true, []string {}}
	for i := 0; i < count; i++ {
		if !helper.IsEmail(input.Friends[i]) {
			output.Success = false
			var msg string = fmt.Sprintf("%s isn't an Email address", input.Friends[i])
			output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		}
	}
	if output.Success == false {
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 4
	var users []domain.UserDomain
	if _, error := persistence.DbContext.Select(&users, "select Id, Username From User Where Username In (?,?)", input.Friends[0], input.Friends[1]); error != nil{
		context.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"msg": "Db Connection refused",
		})
		return
	}
	for i := range input.Friends {
		var flag = true
		for j := range users {
			if input.Friends[i] == users[j].Username {
				flag = false
			}
		}
		if flag {
			output.Success = false
			var msg = fmt.Sprintf("%s isn't registered", input.Friends[i])
			output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		}
	}
	if output.Success == false {
		context.JSON(http.StatusBadRequest, output)
		return
	}

	userfriend := &domain.UserFriendDomain{0, users[0].Id, users[1].Id}
	persistence.DbContext.Insert(userfriend)
	context.JSON(http.StatusOK, output)
}