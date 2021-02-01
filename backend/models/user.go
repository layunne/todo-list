package models

import (
	"encoding/json"
)

type User struct {
	Id       string `bson:"_id" json:"id"`
	Name     string `bson:"name" json:"name"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

func (u *User) Bytes() []byte {

	bytes, err := json.Marshal(u)

	if err != nil {
		return nil
	}
	return bytes
}

func (u *User) String() string {

	return string(u.Bytes())
}

type CreateUser struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUser struct {
	Name        string `json:"name"`
	Username    string `json:"username"`
	OldPassword string `json:"oldPassword"`
	Password    string `json:"password"`
}

type UserDTO struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func (u *User) ToDTO(token string) *UserDTO {

	if u == nil {
		return nil
	}

	return &UserDTO{
		Id:       u.Id,
		Name:     u.Name,
		Username: u.Username,
		Token:    token,
	}
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
