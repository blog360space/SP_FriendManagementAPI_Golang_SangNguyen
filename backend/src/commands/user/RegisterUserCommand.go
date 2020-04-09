package user

import (
	helper "spapp/src/common/helpers"
	usermodels "spapp/src/models/apimodels/user"
	"spapp/src/models/domain"
	"spapp/src/persistence"
)

func RegisterUserCommand (input usermodels.RegisterUserInput) usermodels.RegisterUserOutput {

	// 1
	if helper.IsNull(input) {
		var output usermodels.RegisterUserOutput
		output.Success = false
		output.Msgs = []string{"Input isn't null"}
		return output
	}

	// 2
	if len(input.Username) == 0 {
		var output usermodels.RegisterUserOutput
		output.Success = false
		output.Msgs = []string{"Username isn't null or empty"}
		return output
	}

	// 3
	if !helper.IsEmail(input.Username) {
		var output usermodels.RegisterUserOutput
		output.Success = false
		output.Msgs = []string{"Username isn't valid email address"}
		return output
	}

	// 4
	var users []domain.UserDomain
	_, _ = persistence.DbContext.Select(&users,"Select Id, Username From User Where Username=?", input.Username)

	if len(users) > 0 {
		var output usermodels.RegisterUserOutput
		output.Success = false
		output.Msgs = []string{"Username is existed"}
		return output
	}

	user := &domain.UserDomain{ 0, input.Username }
	persistence.DbContext.Insert(user)
	var output usermodels.RegisterUserOutput
	output.Success = true
	output.Data.Id = user.Id
	output.Data.Username =  user.Username
	return output

}
