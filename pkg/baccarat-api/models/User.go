package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type BaseUserDto struct {
	Account  string `json:"account" binding:"required" validate:"required" example:"test001" gorm:"type:varchar(40);uniqueIndex"`
	Password string `json:"password" binding:"required" validate:"required" example:"a12345678" gorm:"type:varchar(255);"`
	Name     string `json:"name" binding:"required" validate:"required" examle:"test001" gorm:"type:varchar(40);"`
	Amount   int    `json:"amount" validate:"required" example:0 gorm:"default:0"`
}

type User struct {
	gorm.Model
	BaseUserDto
}

func hashAndSalt(pwd []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.DefaultCost)
	return string(hash), err
}

func comparePasswords(hashedPwd string, plainPwd []byte) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPwd)
	if err != nil {
		return false
	}

	return true
}

func CreateUser(values BaseUserDto) (User, *gorm.DB) {
	var u User

	hashedPassword, _ := hashAndSalt([]byte(values.Password))
	u.Account = values.Account
	u.Password = hashedPassword
	u.Name = values.Name
	u.Amount = values.Amount
	result := DBConnection.Create(&u)
	return u, result
}

func UpdateUserById(id string, name string) *User {
	var userResult *User
	DBConnection.Select([]string{"id", "name", "account", "password"}).First(&userResult, "id = ?", id)

	userResult.Name = name
	DBConnection.Save(&userResult)
	return userResult
}

func GetUserByAccountAndPassword(account string, password string) *User {
	var u *User
	DBConnection.Select([]string{"id", "name", "account", "password"}).First(&u, "account = ?", account)

	pwdMatch := comparePasswords(u.Password, []byte(password))

	if pwdMatch {
		return u
	}

	return nil
}

func GetUserById(id string) *User {
	var u *User
	DBConnection.Select([]string{"id", "name", "account", "password"}).First(&u, "id = ?", id)

	return u
}

func DeleteUserById(id string) {
	var userResult *User
	DBConnection.Select([]string{"id"}).First(&userResult, "id = ?", id)
	if userResult != nil && userResult.ID != 0 {
		DBConnection.Delete(&userResult)
	}
}
