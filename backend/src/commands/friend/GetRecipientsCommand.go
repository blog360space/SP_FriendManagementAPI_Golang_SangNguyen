package friend

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"spapp/src/common/constants"
	helper "spapp/src/common/helpers"
	apimodels "spapp/src/models/apimodels/friend"
	"spapp/src/models/domain"
	"spapp/src/persistence"
	"strconv"
)

// Get Recipients docs
// @Summary Get Recipients
// @Description As a user, I need an API to retrieve all email addresses that can receive updates from an email address.
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param sender body string true "Email"
// @Param text body string true "Text"
// @Success 201 {object} friend.GetRecipientsOutput
// @Failure 400 {object} friend.GetRecipientsOutput
// @Router /friend/get-recipients [post]
func GetRecipientsCommand(context *gin.Context) {
	var input apimodels.GetRecipientsInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = &apimodels.GetRecipientsOutput{false, []string {"Input isn't null"}, []string{}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 2
	if len(input.Sender) == 0 {
		var output = &apimodels.GetRecipientsOutput{false, []string {"Sender isn't null or empty"}, []string{}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 3
	if !helper.IsEmail(input.Sender) {
		var output = &apimodels.GetRecipientsOutput{false, []string {"Sender isn't valid email address"}, []string{}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 4
	if len(input.Text) == 0 {
		var output = &apimodels.GetRecipientsOutput{false, []string {"Text isn't null or empty"}, []string{}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 5
	var users []domain.UserDomain
	_, err = persistence.DbContext.Select(&users, "select Id, Username From User Where Username=?", input.Sender)
	if len(users) == 0 {
		var msg = fmt.Sprintf("%s isn't registered", input.Sender)
		var output = &apimodels.GetRecipientsOutput{false,  []string {msg},[]string{}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// Return data
	var currentUser = users[0]

	//if len(matchEmails) > 0 {

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

	// subscribeUserIds
	var subscribeUserIds []int
	_,  err = persistence.DbContext.Select(&subscribeUserIds,"Select Target From Subscribe_User Where Requestor = ? And Status=?", currentUser.Id, constants.Subscribed)

	// Extract Emails from Text
	var matchEmails = helper.ExtractEmails(input.Text)
	var matchedUserIds = []int{}
	if len(matchEmails) > 0 {
		var emailsParam = ""
		for i := range matchEmails {
			matchEmail := matchEmails[i]
			emailsParam = fmt.Sprintf("%s, '%s'",emailsParam, matchEmail)
		}
		var rune = []rune(emailsParam)
		emailsParam = string(rune[1:])

		query = fmt.Sprintf("Select Id From User Where Username In (%s) And Id != %s", emailsParam, strconv.Itoa(currentUser.Id))

		if len(blockUserIdsParam) > 0 {
			query = fmt.Sprintf("%s And Id Not In (%s)", query, blockUserIdsParam)
		}

		_,  err = persistence.DbContext.Select(&matchedUserIds,query)
	}

	var notifyUserIds = []int{}
	notifyUserIds = append(toFriendUserIds, fromFriendUserIds...)
	notifyUserIds = append(notifyUserIds, subscribeUserIds...)
	notifyUserIds = append(notifyUserIds, matchedUserIds...)

	var emails = []string {}
	if len(notifyUserIds) > 0 {
		param := ""
		for i := range notifyUserIds {
			notifyUserId := notifyUserIds[i]
			param = fmt.Sprintf("%s,%s", param, strconv.Itoa(notifyUserId))
		}
		var rune = []rune(param)
		param = string(rune[1:])
		query = fmt.Sprintf("Select Username From User Where Id In (%s)", param)
		_,  err = persistence.DbContext.Select(&emails,query)
	}

	output := &apimodels.GetRecipientsOutput{true,  []string {},emails}
	context.JSON(http.StatusOK, output)
}