package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser(t *testing.T) {
	PrepareTestDatabase()
	user, _ := FindUserByID(1)

	assert.Equal(t, "appleboy", user.Name)
	assert.Equal(t, "test01@gmail.com", user.Email)
}

func TestCreateUser(t *testing.T) {
	PrepareTestDatabase()
	user := &User{
		ID:    1000,
		Name:  "appleboy",
		Email: "test01@gmail.com",
	}

	err := CreateUser(user)
	assert.Nil(t, err)

	user, _ = FindUserByID(1000)
	assert.Equal(t, "appleboy", user.Name)
	assert.Equal(t, "test01@gmail.com", user.Email)
}
