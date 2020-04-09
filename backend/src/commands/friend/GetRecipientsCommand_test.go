package friend

import (
	"github.com/stretchr/testify/assert"
	"spapp/src/commands/user"
	helper "spapp/src/common/helpers"
	friendmodels "spapp/src/models/apimodels/friend"
	usermodels "spapp/src/models/apimodels/user"
	"strconv"
	"testing"
	"fmt"
)

func Test_GetRecipients_Ok(t *testing.T){
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
	_ = SubscribeUserCommand(friendmodels.SubscribeUserInput{username1, username2})

	username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	count += 1
	input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)
	_ = BlockUserCommand(friendmodels.BlockUserInput{username2, username1})

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
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username2, username1}})

	username2 = fmt.Sprintf("%s_%s@%s.com", helper.RandomString(8),strconv.Itoa(count), helper.RandomString(4))
	count += 1
	input2 = usermodels.RegisterUserInput{
		username2,
	}
	_ = user.RegisterUserCommand(input2)
	_ = MakeFriendCommand(friendmodels.MakeFriendInput{[]string{username2, username1}})

	var input = friendmodels.GetRecipientsInput{username1, "Hi sangnv1.pr@gmail.com, sangnv1.prfortesting@gmail.com"}
	var output = GetRecipientsCommand(input)
	assert.True(t, output.Success)
}

func Test_GetRecipients_BadRequestCase2(t *testing.T) {
	// Config
	initConfig()

	var input = friendmodels.GetRecipientsInput{"", "Hi sangnv1.pr@gmail.com, sangnv1.prfortesting@gmail.com"}
	var output = GetRecipientsCommand(input)
	assert.False(t, output.Success)
}

func Test_GetRecipients_BadRequestCase3(t *testing.T) {
	// Config
	initConfig()

	var input = friendmodels.GetRecipientsInput{"sangnv", "Hi sangnv1.pr@gmail.com, sangnv1.prfortesting@gmail.com"}
	var output = GetRecipientsCommand(input)
	assert.False(t, output.Success)
}

func Test_GetRecipients_BadRequestCase4(t *testing.T) {
	// Config
	initConfig()

	var input = friendmodels.GetRecipientsInput{"sangnv@ithink.vn", ""}
	var output = GetRecipientsCommand(input)
	assert.False(t, output.Success)
}

func Test_GetRecipients_BadRequestCase5(t *testing.T) {
	// Config
	initConfig()

	var input = friendmodels.GetRecipientsInput{"sangnv@ithink.vn", "Hi sangnv1.pr@gmail.com, sangnv1.prfortesting@gmail.com"}
	var output = GetRecipientsCommand(input)
	assert.False(t, output.Success)
}