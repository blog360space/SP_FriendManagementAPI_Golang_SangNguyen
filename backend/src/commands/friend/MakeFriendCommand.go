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
// Make Friend docs
// @Summary Make Friend
// @Description As a user, I need an API to create a friend connection between two email addresses
// @Accept  json
// @Produce  json
// @Param friends body []string true "Emails" collectionFormat(multi)
// @Router /friend/make-friend [post]
func MakeFriendCommand (context * gin.Context){
	var input apimodels.MakeFriendInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
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
	_, err = persistence.DbContext.Select(&users, "select Id, Username From User Where Username In (?,?)", input.Friends[0], input.Friends[1])
	for i := range input.Friends {
		var flag = true
		for j := range users {
			if input.Friends[i] == users[j].Username {
				flag = false
			}
		}
		if flag {
			var msg = fmt.Sprintf("%s isn't registered", input.Friends[i])
			output.Success = false
			output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		}
	}
	if output.Success == false {
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 5
	var userfriends []domain.UserFriendDomain
	_, err = persistence.DbContext.Select(&userfriends,"Select Id, FromUserID, ToUserID From User_Friend Where (FromUserID=? And ToUserID=?) Or (FromUserID=? And ToUserID=?)", users[0].Id, users[1].Id, users[1].Id, users[0].Id)

	if len(userfriends) > 0 {
		var msg = fmt.Sprintf("%s and %s are existed connection", input.Friends[0], input.Friends[1])
		output.Success = true
		output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		context.JSON(http.StatusBadRequest, output)
		return
	}

	userfriend := &domain.UserFriendDomain{0, users[0].Id, users[1].Id}
	persistence.DbContext.Insert(userfriend)

	context.JSON(http.StatusOK, output)
}