package friend

import (
	"github.com/gin-gonic/gin"
	commands "spapp/src/commands/friend"
	helper "spapp/src/common/helpers"
	"spapp/src/models/apimodels"
	friendmodels "spapp/src/models/apimodels/friend"
)

// Block an User docs
// @Summary Block an User
// @Description As a user, I need an API to block updates from an email address.
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param input body friend.BlockUserInput true "Input"
// @Success 200 {object} friend.BlockUserOutput
// @Success 201 {object} friend.BlockUserOutput
// @Failure 400 {object} friend.BlockUserOutput
// @Router /friend/block-user [post]
func BlockUserHandle(context *gin.Context)  {
	var input friendmodels.BlockUserInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = &friendmodels.BlockUserOutput{apimodels.ApiResult{false, []string {"Input isn't null"}}}
		helper.BadRequest(context, output)
		return
	}

	var data = commands.BlockUserCommand(input)
	helper.ApiReturn(context, data)
}