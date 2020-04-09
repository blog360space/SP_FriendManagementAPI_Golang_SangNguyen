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

func Test_GetCommonFriends_Ok(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()
	var count = len(users)
	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count + 1), helper.RandomString(4))
	var username3 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count + 2), helper.RandomString(4))
	count += 3
	var input1 = usermodels.RegisterUserInput{
		username1,
	}
	_ = user.RegisterUserCommand(input1)
	var input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)

	var input3 = usermodels.RegisterUserInput{
		username3,
	}
	_ = user.RegisterUserCommand(input3)

	// make friends
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username1, username3}})
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username2, username3}})
	_ = BlockUserCommand(friendmodels.BlockUserInput{username1, username3})
	_ = BlockUserCommand(friendmodels.BlockUserInput{username2, username3})


	username3 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	count += 1
	input3 = usermodels.RegisterUserInput{
		username3,
	}
	_ = user.RegisterUserCommand(input3)
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username1, username3}})
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username2, username3}})

	username3 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	count += 1
	input3 = usermodels.RegisterUserInput{
		username3,
	}
	_ = user.RegisterUserCommand(input3)
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username1, username3}})
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username2, username3}})

	var input = friendmodels.GetCommonFriendsInput{ []string{username1,username2}}
	var output = GetCommonFriendsCommand(input)
	assert.True(t, output.Success)
}

func Test_GetCommonFriends_BadRequestCase2(t *testing.T) {
	// Config
	initConfig()

	var username1 = "sangnv1"
	var input = friendmodels.GetCommonFriendsInput{ []string{username1}}
	var output = GetCommonFriendsCommand(input)
	assert.False(t, output.Success)
}

func Test_GetCommonFriends_BadRequestCase3(t *testing.T) {
	// Config
	initConfig()

	var username1 = "sangnv1"
	var username2 = "sangnv2"
	var input = friendmodels.GetCommonFriendsInput{ []string{username1, username2}}
	var output = GetCommonFriendsCommand(input)
	assert.False(t, output.Success)
}

func Test_GetCommonFriends_BadRequestCase4(t *testing.T) {
	// Config
	initConfig()


	var username1 = "sangnv1@gmail.com"
	var username2 = "sangnv1@gmail.com"
	var input = friendmodels.GetCommonFriendsInput{ []string{username1, username2}}
	var output = GetCommonFriendsCommand(input)
	assert.False(t, output.Success)
}

func Test_GetCommonFriends_BadRequestCase5(t *testing.T) {
	// Config
	initConfig()

	var users = getAllUsers()
	var count = len(users)

	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count + 1), helper.RandomString(4))
	var input = friendmodels.GetCommonFriendsInput{ []string{username1, username2}}
	var output = GetCommonFriendsCommand(input)
	assert.False(t, output.Success)
}

