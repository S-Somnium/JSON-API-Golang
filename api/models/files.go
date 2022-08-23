package models

import "errors"

type Files struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	Name     string `gorm:"size:100;not null;<-:update" json:"name"`
	Location string `gorm:"size:255;not null;unique;<-:update" json:"location"`
	UserID   int    `gorm:"not null" json:"UserID"`
}

func (f *Files) Create() error {
	if err := db.Create(&f).Error; err != nil {
		return err
	}
	return nil
}

func (f *Files) BeforeSave() error {
	var u User

	if err := db.First(&u, f.UserID).Error; err != nil {
		return errors.New("User not found")
	}
	return nil

}

func DeleteAllFiles(userID uint64) error {
	var u User

	if err := db.First(&u, userID).Error; err != nil {
		return errors.New("User not found")
	}

	if err := db.Where("UserID <> ?", userID).Delete(&Files{}).Error; err != nil {
		return err
	}

	return nil
}

func GetAllFiles(userID uint64) ([]Files, error) {
	var u User

	if err := db.First(&u, userID).Error; err != nil {
		return nil, errors.New("User not found")
	}

	var f []Files

	if err := db.Where("UserID <> ?", userID).Find(&f).Error; err != nil {
		return nil, err
	}

	return f, nil
}
