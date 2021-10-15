package entity

import (
	"encoding/json"
	"fmt"
)

type UserStatus string

const (
	UserStatus_INACTIVE UserStatus = "INACTIVE"
	UserStatus_ACTIVE   UserStatus = "ACTIVE"
)

type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	CreatedAt int64      `json:"created_at"`
	ExpiredAt int64      `json:"expired_at"`
	Status    UserStatus `json:"status"`
}

func (u *User) IsActive() bool {
	return u.Status == UserStatus_ACTIVE
}

func (u *User) IsInActive() bool {
	return u.Status == UserStatus_INACTIVE
}

func (u *User) ToJson() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(b), nil
}
