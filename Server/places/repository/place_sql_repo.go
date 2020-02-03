package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/fasikawkn/Web_user_App/Server/entity"
)

//PlaceRepository implements the repository actitvities
type PlaceRepository struct {
	conn *sql.DB
}

//NewPlaceRepository returns a new UserRepository
func NewPlaceRepository(Conn *sql.DB) *PlaceRepository {
	return &PlaceRepository{conn: Conn}
}

//GetSinglePlace returns a single place from the database
func (usrRepo *PlaceRepository) GetSinglePlace(id int) (*entity.Place, error) {
	row := usrRepo.conn.QueryRow("SELECT * FROM places where id = $1", id)

	place := entity.Place{}

	err := row.Scan(&place.ID, &place.UserID, &place.Name, &place.Location, &place.Description)

	if err != nil {
		return nil, err
	}
	return &place, nil

}

//GetManyPlaces returns Many places from the database by userID
func (usrRepo *PlaceRepository) GetManyPlaces(userid int) ([]entity.Place, error) {
	log.Println("returning many users")

	rows, err := usrRepo.conn.Query("SELECT * FROM places where userid = $1", userid)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("can't query database")
	}

	places := []entity.Place{}

	for rows.Next() {
		place := entity.Place{}
		err = rows.Scan(&place.ID, &place.UserID, &place.Name, &place.Location, &place.Description)
		if err != nil {
			fmt.Println("error hanppern")
			return nil, err
		}
		places = append(places, place)
	}
	log.Println("users", places)

	return places, nil

}

//AddPlace adds a new user to the database
func (usrRepo *PlaceRepository) AddPlace(place *entity.Place) (*entity.Place, error) {
	log.Println("posting a user")
	_, err := usrRepo.conn.Exec("INSERT into places (userid,name,location,description) values($1,$2,$3,$4)", place.UserID, place.Name, place.Location, place.Description)
	if err != nil {
		return nil, err
	}
	return place, nil

}

//DeletePlace deletes a specific id from the database using id
func (usrRepo *PlaceRepository) DeletePlace(id int) error {
	_, err := usrRepo.conn.Exec("DELETE from places where id = $1", id)

	if err != nil {
		return err
	}
	return nil
}

//UpdatePlace updates a single user from the database
func (usrRepo *PlaceRepository) UpdatePlace(place *entity.Place) (*entity.Place, error) {
	fmt.Println("updating")
	_, err := usrRepo.conn.Exec("UPDATE places set name = $1,location = $2, description = $3 where id = $4", place.Name, place.Location, place.Description, place.ID)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return place, nil
}
