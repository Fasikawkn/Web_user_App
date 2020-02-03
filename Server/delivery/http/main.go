package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fasikawkn/Web_user_App/Server/delivery/http/handler"
	"github.com/fasikawkn/Web_user_App/Server/users/repository"
	"github.com/fasikawkn/Web_user_App/Server/users/service"

	placeHdlr "github.com/fasikawkn/Web_user_App/Server/delivery/http/handler"
	placeRepo "github.com/fasikawkn/Web_user_App/Server/places/repository"
	placeSrv "github.com/fasikawkn/Web_user_App/Server/places/service"

	picHdlr "github.com/fasikawkn/Web_user_App/Server/delivery/http/handler"
	picRepo "github.com/fasikawkn/Web_user_App/Server/pictures/repository"
	picSrv "github.com/fasikawkn/Web_user_App/Server/pictures/service"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
)

func main() {

	dbconn, err := sql.Open("postgres", "postgres://postgres:user1@localhost/hostdb?sslmode=disable")

	if err != nil {
		log.Println("connection refused!")
		panic(err)
		return
	}
	log.Println("database connection established")
	router := httprouter.New()

	// Users
	userRepository := repository.NewUserRepository(dbconn)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/host/user/:id", userHandler.GetSingleUser)
	router.GET("/host/user", userHandler.GetManyUsers)
	router.POST("/host/user", userHandler.AddUser)
	router.PUT("/host/user/:id", userHandler.UpdateUser)
	router.DELETE("/host/user/:id", userHandler.DeleteUser)

	// Places

	placeRepository := placeRepo.NewPlaceRepository(dbconn)
	placeService := placeSrv.NewPlaceService(placeRepository)
	placeHandler := placeHdlr.NewPlacePHandler(placeService)

	router.GET("/host/place/:id", placeHandler.GetSinglePlace)
	router.GET("/host/places/:userid", placeHandler.GetManyPlaces)
	router.POST("/host/place", placeHandler.AddPlace)
	router.PUT("/host/place/:id", placeHandler.UpdatePlace)
	router.DELETE("/host/place/:id", placeHandler.DeletePlace)

	//Pictures
	picRepository := picRepo.NewPictureRepository(dbconn)
	picService := picSrv.NewPictureService(picRepository)
	picHandler := picHdlr.NewPictureHandler(picService)

	router.GET("/host/picture/:id", picHandler.GetSinglePicture)
	router.GET("/host/pictures/:placeid", picHandler.GetManyPictures)
	router.POST("/host/picture", picHandler.AddPicture)
	router.PUT("/host/picture/:id", picHandler.UpdatePicture)
	router.DELETE("/host/picture/:id", picHandler.DeletePicture)
	

	http.ListenAndServe(":8181", router)

}
