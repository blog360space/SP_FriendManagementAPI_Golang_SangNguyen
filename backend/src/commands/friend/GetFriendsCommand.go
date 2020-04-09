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


func GetFriendsCommand (input friendmodels.GetFriendsInput) friendmodels.GetFriendsOutput{

	// 1
	if helper.IsNull(input) {
		var output = friendmodels.GetFriendsOutput{
			apimodels.ApiResult{false, []string {"Input isn't null"}},
			0,
			[]string{}}

		return output
	}

	// 2
	if len(input.Email) == 0 {
		var output = friendmodels.GetFriendsOutput{
			apimodels.ApiResult{false, []string {"Email isn't null or empty"}},
			0,
			[]string{}}
		return output
	}

	// 3
	if !helper.IsEmail(input.Email) {
		var msg = fmt.Sprintf("%s isn't an Email address", input.Email)
		var output = friendmodels.GetFriendsOutput{
			apimodels.ApiResult{false, []string {msg}} ,
			0,
			[]string{}}
		output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		return output
	}

	// 4
	var users []domain.UserDomain
	_, _ = persistence.DbContext.Select(&users, "select Id, Username From User Where Username=?", input.Email)
	if len(users) == 0 {
		var msg = fmt.Sprintf("%s isn't registered", input.Email)
		var output = friendmodels.GetFriendsOutput{
			apimodels.ApiResult{false,  []string {msg}},
			0,
			[]string{}}
		return output
	}
	// Return Data
	var currentUser = users[0]

	// Blocked Users
	var blockedIds []int
	_,  _ = persistence.DbContext.Select(&blockedIds,"Select Target From Subscribe_User Where Requestor = ? And Status=?", currentUser.Id, constants.Blocked)
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
		_,  _ = persistence.DbContext.Select(&emails,query)
	}
	var output = friendmodels.GetFriendsOutput{
		apimodels.ApiResult{true, []string {}},
		len(emails),
		emails}
	return output
}