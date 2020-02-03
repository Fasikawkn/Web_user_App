package service

import (
	"fmt"

	"github.com/fasikawkn/Web_user_App/Server/entity"
	"github.com/fasikawkn/Web_user_App/Server/places/repository"
)

//PlaceService implements the repository actitvities
type PlaceService struct {
	placeRepo *repository.PlaceRepository
}

//NewPlaceService returns a new UserRepository
func NewPlaceService(repo *repository.PlaceRepository) *PlaceService {
	return &PlaceService{placeRepo: repo}
}

//GetSinglePlace returns a single place from the database
func (placeSrv *PlaceService) GetSinglePlace(id int) (*entity.Place, error) {
	place, err := placeSrv.placeRepo.GetSinglePlace(id)
	if err != nil {
		return nil, err
	}
	return place, nil

}

//GetManyPlaces returns Many places from the database by userID
func (placeSrv *PlaceService) GetManyPlaces(userid int) ([]entity.Place, error) {
	places, err := placeSrv.placeRepo.GetManyPlaces(userid)
	if err != nil {
		return nil, err

	}
	return places, nil

}

//AddPlace adds a new user to the database
func (placeSrv *PlaceService) AddPlace(place *entity.Place) (*entity.Place, error) {
	place, err := placeSrv.placeRepo.AddPlace(place)
	if err != nil {
		return nil, err
	}
	return place, nil

}

//DeletePlace deletes a specific id from the database using id
func (placeSrv *PlaceService) DeletePlace(id int) error {
	err := placeSrv.placeRepo.DeletePlace(id)

	if err != nil {
		return err
	}
	return nil
}

//UpdatePlace updates a single user from the database
func (placeSrv *PlaceService) UpdatePlace(place *entity.Place) (*entity.Place, error) {
	fmt.Println("updating")
	place, err := placeSrv.placeRepo.UpdatePlace(place)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return place, nil
}
