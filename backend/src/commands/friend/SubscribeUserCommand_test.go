package friend

import (
	"github.com/stretchr/testify/assert"
	"spapp/src/commands/user"
	"fmt"
	helper "spapp/src/common/helpers"
	friendmodels "spapp/src/models/apimodels/friend"
	usermodels "spapp/src/models/apimodels/user"
	"strconv"
	"testing"
)

func Test_SubscribeUser_Ok(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()
	var count = len(users)
	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count + 1), helper.RandomString(4))
	var input1 = usermodels.RegisterUserInput{
		username1,
	}
	_ = user.RegisterUserCommand(input1)
	var input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)

	var input = friendmodels.SubscribeUserInput{username1,username2}

	var output = SubscribeUserCommand(input)

	assert.True(t, output.Success)
}

func Test_SubscribeUser_BadRequestCase2(t *testing.T){
	// Config
	initConfig()

	var input = friendmodels.SubscribeUserInput{"username1",""}
	var output = SubscribeUserCommand(input)

	assert.False(t, output.Success)
}

func Test_SubscribeUser_BadRequestCase3(t *testing.T){
	// Config
	initConfig()


	var input = friendmodels.SubscribeUserInput{"username1","username2"}
	var output = SubscribeUserCommand(input)

	assert.False(t, output.Success)
}

func Test_SubscribeUser_BadRequestCase4(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()
	var count= len(users)
	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))

	var input = friendmodels.SubscribeUserInput{username1,"username1"}

	var output = SubscribeUserCommand(input)

	assert.False(t, output.Success)
}

func Test_SubscribeUser_BadRequestCase5(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()
	var count= len(users)
	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))

	var input = friendmodels.SubscribeUserInput{username1,username1}

	var output = SubscribeUserCommand(input)

	assert.False(t, output.Success)
}

func Test_SubscribeUser_BadRequestCase6(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()
	var count= len(users)
	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count + 1), helper.RandomString(4))

	var input1 = usermodels.RegisterUserInput{
		username1,
	}
	_ = user.RegisterUserCommand(input1)


	var input = friendmodels.SubscribeUserInput{username1,username2}

	var output = SubscribeUserCommand(input)

	assert.False(t, output.Success)
}

func Test_SubscribeUser_BadRequestCase7_1(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()
	var count= len(users)
	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count + 1), helper.RandomString(4))
	var input1 = usermodels.RegisterUserInput{
		username1,
	}
	_ = user.RegisterUserCommand(input1)
	var input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)

	var input = friendmodels.SubscribeUserInput{username1,username2}
	var output = SubscribeUserCommand(input)
	output = SubscribeUserCommand(input)

	assert.False(t, output.Success)
}

func Test_SubscribeUser_BadRequestCase7_2(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()
	var count= len(users)
	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count + 1), helper.RandomString(4))
	var input1 = usermodels.RegisterUserInput{
		username1,
	}
	_ = user.RegisterUserCommand(input1)
	var input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)
	var temp = friendmodels.BlockUserInput{username1, username2}
	_ = BlockUserCommand(temp)

	var input = friendmodels.SubscribeUserInput{username1,username2}
	var output = SubscribeUserCommand(input)
	output = SubscribeUserCommand(input)

	assert.False(t, output.Success)
}