package repository

import (
	"apartment_rent/db"
	"apartment_rent/logger"
	"apartment_rent/models"
)

func GetAllUsers() (users []models.User, err error) {
	err = db.GetDBConn().Find(&users).Error
	if err != nil {
		logger.Error.Printf("[repository.GetAllUsers] error getting all users: %s\n", err.Error())
		return nil, translateError(err)
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	err = db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByID] error getting user by id: %v\n", err)
		return user, translateError(err)
	}
	return user, nil
}

func GetUserByUsername(username string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ?", username).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsername] error getting user by username: %v\n", err)
		return user, translateError(err)
	}

	return user, nil
}

func GetUserByUsernameAndPassword(username string, password string) (user models.User, err error) {
	err = db.GetDBConn().Where("username = ? AND password = ?", username, password).First(&user).Error
	if err != nil {
		logger.Error.Printf("[repository.GetUserByUsernameAndPassword] error getting user by username and password: %v\n", err)
		return user, translateError(err)
	}

	return user, nil
}

func CreateUser(user models.User) (err error) {
	if err = db.GetDBConn().Create(&user).Error; err != nil {
		logger.Error.Printf("[repository.CreateUser] error creating user: %v\n", err)
		return translateError(err)
	}

	return nil
}
func DeleteUser(id int) error {
	var user models.User
	err := db.GetDBConn().Model(&user).Where("id = ?",id).Update("is_deleted", true).Error
	return err
}

func UpdateUser(id int, userUpdate models.UserFilterForUpdate) (*models.User, error) {
	updates := map[string]interface{}{}

	if userUpdate.FullName != "" {
		updates["full_name"] = userUpdate.FullName
	}
	if userUpdate.Username != "" {
		updates["username"] = userUpdate.Username
	}
	
	if len(updates) > 0 {
		err := db.GetDBConn().Model(&models.User{}).Where("id = ?", id).Updates(updates).Error
		if err != nil {
			return nil, err
		}
	}

	var user models.User
	err := db.GetDBConn().Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func BlockUser(userID uint) error {
    var user models.User
    err := db.GetDBConn().Model(&user).Where("id = ?", userID).Update("is_blocked", true).Error
    return err
}