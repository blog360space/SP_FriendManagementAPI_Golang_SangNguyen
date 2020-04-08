package friend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	apimodels "spapp/src/models/apimodels/friend"
	helper "spapp/src/common/helpers"

	"spapp/src/models/domain"
	"spapp/src/persistence"
)

func GetFriendsCommand (context * gin.Context){
	var input apimodels.GetFriendsInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = &apimodels.GetFriendsOutput{false, []string {"Input isn't null"}, 0, []string{}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 2
	if len(input.Email) == 0 {
		var output = &apimodels.GetFriendsOutput{false, []string {"Email isn't null or empty"},0, []string{}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 3
	if !helper.IsEmail(input.Email) {
		var msg = fmt.Sprintf("%s isn't an Email address", input.Email)
		var output = &apimodels.GetFriendsOutput{false, []string {msg},0, []string{}}
		output.Msgs = helper.AddItemToArray(output.Msgs, msg)
	}

	// 4
	var users []domain.UserDomain
	_, err = persistence.DbContext.Select(&users, "select Id, Username From User Where Username=?", input.Email)
	if len(users) == 0 {
		var msg = fmt.Sprintf("%s isn't registered", input.Email)
		var output = &apimodels.GetFriendsOutput{false,  []string {msg},0, []string{}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// TODO: Block User

	// Get Friends
	var user = users[0]
	var fromUserIds []int
	_, err = persistence.DbContext.Select(&fromUserIds,"Select FromUserID From User_Friend Where ToUserID=?", user.Id)
	var toUserIds []int
	_, err = persistence.DbContext.Select(&toUserIds,"Select ToUserID From User_Friend Where FromUserID=?", user.Id)

	var userIds = append(fromUserIds, toUserIds...)

	// Build Query
	var param = ""
	for i := range userIds {
		userId := userIds[i]
		param = param + "," + strconv.Itoa(userId)
	}
	var emails []string
	if len(param) > 0 {
		var rune = []rune(param)
		params := string(rune[1:])
		query := fmt.Sprintf("Select Username From User Where Id In (%s) And Id != %s", params, strconv.Itoa(user.Id))
		_,  err = persistence.DbContext.Select(&emails,query)
	}
	var output = &apimodels.GetFriendsOutput{true, []string {}, len(emails), emails}
	context.JSON(http.StatusOK, output)
}