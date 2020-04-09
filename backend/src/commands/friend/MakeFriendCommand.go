package friend

import (
	"fmt"
	"spapp/src/common/constants"
	"spapp/src/models/apimodels"

	helper "spapp/src/common/helpers"
	friendmodels "spapp/src/models/apimodels/friend"

	"spapp/src/models/domain"
	"spapp/src/persistence"
)


func MakeFriendCommand (input friendmodels.MakeFriendInput) friendmodels.MakeFriendOutput{

	// 2
	if helper.IsNull(input.Friends) || len(input.Friends) < 2 {
		var output = friendmodels.MakeFriendOutput{apimodels.ApiResult{false, []string {"Input isn't valid"}}}
		return output
	}

	// 3
	var count = len(input.Friends)
	var output = friendmodels.MakeFriendOutput{apimodels.ApiResult{true, []string {}}}
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
	_, _ = persistence.DbContext.Select(&users, "select Id, Username From User Where Username In (?,?) Order By Id", input.Friends[0], input.Friends[1])
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

	// 6
	var userfriends []domain.UserFriendDomain
	_, _ = persistence.DbContext.Select(&userfriends,"Select Id, FromUserID, ToUserID From User_Friend Where (FromUserID=? And ToUserID=?) Or (FromUserID=? And ToUserID=?)", users[0].Id, users[1].Id, users[1].Id, users[0].Id)

	if len(userfriends) > 0 {
		var msg = fmt.Sprintf("%s and %s are existed connection", input.Friends[0], input.Friends[1])
		output.Success = false
		output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		return output
	}

	// 7
	var blockUsers []domain.SubscribeUserDomain
	_, _ = persistence.DbContext.Select(&blockUsers, "Select Id, Requestor, Target, Status From Subscribe_User Where Requestor=? And Target=? And Status=?", users[0].Id, users[1].Id, constants.Blocked)

	if len(blockUsers) > 0 {
		var msg = fmt.Sprintf("%s blocked %s", users[0].Username, users[1].Username)
		output.Success = false
		output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		return output
	}

	userfriend := &domain.UserFriendDomain{0, users[0].Id, users[1].Id}
	persistence.DbContext.Insert(userfriend)
	return output
}