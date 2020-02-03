package users

import "github.com/fasikawkn/Web_user_App/Server/entity"

//UserRepository contains user repositoy actions
type UserRepository interface {
	GetSingleUser(id int) (*entity.User, error)
	GetManyUsers() ([]entity.User, error)
	AddUser(user *entity.User) (*entity.User, error)
	UpdateUser(user *entity.User) (*entity.User, error)
	DeleteUser(id int) error
}
