package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/fasikawkn/Web_user_App/Server/entity"
)

//UserRepository implements the repository actitvities
type UserRepository struct {
	conn *sql.DB
}

//NewUserRepository returns a new UserRepository
func NewUserRepository(Conn *sql.DB) *UserRepository {
	return &UserRepository{conn: Conn}
}

//GetSingleUser returns a single user from the database
func (usrRepo *UserRepository) GetSingleUser(id int) (*entity.User, error) {
	row := usrRepo.conn.QueryRow("SELECT * FROM users where userid = $1", id)

	user := entity.User{}

	err := row.Scan(&user.ID, &user.FullName, &user.UserName, &user.Password, &user.Picture, &user.Address)

	if err != nil {
		return nil, err
	}
	return &user, nil

}

//GetManyUsers returns Many users from the database
func (usrRepo *UserRepository) GetManyUsers() ([]entity.User, error) {
	log.Println("returning many users")

	rows, err := usrRepo.conn.Query("SELECT * FROM users")
	if err != nil {
		return nil, errors.New("can't query database")
	}

	users := []entity.User{}

	for rows.Next() {
		user := entity.User{}
		err = rows.Scan(&user.ID, &user.FullName, &user.UserName, &user.Password, &user.Picture, &user.Address)
		if err != nil {
			fmt.Println("error hanppern")
			return nil, err
		}
		users = append(users, user)
	}
	log.Println("users", users)

	return users, nil

}

//AddUser adds a new user to the database
func (usrRepo *UserRepository) AddUser(user *entity.User) (*entity.User, error) {
	log.Println("posting a user")
	_, err := usrRepo.conn.Exec("INSERT into users (fullname,username,password,picture,address) values($1,$2,$3,$4,$5)", user.FullName, user.UserName, user.Password, user.Picture, user.Address)
	if err != nil {
		return nil, err
	}
	return user, nil

}

//DeleteUser deletes a specific id from the database using id
func (usrRepo *UserRepository) DeleteUser(id int) error {
	_, err := usrRepo.conn.Exec("DELETE from users where userid = $1", id)

	if err != nil {
		return err
	}
	return nil
}

//UpdateUser updates a single user from the database
func (usrRepo *UserRepository) UpdateUser(user *entity.User) (*entity.User, error) {
	fmt.Println("updating")
	_, err := usrRepo.conn.Exec("UPDATE users set fullname = $1,username = $2, password = $3,picture = $4,address = $5 where userid = $6", user.FullName, user.UserName, user.Password, user.Picture, user.Address, user.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
