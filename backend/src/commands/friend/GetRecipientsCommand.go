package friend

import (
	"fmt"
	"spapp/src/common/constants"
	helper "spapp/src/common/helpers"
	"spapp/src/models/apimodels"
	friendmodels "spapp/src/models/apimodels/friend"
	"spapp/src/models/domain"
	"spapp/src/persistence"
	"strconv"
)

func GetRecipientsCommand(input friendmodels.GetRecipientsInput) friendmodels.GetRecipientsOutput {
	// 2
	if len(input.Sender) == 0 {
		var output = friendmodels.GetRecipientsOutput{
			apimodels.ApiResult{ false, []string {"Sender isn't null or empty"}},
			[]string{}}
		return output
	}
	// 3
	if !helper.IsEmail(input.Sender) {
		var output = friendmodels.GetRecipientsOutput{
			apimodels.ApiResult{false, []string {"Sender isn't valid email address"}},
			[]string{}}
		return output
	}
	// 4
	if len(input.Text) == 0 {
		var output = friendmodels.GetRecipientsOutput{
			apimodels.ApiResult{false, []string {"Text isn't null or empty"}},
			[]string{}}

		return output
	}
	var users []domain.UserDomain
	_, _ = persistence.DbContext.Select(&users, "select Id, Username From User Where Username=?", input.Sender)
	if len(users) == 0 {
		var msg = fmt.Sprintf("%s isn't registered", input.Sender)
		var output = friendmodels.GetRecipientsOutput{
			apimodels.ApiResult{false,  []string {msg}},
			[]string{}}
		return output
	}
	var currentUser = users[0]
	// Blocked Users
	var blockedIds []int
	_,  _ = persistence.DbContext.Select(&blockedIds,"Select Requestor From Subscribe_User Where Target = ? And Status=?", currentUser.Id, constants.Blocked)

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
	_,  _ = persistence.DbContext.Select(&toFriendUserIds, query, currentUser.Id)
	// fromUserIds
	query = "Select FromUserID From User_Friend Where ToUserID=?"
	if len(blockUserIdsParam) > 0 {
		query = fmt.Sprintf("%s And FromUserID Not In (%s)", query, blockUserIdsParam)
	}
	var fromFriendUserIds []int
	_,  _ = persistence.DbContext.Select(&fromFriendUserIds, query, currentUser.Id)
	// subscribeUserIds
	var subscribeUserIds []int
	_,  _ = persistence.DbContext.Select(&subscribeUserIds,"Select Requestor From Subscribe_User Where Target = ? And Status=?", currentUser.Id, constants.Subscribed)
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
		_,  _ = persistence.DbContext.Select(&matchedUserIds,query)
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
		_,  _ = persistence.DbContext.Select(&emails,query)
	}
	output := friendmodels.GetRecipientsOutput{apimodels.ApiResult{true,  []string {}},emails}
	return output
}