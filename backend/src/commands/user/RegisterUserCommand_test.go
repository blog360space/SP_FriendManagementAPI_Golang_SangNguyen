package user

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"log"
	helper "spapp/src/common/helpers"
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

func closeDb(){
	defer persistence.DbContext.Db.Close()
}

func getAllUsers() []domain.UserDomain {
	var users []domain.UserDomain
	_, _ = persistence.DbContext.Select(&users, "Select Id, UserName From User")
	return users
}

func Test_RegisterUser_Ok(t *testing.T) {
	// Config
	initConfig()
	var users = getAllUsers()
	var username = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(len(users)), helper.RandomString(4))
	log.Printf(username)
	var input = usermodels.RegisterUserInput{
		username,
	}
	var output = RegisterUserCommand(input)
	closeDb()
	assert.True(t, output.Success)
}

func Test_RegisterUser_BadRequest_Case2(t *testing.T) {
	// Config
	initConfig()

	var username = ""
	var input = usermodels.RegisterUserInput{username}
	var output = RegisterUserCommand(input)

	closeDb()
	assert.False(t, output.Success)
	assert.Equal(t, output.Msgs[0], "Username isn't null or empty")
}

func Test_RegisterUser_BadRequest_Case3(t *testing.T) {
	// Config
	initConfig()

	var username = "test"
	var input = usermodels.RegisterUserInput{username}
	var output = RegisterUserCommand(input)
	closeDb()
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
	closeDb()
	assert.False(t, output.Success)
	assert.Equal(t, output.Msgs[0], "Username is existed")
}

