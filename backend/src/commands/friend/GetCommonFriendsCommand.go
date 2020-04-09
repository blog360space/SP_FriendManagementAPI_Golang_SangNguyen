package friend

import (
	"fmt"
	"spapp/src/common/constants"
	"spapp/src/models/apimodels"
	"strconv"

	helper "spapp/src/common/helpers"
	friendmodels "spapp/src/models/apimodels/friend"

	"spapp/src/models/domain"
	"spapp/src/persistence"
)

func GetCommonFriendsCommand(input friendmodels.GetCommonFriendsInput) friendmodels.GetCommonFriendsOutput {
	// 2
	if helper.IsNull(input.Friends) || len(input.Friends) < 2 {
		var output = friendmodels.GetCommonFriendsOutput{
			apimodels.ApiResult{false, []string {"Input isn't valid"}},
			0,
			[]string{} }
		return output
	}
	// 3
	var count = len(input.Friends)
	var output = friendmodels.GetCommonFriendsOutput{
		apimodels.ApiResult{true, []string {}},
		0,
		[]string{}}
	for i := 0; i < count; i++ {
		if !helper.IsEmail(input.Friends[i]) {
			output.Success = false
			var msg string = fmt.Sprintf("%s isn't an Email address", input.Friends[i])
			output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		}
	}
	if output.Success == false {
		return output
	}
	// 4
	if input.Friends[0] == input.Friends[1] {
		output.Success = false
		output.Msgs = helper.AddItemToArray(output.Msgs, "Emails are the same")
		return output
	}
	// 5
	var users []domain.UserDomain
	_, _ = persistence.DbContext.Select(&users, "select Id, Username From User Where Username In (?,?)", input.Friends[0], input.Friends[1])
	for i := range input.Friends {
		var flag = true
		for j := range users {
			if input.Friends[i] == users[j].Username {
				flag = false
			}
		}
		if flag {
			var msg = fmt.Sprintf("%s isn't registered", input.Friends[i])
			output.Success = false
			output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		}
	}
	if output.Success == false {
		return output
	}
	var user1 = users[0]
	var user2 = users[1]
	// Blocked Users From User 1
	var blockedIds []int
	_,  _ = persistence.DbContext.Select(&blockedIds,"Select Target  From Subscribe_User Where Requestor = ? And Status=?", user1.Id, constants.Blocked)
	var blockUserIdsParam = ""
	if len(blockedIds) > 0 {
		for i := range blockedIds {
			blockedId := blockedIds[i]
			blockUserIdsParam = blockUserIdsParam + "," + strconv.Itoa(blockedId)
		}
		var rune = []rune(blockUserIdsParam)
		blockUserIdsParam = string(rune[1:])
	}
	var query = "Select ToUserID From User_Friend Where FromUserID =? And ToUserID != ?"
	if len(blockUserIdsParam) > 0 {
		query = fmt.Sprintf("%s And ToUserID Not In (%s)", query, blockUserIdsParam)
	}
	var toUserIds1 []int
	_, _ = persistence.DbContext.Select(&toUserIds1, query, user1.Id, user2.Id)

	query = "Select FromUserID From User_Friend Where ToUserID = ? And FromUserID != ?"
	if len(blockUserIdsParam) > 0 {
		query = fmt.Sprintf("%s And FromUserID Not In (%s)", query, blockUserIdsParam)
	}
	var fromUserIds1 []int
	_, _ = persistence.DbContext.Select(&fromUserIds1,query, user1.Id, user2.Id)

	var userIds1 = append(fromUserIds1, toUserIds1...)

	// Blocked Users From User 2
	_,  _ = persistence.DbContext.Select(&blockedIds,"Select Target From Subscribe_User Where Requestor = ? And Status=?", user2.Id, constants.Blocked)
	blockUserIdsParam = ""
	if len(blockedIds) > 0 {
		for i := range blockedIds {
			blockedId := blockedIds[i]
			blockUserIdsParam = blockUserIdsParam + "," + strconv.Itoa(blockedId)
		}
		var rune = []rune(blockUserIdsParam)
		blockUserIdsParam = string(rune[1:])
	}
	query = "Select ToUserID From User_Friend Where FromUserID =? And ToUserID != ?"
	if len(blockUserIdsParam) > 0 {
		query = fmt.Sprintf("%s And ToUserID Not In (%s)", query, blockUserIdsParam)
	}
	var toUserIds2 []int
	_, _ = persistence.DbContext.Select(&toUserIds2, query, user2.Id, user1.Id)

	query = "Select FromUserID From User_Friend Where ToUserID = ? And FromUserID != ?"
	if len(blockUserIdsParam) > 0 {
		query = fmt.Sprintf("%s And FromUserID Not In (%s)", query, blockUserIdsParam)
	}
	var fromUserIds2 []int
	_, _ = persistence.DbContext.Select(&fromUserIds2,query, user2.Id, user1.Id)

	var userIds2 = append(fromUserIds2, toUserIds2...)

	// Search Intersection Users
	var userIds = helper.Intersection(userIds1, userIds2)

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
		query := fmt.Sprintf("Select Username From User Where Id In (%s)", params)


		_,  _ = persistence.DbContext.Select(&emails,query)
		output = friendmodels.GetCommonFriendsOutput{
			apimodels.ApiResult{true, []string {}},
			len(emails),
			emails}
	}
	return output
}