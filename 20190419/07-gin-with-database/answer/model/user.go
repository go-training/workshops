package model

import (
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog/log"
)

// User struct
type User struct {
	ID    int64  `json:"id" gorm:"primary_key"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateUser data
func CreateUser(user *User) {
	DB.Create(user)
}

// FindAllUsers data
func FindAllUsers(users *[]User) *gorm.DB {
	return DB.Find(users)
}

// FindUserByID ...
func FindUserByID(id int64) *User {
	user := &User{}
	DB.First(user, id)

	return user
}

// DeleteUser ...
func DeleteUser(id int64) *User {
	user := &User{}
	if err := DB.First(user, id).Error; err != nil {
		log.Debug().Err(err).Msg("can't find user")
	}

	DB.Delete(user, id)

	return user
}
