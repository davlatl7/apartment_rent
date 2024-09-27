package service

import (
	"apartment_rent/errs"
	"apartment_rent/models"
	"strconv"
	"apartment_rent/pkg/repository"
	"apartment_rent/utils"
	"errors"
)

func GetAllUsers() (users []models.User, err error) {
	users, err = repository.GetAllUsers()
	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByID(id uint) (user models.User, err error) {
	user, err = repository.GetUserByID(id)
	if err != nil {
		return user, err
	}

	return user, nil
}

func CreateUser(user models.User) error {

	userFromDB, err := repository.GetUserByUsername(user.Username)
	if err != nil && !errors.Is(err, errs.ErrRecordNotFound) {
		return err
	}

	if userFromDB.ID > 0 {
		return errs.ErrUsernameUniquenessFailed
	}

	user.Role = "user"


	user.Password = utils.GenerateHash(user.Password)

	
	err = repository.CreateUser(user)
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	err := repository.DeleteUser(id) 
    if err != nil {
        return err
    }
    return nil
}
func UpdateUser(id int, userUpdate models.UserFilterForUpdate) (*models.User, error) {
	user, err := repository.UpdateUser(id, userUpdate)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func BlockUser(userID string) error {
    id, err := strconv.ParseUint(userID, 10, 32)
    if err != nil {
        return err 
    }

    err = repository.BlockUser(uint(id))
    return err
}