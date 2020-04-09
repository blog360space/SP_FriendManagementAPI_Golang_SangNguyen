package friend

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"spapp/src/commands/user"
	helper "spapp/src/common/helpers"
	friendmodels "spapp/src/models/apimodels/friend"
	usermodels "spapp/src/models/apimodels/user"
	"strconv"
	"testing"
)

func Test_GetFriends_Ok(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()
	var count = len(users)
	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count + 1), helper.RandomString(4))
	count += 2
	var input1 = usermodels.RegisterUserInput{
		username1,
	}
	_ = user.RegisterUserCommand(input1)
	var input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)

	// make friends
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username1, username2}})
	_ = BlockUserCommand(friendmodels.BlockUserInput{username1, username2})

	username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	count += 1
	input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username1, username2}})

	username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	count += 1
	input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username1, username2}})

	var input = friendmodels.GetFriendsInput{username1}
	var output = GetFriendsCommand(input)
	assert.True(t, output.Success)
}

func Test_GetFriends_BadRequestCase2(t *testing.T){
	// Config
	initConfig()

	var input = friendmodels.GetFriendsInput{""}

	var output = GetFriendsCommand(input)

	assert.False(t, output.Success)
}

func Test_GetFriends_BadRequestCase3(t *testing.T){
	// Config
	initConfig()

	var input = friendmodels.GetFriendsInput{"sangnv"}

	var output = GetFriendsCommand(input)

	assert.False(t, output.Success)
}

func Test_GetFriends_BadRequestCase4(t *testing.T){
	// Config
	initConfig()
	var users = getAllUsers()

	var username = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(len(users)), helper.RandomString(4))
	var input = friendmodels.GetFriendsInput{username}

	var output = GetFriendsCommand(input)

	assert.False(t, output.Success)
}