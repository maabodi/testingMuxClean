package handler

import (
	"cobaclean/internal/config"
	"cobaclean/internal/database"
	user "cobaclean/internal/handler/controller/user_controller"
	"cobaclean/internal/repository"
	"cobaclean/internal/service"

	"github.com/gorilla/mux"
)

func ConfigureRouter(conf config.Config) *mux.Router {

	db := database.ConnectDatabase(conf)
	repo := repository.NewUserMysql(db)

	SVC := service.NewServiceUser(repo, conf)

	cont := user.UserController{
		Service: SVC,
	}

	router := mux.NewRouter()
	router.HandleFunc("/api/users", cont.Index).Methods("GET")
	// router.HandleFunc("/api/user/:id", cont.Show).Methods("GET")

	return router
}
