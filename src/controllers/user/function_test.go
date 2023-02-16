package user

import (
	"testing"

	"FDPD-BACKEND/src/controllers/user/models"
)

func TestAdd(t *testing.T) {

	want := true

	loginSuccesExample := &models.Login{
		Email:    "nicolas@ucn.cl",
		Password: "xxxxxx",
	}

	if gotFromSuccesExample := loginSuccesExample.ValidateDomain(); gotFromSuccesExample != want {
		t.Errorf("error in succes login example, got %t, wanted %t", gotFromSuccesExample, want)
	}

	loginErrorExample := &models.Login{
		Email:    "nicolas@gmail.cl",
		Password: "xxxxxx",
	}

	if gotFromErrorExample := loginErrorExample.ValidateDomain(); gotFromErrorExample == want {
		t.Errorf("error in error login example, got %t, wanted %t", gotFromErrorExample, want)
	}

}
