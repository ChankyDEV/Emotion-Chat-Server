package models

import (
	"encoding/json"
	"io"
)

type User struct {
	Uid      string `json:"uid"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func (user *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(user)
}
