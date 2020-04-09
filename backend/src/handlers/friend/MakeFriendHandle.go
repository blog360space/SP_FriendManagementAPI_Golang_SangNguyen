package friend

import (
	"github.com/gin-gonic/gin"
	commands "spapp/src/commands/friend"
	"spapp/src/models/apimodels"

	helper "spapp/src/common/helpers"
	friendmodels "spapp/src/models/apimodels/friend"
)
// Make Friend docs
// @Summary Make Friend
// @Description As a user, I need an API to create a friend connection between two email addresses
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param input body friend.MakeFriendInput true "Input"
// @Success 200 {object} friend.MakeFriendOutput
// @Success 201 {object} friend.MakeFriendOutput
// @Failure 400 {object} friend.MakeFriendOutput
// @Router /friend/make-friend [post]
func MakeFriendHandle (context * gin.Context){
	var input friendmodels.MakeFriendInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = friendmodels.MakeFriendOutput{ apimodels.ApiResult{false, []string {"Input isn't null"}}}
		helper.BadRequest(context, output)
		return
	}

	var data = commands.MakeFriendCommand(input)
	helper.ApiReturn(context, data)
}