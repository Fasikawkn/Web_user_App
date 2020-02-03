package service

import (
	"github.com/fasikawkn/Web_user_App/Server/entity"
	"github.com/fasikawkn/Web_user_App/Server/users/repository"
)

//UserServices implements the userservice interfaces
type UserServices struct {
	repo *repository.UserRepository
}

//NewUserService returns new UserService
func NewUserService(repoC *repository.UserRepository) *UserServices {
	return &UserServices{repo: repoC}
}

//GetSingleUser returns a single user by id
func (usrSrv *UserServices) GetSingleUser(id int) (*entity.User, error) {
	user, err := usrSrv.repo.GetSingleUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//GetManyUsers returs all the users in the database
func (usrSrv *UserServices) GetManyUsers() ([]entity.User, error) {
	users, err := usrSrv.repo.GetManyUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

//AddUser adds a new user to the database
func (usrSrv *UserServices) AddUser(user *entity.User) (*entity.User, error) {
	user, err := usrSrv.repo.AddUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

//UpdateUser updates a single user
func (usrSrv *UserServices) UpdateUser(user *entity.User) (*entity.User, error) {
	user, err := usrSrv.repo.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return user, nil

}

//DeleteUser deletes a single user by id
func (usrSrv *UserServices) DeleteUser(id int) error {

	err := usrSrv.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil

}
