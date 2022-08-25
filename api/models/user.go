package models

import (
	"JSON-API-GOLANG/api/utils"
	"errors"
	"regexp"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `gorm:"size:100;not null;<-:update" json:"name"`
	Email    string `gorm:"size:255;not null;unique;<-:create" json:"email"`
	Password string `gorm:"size:255;not null;<-:update" json:"password"`
	Age      int    `gorm:"size:3;not null;<-:update" json:"age"`
}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, password string) (string, error) {
	u := User{}

	if err := db.Model(User{}).Where("Email = ?", email).Take(&u).Error; err != nil {
		return "", err
	}

	if err := VerifyPassword(password, u.Password); err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := utils.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *User) Create() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	if err := db.Create(&u).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Update() error {
	if err := db.Model(&u).Updates(User{Email: u.Email, Age: u.Age, Name: u.Name}).Error; err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	if err := db.Delete(&u).Error; err != nil {
		return err
	}
	db.Where("UserID <> ?", u.ID).Delete(&Files{})
	return nil
}

func (u *User) BeforeSave() error {

	if u.Age <= 18 {
		return errors.New("Age must be greater than 18!")
	}

	if match, _ := regexp.MatchString("^[\\w-\\.]+@([\\w-]+\\.)+[\\w-]{2,4}$", u.Email); !match {
		return errors.New("You must provide a valid email")
	}

	return nil
}

func GetUserByID(userID uint64) (*User, error) {
	var u User

	if err := db.First(&u, userID).Error; err != nil {
		return nil, errors.New("User not found")
	}

	u.Password = ""
	return &u, nil
}

func GetAllUsers() (*[]User, error) {
	var Users []User
	if err := db.Find(&Users).Error; err != nil {
		return nil, err
	}
	for i := range Users {
		Users[i].Password = ""
	}

	return &Users, nil
}
