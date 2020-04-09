package user

import (
	"github.com/joho/godotenv"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	helper "spapp/src/common/helpers"
	"spapp/src/models/domain"
	usermodels "spapp/src/models/apimodels/user"
	"spapp/src/persistence"
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

func Test_RegisterUser_Ok(t *testing.T) {
	// Config
	initConfig()

	var username = fmt.Sprintf("%s@%s.com", helper.RandomString(5), helper.RandomString(4))
	var input = usermodels.RegisterUserInput{
		username,
	}
	var output = RegisterUserCommand(input)

	assert.True(t, output.Success)
}

func Test_RegisterUser_BadRequest_Case2(t *testing.T) {
	// Config
	initConfig()

	var username = ""
	var input = usermodels.RegisterUserInput{username}
	var output = RegisterUserCommand(input)

	assert.False(t, output.Success)
	assert.Equal(t, output.Msgs[0], "Username isn't null or empty")
}

func Test_RegisterUser_BadRequest_Case3(t *testing.T) {
	// Config
	initConfig()

	var username = "test"
	var input = usermodels.RegisterUserInput{username}
	var output = RegisterUserCommand(input)

	assert.False(t, output.Success)
	assert.Equal(t, output.Msgs[0], "Username isn't valid email address")
}

func Test_RegisterUser_BadRequest_Case4(t *testing.T) {
	// Config
	initConfig()

	var users = getAllUsers()

	var username = users[0].Username
	var input = usermodels.RegisterUserInput{username}
	var output = RegisterUserCommand(input)

	assert.False(t, output.Success)
	assert.Equal(t, output.Msgs[0], "Username is existed")
}

