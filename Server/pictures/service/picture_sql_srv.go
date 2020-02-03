package service

import (
	"fmt"

	"github.com/fasikawkn/Web_user_App/Server/entity"
	"github.com/fasikawkn/Web_user_App/Server/pictures/repository"
)

//PictureService implements the repository actitvities
type PictureService struct {
	picRepo *repository.PictureRepository
}

//NewPictureService returns a new UserRepository
func NewPictureService(repo *repository.PictureRepository) *PictureService {
	return &PictureService{picRepo: repo}
}

//GetSinglePicture returns a single place from the database
func (piceSrv *PictureService) GetSinglePicture(id int) (*entity.Picture, error) {
	place, err := piceSrv.picRepo.GetSinglePicture(id)
	if err != nil {
		return nil, err
	}
	return place, nil

}

//GetManyPictures returns Many places from the database by userID
func (piceSrv *PictureService) GetManyPictures(placeID int) ([]entity.Picture, error) {
	pictures, err := piceSrv.picRepo.GetManyPictures(placeID)
	if err != nil {
		return nil, err

	}
	return pictures, nil

}

//AddPicture adds a new user to the database
func (piceSrv *PictureService) AddPicture(picture *entity.Picture) (*entity.Picture, error) {
	picture, err := piceSrv.picRepo.AddPicture(picture)
	if err != nil {
		return nil, err
	}
	return picture, nil

}

//DeletePicture deletes a specific id from the database using id
func (piceSrv *PictureService) DeletePicture(id int) error {
	err := piceSrv.picRepo.DeletePicture(id)

	if err != nil {
		return err
	}
	return nil
}

//UpdatePicture updates a single user from the database
func (piceSrv *PictureService) UpdatePicture(picture *entity.Picture) (*entity.Picture, error) {
	fmt.Println("updating")
	picture, err := piceSrv.picRepo.UpdatePicture(picture)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return picture, nil
}
