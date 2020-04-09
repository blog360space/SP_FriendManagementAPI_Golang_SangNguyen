package friend

import (
	"github.com/gin-gonic/gin"
	commands "spapp/src/commands/friend"
	"spapp/src/models/apimodels"

	helper "spapp/src/common/helpers"
	friendmodels "spapp/src/models/apimodels/friend"
)

// Subscribe an User docs
// @Summary Subscribe an User
// @Description As a user, I need an API to subscribe to updates from an email address.
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param input body friend.SubscribeUserInput true "Input"
// @Success 200 {object} friend.SubscribeUserOutput
// @Success 201 {object} friend.SubscribeUserOutput
// @Failure 400 {object} friend.SubscribeUserOutput
// @Router /friend/subscribe-user [post]
func SubscribeUserHandle(context *gin.Context)  {
	var input friendmodels.SubscribeUserInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = friendmodels.SubscribeUserOutput{ apimodels.ApiResult{false, []string {"Input isn't null"}} }
		helper.BadRequest(context, output)
		return
	}

	var data = commands.SubscribeUserCommand(input)
	helper.ApiReturn(context, data)
}