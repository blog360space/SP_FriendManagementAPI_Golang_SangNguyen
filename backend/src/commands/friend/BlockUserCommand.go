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

// Block an User docs
// @Summary Block an User
// @Description As a user, I need an API to block updates from an email address.
// @Tags Friend
// @Accept  json
// @Produce  json
// @Param requestor body string true "Email"
// @Param target body string true "Email"
// @Success 201 {object} friend.BlockUserOutput
// @Failure 400 {object} friend.BlockUserOutput
// @Router /friend/block-user [post]
func BlockUserCommand(context *gin.Context)  {
	var input apimodels.BlockUserInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = &apimodels.BlockUserOutput{false, []string {"Input isn't null"}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 2
	if len(input.Requestor) == 0 || len(input.Target) == 0 {
		var output = &apimodels.BlockUserOutput{false, []string {"Input isn't valid"} }
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 3
	if !helper.IsEmail(input.Requestor) {
		var output = &apimodels.BlockUserOutput{false, []string {"Requestor isn't valid email address"} }
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 4
	if !helper.IsEmail(input.Target) {
		var output = &apimodels.BlockUserOutput{false, []string {"Target isn't valid email address"} }
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 5
	var output = &apimodels.BlockUserOutput{true, []string {}}
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
		// 7
		if subscribeUser.Status == constants.Blocked {
			var msg = fmt.Sprintf("%s blocked %s", requestor.Username, target.Username)
			output.Success = false
			output.Msgs = helper.AddItemToArray(output.Msgs, msg)

			context.JSON(http.StatusBadRequest, output)
			return
		}

		// Update
		subscribeUser.Status = constants.Blocked
		persistence.DbContext.Update(subscribeUser)
		context.JSON(http.StatusOK, output)
	} else {
		// Insert
		subscribeUser:= &domain.SubscribeUserDomain{0, requestor.Id, target.Id, constants.Blocked}
		persistence.DbContext.Insert(subscribeUser)
		context.JSON(http.StatusOK, output)
	}
}