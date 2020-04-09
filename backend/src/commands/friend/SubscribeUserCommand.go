package friend

import (
	"fmt"
	"reflect"
	"spapp/src/common/constants"
	"spapp/src/models/apimodels"

	helper "spapp/src/common/helpers"
	friendmodels "spapp/src/models/apimodels/friend"
	"spapp/src/models/domain"
	"spapp/src/persistence"
)

func SubscribeUserCommand(input friendmodels.SubscribeUserInput)  friendmodels.SubscribeUserOutput {

	// 1
	if helper.IsNull(input) {
		var output = friendmodels.SubscribeUserOutput{apimodels.ApiResult{false, []string {"Input isn't null"}}}
		return output
	}

	// 2
	if len(input.Requestor) == 0 || len(input.Target) == 0 {
		var output = friendmodels.SubscribeUserOutput{apimodels.ApiResult{false, []string {"Input isn't valid"}}}
		return output
	}

	// 3
	if !helper.IsEmail(input.Requestor) {
		var output = friendmodels.SubscribeUserOutput{apimodels.ApiResult{false, []string {"Requestor isn't valid email address"}}}
		return output
	}

	// 4
	if !helper.IsEmail(input.Target) {
		var output = friendmodels.SubscribeUserOutput{apimodels.ApiResult{false, []string {"Target isn't valid email address"}}}
		return output
	}

	// 5
	if input.Requestor == input.Target {
		var output = friendmodels.SubscribeUserOutput{apimodels.ApiResult{false, []string {"Requestor and Target are the same"}}}
		return output
	}

	// 6
	var output = friendmodels.SubscribeUserOutput{apimodels.ApiResult{true, []string {}}}
	var users []domain.UserDomain
	_, _ = persistence.DbContext.Select(&users,"Select Id, Username From User Where Username=? Or Username=?", input.Requestor, input.Target)

	if len(users) != 2 {
		v := reflect.ValueOf(input)
		values := make([]interface{}, v.NumField())
		for i := 0; i < v.NumField(); i++ {
			values[i] = v.Field(i).Interface()
			flag := true
			for j := range users {
				if values[i] == users[j].Username {
					flag = false
				}
			}
			if flag {
				var msg = fmt.Sprintf("%s isn't registered", values[i])
				output.Success = false
				output.Msgs = helper.AddItemToArray(output.Msgs, msg)
			}
		}
		return output
	}

	// 7
	index := helper.ArrayIndex(len(users), func(i int) bool {
		return users[i].Username == input.Requestor
	})
	var requestor = users[index]
	index = helper.ArrayIndex(len(users), func(i int) bool {
		return users[i].Username == input.Target
	})
	var target = users[index]

	var subscribeUsers []domain.SubscribeUserDomain
	_, _ = persistence.DbContext.Select(&subscribeUsers, "Select Id, Requestor, Target, Status From Subscribe_User Where Requestor=? And Target=?", requestor.Id, target.Id)

	if len(subscribeUsers) > 0 {
		subscribeUser := subscribeUsers[0]
		// Check Status
		if subscribeUser.Status == constants.Subscribed {
			var msg = fmt.Sprintf("%s subscribed %s", requestor.Username, target.Username)
			output.Success = false
			output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		}

		if subscribeUser.Status == constants.Blocked {
			var msg = fmt.Sprintf("%s blocked %s", requestor.Username, target.Username)
			output.Success = false
			output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		}

		return output
	}

	// Return Data
	subscribeUser:= &domain.SubscribeUserDomain{0, requestor.Id, target.Id, constants.Subscribed}
	persistence.DbContext.Insert(subscribeUser)

	return output
}