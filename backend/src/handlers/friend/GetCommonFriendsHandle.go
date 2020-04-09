package friend

import (
	"github.com/gin-gonic/gin"
	commands "spapp/src/commands/friend"
	helper "spapp/src/common/helpers"
	"spapp/src/models/apimodels"
	friendmodels "spapp/src/models/apimodels/friend"
)


// Get Common Friends docs
// @Summary Get Common Friends
// @Description As a user, I need an API to retrieve the common friends list between two email addresses.
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param input body friend.GetCommonFriendsInput true "Input"
// @Success 200 {object} friend.GetCommonFriendsOutput
// @Failure 400 {object} friend.GetCommonFriendsOutput
// @Router /friend/get-common-friends [post]
func GetCommonFriendsHandle(context *gin.Context) {
	var input friendmodels.GetCommonFriendsInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = friendmodels.GetCommonFriendsOutput{
			apimodels.ApiResult{false,  []string {"Input isn't null"}},
			0,
			[]string{}}
		helper.BadRequest(context, output)
		return
	}

	var data = commands.GetCommonFriendsCommand(input)
	helper.ApiReturn(context, data)
}