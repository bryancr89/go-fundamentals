package model

import (
	"crypto/sha512"
	"encoding/base64"
	"database/sql"
	"fmt"
)

const passwordSalt = "asfq23412dsafasdfasdfa1+2"

type User struct {
	Id int
	Email string
	Password string
}

func Login(email, password string) (*User, error) {
	result := &User{}
	hasher := sha512.New()
	hasher.Write([]byte(passwordSalt))
	hasher.Write([]byte(email))
	hasher.Write([]byte(password))
	pwd := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println(pwd)
	row := db.QueryRow(`
		SELECT id, email
		FROM public.users
		WHERE email = $1
			AND password = $2`, email, pwd)
	err := row.Scan(&result.Id, &result.Email)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("User not found")
	case err != nil:
		return nil, err
	}
	return result, err
}