package friend

import (
	"github.com/gin-gonic/gin"
	commands "spapp/src/commands/friend"
	helper "spapp/src/common/helpers"
	"spapp/src/models/apimodels"
	friendmodels "spapp/src/models/apimodels/friend"
)


// Get Friends docs
// @Summary Get Friends
// @Description As a user, I need an API to retrieve the friends list for an email address.
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param input body friend.GetFriendsInput true "Input"
// @Success 200 {object} friend.GetFriendsOutput
// @Failure 400 {object} friend.GetFriendsOutput
// @Router /friend/get-friends [post]
func GetFriendsHandle (context * gin.Context){
	var input friendmodels.GetFriendsInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = &friendmodels.GetFriendsOutput{
			apimodels.ApiResult{false, []string {"Input isn't null"}},
			0,
			[]string{}}
		helper.BadRequest(context, output)
		return
	}

	var data = commands.GetFriendsCommand(input)
	helper.ApiReturn(context, data)
}