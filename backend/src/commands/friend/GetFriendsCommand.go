package friend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"spapp/src/common/constants"
	"strconv"

	apimodels "spapp/src/models/apimodels/friend"
	helper "spapp/src/common/helpers"

	"spapp/src/models/domain"
	"spapp/src/persistence"
)


// Get Friends docs
// @Summary Get Friends
// @Description As a user, I need an API to retrieve the friends list for an email address.
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param email body string true "Email"
// @Success 201 {object} friend.GetFriendsOutput
// @Failure 400 {object} friend.GetFriendsOutput
// @Router /friend/get-friends [post]
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
	// Return Data
	var currentUser = users[0]

	// Blocked Users
	var blockedIds []int
	_,  err = persistence.DbContext.Select(&blockedIds,"Select Target From Subscribe_User Where Requestor = ? And Status=?", currentUser.Id, constants.Blocked)
	var blockUserIdsParam = ""
	if len(blockedIds) > 0 {
		for i := range blockedIds {
			blockedId := blockedIds[i]
			blockUserIdsParam = blockUserIdsParam + "," + strconv.Itoa(blockedId)
		}
		var rune = []rune(blockUserIdsParam)
		blockUserIdsParam = string(rune[1:])
	}

	// toUserIds
	var query = "Select ToUserID From User_Friend Where FromUserID=?"
	if len(blockUserIdsParam) > 0 {
		query = fmt.Sprintf("%s And ToUserID Not In (%s)", query, blockUserIdsParam)
	}
	var toFriendUserIds []int
	_,  err = persistence.DbContext.Select(&toFriendUserIds, query, currentUser.Id)

	// fromUserIds
	query = "Select FromUserID From User_Friend Where ToUserID=?"
	if len(blockUserIdsParam) > 0 {
		query = fmt.Sprintf("%s And FromUserID Not In (%s)", query, blockUserIdsParam)
	}
	var fromFriendUserIds []int
	_,  err = persistence.DbContext.Select(&fromFriendUserIds, query, currentUser.Id)


	var userIds = append(toFriendUserIds, fromFriendUserIds...)

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
		query = fmt.Sprintf("Select Username From User Where Id In (%s) And Id != %s", params, strconv.Itoa(currentUser.Id))
		_,  err = persistence.DbContext.Select(&emails,query)
	}
	var output = &apimodels.GetFriendsOutput{true, []string {}, len(emails), emails}
	context.JSON(http.StatusOK, output)
}