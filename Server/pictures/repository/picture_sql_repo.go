package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/fasikawkn/Web_user_App/Server/entity"
)

//PictureRepository implements the repository actitvities
type PictureRepository struct {
	conn *sql.DB
}

//NewPictureRepository returns a new UserRepository
func NewPictureRepository(Conn *sql.DB) *PictureRepository {
	return &PictureRepository{conn: Conn}
}

//GetSinglePicture returns a single place from the database
func (usrRepo *PictureRepository) GetSinglePicture(id int) (*entity.Picture, error) {
	row := usrRepo.conn.QueryRow("SELECT * FROM pictures where id = $1", id)

	picture := entity.Picture{}

	err := row.Scan(&picture.ID, &picture.PlaceID, &picture.Name)

	if err != nil {
		return nil, err
	}
	return &picture, nil

}

//GetManyPictures returns Many places from the database by userID
func (usrRepo *PictureRepository) GetManyPictures(placeID int) ([]entity.Picture, error) {
	log.Println("returning many users")

	rows, err := usrRepo.conn.Query("SELECT * FROM pictures where placeid = $1", placeID)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("can't query database")
	}

	pictures := []entity.Picture{}

	for rows.Next() {
		picture := entity.Picture{}
		err = rows.Scan(&picture.ID, &picture.PlaceID, &picture.Name)
		if err != nil {
			fmt.Println("error hanppern")
			return nil, err
		}
		pictures = append(pictures, picture)
	}
	log.Println("users", pictures)

	return pictures, nil

}

//AddPicture adds a new user to the database
func (usrRepo *PictureRepository) AddPicture(picture *entity.Picture) (*entity.Picture, error) {
	log.Println("posting a picture")
	_, err := usrRepo.conn.Exec("INSERT into pictures (placeid,name) values($1,$2)", &picture.PlaceID, &picture.Name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return picture, nil

}

//DeletePicture deletes a specific id from the database using id
func (usrRepo *PictureRepository) DeletePicture(id int) error {
	_, err := usrRepo.conn.Exec("DELETE from pictures where id = $1", id)

	if err != nil {
		return err
	}
	return nil
}

//UpdatePicture updates a single user from the database
func (usrRepo *PictureRepository) UpdatePicture(picture *entity.Picture) (*entity.Picture, error) {
	fmt.Println("updating")
	_, err := usrRepo.conn.Exec("UPDATE pictures set name = $1where id = $2", picture.Name, picture.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return picture, nil
}
