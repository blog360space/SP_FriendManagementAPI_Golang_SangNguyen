package friend

import (
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	"fmt"
	"spapp/src/commands/user"
	helper "spapp/src/common/helpers"
	friendmodels "spapp/src/models/apimodels/friend"
	usermodels "spapp/src/models/apimodels/user"
	"spapp/src/models/domain"
	"spapp/src/persistence"
	"strconv"
	"testing"
)

func initConfig() {
	// Config
	err := godotenv.Load("../../../env/env.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	persistence.UseMySql()
}

func getAllUsers() []domain.UserDomain {
	var users []domain.UserDomain
	_, _ = persistence.DbContext.Select(&users, "Select Id, UserName From User")
	return users
}

func getAllFriendUsers() []domain.UserFriendDomain{
	var friendUsers []domain.UserFriendDomain
	_, _ = persistence.DbContext.Select(&friendUsers, "Select Id, ToUserID, FromUserID From User_Friend")
	return friendUsers
}


func Test_MakeFriend_Ok(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()

	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(len(users)), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(len(users) + 1), helper.RandomString(4))
	var input1 = usermodels.RegisterUserInput{
		username1,
	}
	_ = user.RegisterUserCommand(input1)
	var input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)

	var input = friendmodels.MakeFriendInput{[]string{username1,username2}}

	var output = MakeFriendCommand(input)

	assert.True(t, output.Success)
}

func Test_MakeFriend_BadRequestCase2(t *testing.T){
	// Config
	initConfig()

	var input = friendmodels.MakeFriendInput{[]string{}}

	var output = MakeFriendCommand(input)

	assert.False(t, output.Success)
}

func Test_MakeFriend_BadRequestCase3(t *testing.T){
	// Config
	initConfig()

	var input = friendmodels.MakeFriendInput{[]string{"sangnv1", "sangnv2"}}

	var output = MakeFriendCommand(input)

	assert.False(t, output.Success)
}

func Test_MakeFriend_BadRequestCase4(t *testing.T){
	// Config
	initConfig()

	var input = friendmodels.MakeFriendInput{[]string{"sangnv.pr@gmail.com", "sangnv.pr@gmail.com"}}

	var output = MakeFriendCommand(input)

	assert.False(t, output.Success)
}

func Test_MakeFriend_BadRequestCase5(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()

	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(len(users)), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(len(users) + 1), helper.RandomString(4))
	var input = friendmodels.MakeFriendInput{[]string{username1, username2}}

	var output = MakeFriendCommand(input)

	assert.False(t, output.Success)
}

func Test_MakeFriend_BadRequestCase6(t *testing.T){
	// Config
	initConfig()

	var users = getAllUsers()

	var username1 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(len(users)), helper.RandomString(4))
	var username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(len(users) + 1), helper.RandomString(4))
	var input1 = usermodels.RegisterUserInput{
		username1,
	}
	_ = user.RegisterUserCommand(input1)
	var input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)

	var input = friendmodels.MakeFriendInput{[]string{username1,username2}}

	_ = MakeFriendCommand(input)
	var output = MakeFriendCommand(input)

	assert.False(t, output.Success)
}

