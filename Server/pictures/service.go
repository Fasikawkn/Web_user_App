package pictures

import "github.com/fasikawkn/Web_user_App/Server/entity"

//PictureService  ...
type PictureService interface {
	GetSinglePicture(id int) (*entity.Picture, error)
	GetManyPictures(userID int) ([]entity.Picture, error)
	AddPicture(place *entity.Picture) (*entity.Picture, error)
	UpdatePicture(place *entity.Picture) (*entity.Picture, error)
	DeletePicture(id int) error
}
