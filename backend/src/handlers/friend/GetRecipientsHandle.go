package friend

import (
	"github.com/gin-gonic/gin"
	commands "spapp/src/commands/friend"
	helper "spapp/src/common/helpers"
	"spapp/src/models/apimodels"
	friendmodels "spapp/src/models/apimodels/friend"
)

// Get Recipients docs
// @Summary Get Recipients
// @Description As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param input body friend.GetRecipientsInput true "Input"
// @Success 200 {object} friend.GetRecipientsOutput
// @Failure 400 {object} friend.GetRecipientsOutput
// @Router /friend/get-recipients [post]
func GetRecipientsHandle(context *gin.Context) {
	var input friendmodels.GetRecipientsInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = friendmodels.GetRecipientsOutput{
			apimodels.ApiResult{false, []string {"Input isn't null"}},
			[]string{}}
		helper.BadRequest(context, output)
		return
	}

	var data = commands.GetRecipientsCommand(input)
	helper.ApiReturn(context, data)
}