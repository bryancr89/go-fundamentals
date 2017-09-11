package model

import (
	"testing"
	"crypto/sha512"
	"encoding/base64"
)

func TestLoginSendCorrectPasswordHash(t *testing.T) {
	testDB := new(mockDB)
	testDB.returnedRow = &mockRow{}
	db = testDB

	password := "foo"
	email := "bar"
	Login(email, password)

	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))

	if testDB.lastArgs[1] != pwd {
		t.Errorf("Login function failed to send correct password")
	}
}