package friend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"fmt"
	"spapp/src/common/constants"

	helper "spapp/src/common/helpers"
	apimodels "spapp/src/models/apimodels/friend"
	"spapp/src/models/domain"
	"spapp/src/persistence"
)

// Subscribe an User docs
// @Summary Subscribe an User
// @Description As a user, I need an API to subscribe to updates from an email address.
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param input body friend.SubscribeUserInput true "Input"
// @Success 201 {object} friend.SubscribeUserOutput
// @Failure 400 {object} friend.SubscribeUserOutput
// @Router /friend/subscribe-user [post]
func SubscribeUserCommand(context *gin.Context)  {
	var input apimodels.SubscribeUserInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = &apimodels.SubscribeUserOutput{false, []string {"Input isn't null"}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 2
	if len(input.Requestor) == 0 || len(input.Target) == 0 {
		var output = &apimodels.SubscribeUserOutput{false, []string {"Input isn't valid"} }
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 3
	if !helper.IsEmail(input.Requestor) {
		var output = &apimodels.SubscribeUserOutput{false, []string {"Requestor isn't valid email address"} }
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 4
	if !helper.IsEmail(input.Target) {
		var output = &apimodels.SubscribeUserOutput{false, []string {"Target isn't valid email address"} }
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 5
	var output = &apimodels.SubscribeUserOutput{true, []string {}}
	var users []domain.UserDomain
	_, err = persistence.DbContext.Select(&users,"Select Id, Username From User Where Username=? Or Username=?", input.Requestor, input.Target)

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
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 6
	index := helper.ArrayIndex(len(users), func(i int) bool {
		return users[i].Username == input.Requestor
	})
	var requestor = users[index]
	index = helper.ArrayIndex(len(users), func(i int) bool {
		return users[i].Username == input.Target
	})
	var target = users[index]

	var subscribeUsers []domain.SubscribeUserDomain
	_, err = persistence.DbContext.Select(&subscribeUsers, "Select Id, Requestor, Target, Status From Subscribe_User Where Requestor=? And Target=?", requestor.Id, target.Id)

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

		context.JSON(http.StatusBadRequest, output)
		return
	}

	// Return Data
	subscribeUser:= &domain.SubscribeUserDomain{0, requestor.Id, target.Id, constants.Subscribed}
	persistence.DbContext.Insert(subscribeUser)
	context.JSON(http.StatusOK, output)
}