package places

import "github.com/fasikawkn/Web_user_App/Server/entity"

//PlaceService ...
type PlaceService interface {
	GetSinglePlace(id int) (*entity.Place, error)
	GetManyPlaces(placeID int) ([]entity.Place, error)
	AddPlace(place *entity.Place) (*entity.Place, error)
	UpdatePlace(place *entity.Place) (*entity.Place, error)
	DeletePlace(id int) error
}
