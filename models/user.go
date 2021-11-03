package models

import "gorm.io/gorm"

type User struct { // User schema or blueprint
	gorm.Model
	ID    int    `json:"ID"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// create user
func CreateUser(db *gorm.DB, User *User) (err error) {
	err = db.Create(User).Error
	if err != nil {
		return err
	}

	return nil
}

// get users
func GetUsers(db *gorm.DB, User *[]User) (err error) {
	err = db.Find(User).Error
	if err != nil {
		return err
	}

	return nil
}

// get user by id
func GetUserById(db *gorm.DB, User *User, id string) (err error) {
	err = db.Where("id = ?", id).First(User).Error
	if err != nil {
		return err
	}

	return nil
}

// update user
func UpdateUser(db *gorm.DB, User *User) (err error) {
	db.Save(User)

	return nil
}

// delete user
func DeleteUser(db *gorm.DB, User *User, id string) (err error) {
	db.Where("id = ?", id).Delete(User)

	return nil
}
