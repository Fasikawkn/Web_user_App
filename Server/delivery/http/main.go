package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/fasikawkn/Web_user_App/Server/delivery/http/handler"
	"github.com/fasikawkn/Web_user_App/Server/users/repository"
	"github.com/fasikawkn/Web_user_App/Server/users/service"
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

	userRepository := repository.NewUserRepository(dbconn)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.GET("/host/user/:id", userHandler.GetSingleUser)
	router.GET("/host/user", userHandler.GetManyUsers)
	router.POST("/host/user", userHandler.AddUser)
	router.PUT("/host/user/:id", userHandler.UpdateUser)
	router.DELETE("/host/user/:id", userHandler.DeleteUser)
	http.ListenAndServe(":8181", router)

}
