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



func GetCommonFriendsCommand(context *gin.Context) {
	var input apimodels.GetCommonFriendsInput
	var err error

	// 1
	if err = context.BindJSON(&input) ; err != nil {
		var output = &apimodels.GetCommonFriendsOutput{false, []string {"Input isn't null"},0, []string{}}
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 2
	if helper.IsNull(input.Friends) || len(input.Friends) < 2 {
		var output = &apimodels.GetCommonFriendsOutput{false, []string {"Input isn't valid"},0, []string{} }
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 3
	var count = len(input.Friends)
	var output = &apimodels.GetCommonFriendsOutput{true, []string {},0, []string{}}
	for i := 0; i < count; i++ {
		if !helper.IsEmail(input.Friends[i]) {
			output.Success = false
			var msg string = fmt.Sprintf("%s isn't an Email address", input.Friends[i])
			output.Msgs = helper.AddItemToArray(output.Msgs, msg)
		}
	}
	if output.Success == false {
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// 4
	var users []domain.UserDomain
	_, err = persistence.DbContext.Select(&users, "select Id, Username From User Where Username In (?,?)", input.Friends[0], input.Friends[1])
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
		context.JSON(http.StatusBadRequest, output)
		return
	}

	// TODO: Block User

	// Search Intersection Users
	var user1 = users[0]
	var user2 = users[1]

	var fromUserIds1 []int
	_, err = persistence.DbContext.Select(&fromUserIds1,"Select FromUserID From User_Friend Where ToUserID = ? And FromUserID != ?", user1.Id, user2.Id)
	var toUserIds1 []int
	_, err = persistence.DbContext.Select(&toUserIds1,"Select ToUserID From User_Friend Where FromUserID =? And ToUserID != ?", user1.Id, user2.Id)

	var userIds1 = append(fromUserIds1, toUserIds1...)


	var fromUserIds2 []int
	_, err = persistence.DbContext.Select(&fromUserIds2,"Select FromUserID From User_Friend Where ToUserID = ? And FromUserID != ?", user2.Id, user1.Id)
	var toUserIds2 []int
	_, err = persistence.DbContext.Select(&toUserIds2,"Select ToUserID From User_Friend Where FromUserID = ? And ToUserID != ?", user2.Id, user1.Id)

	var userIds2 = append(fromUserIds1, toUserIds1...)

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


		_,  err = persistence.DbContext.Select(&emails,query)
		output = &apimodels.GetCommonFriendsOutput{true, []string {}, len(emails), emails}
	}

	context.JSON(http.StatusOK, output)
}