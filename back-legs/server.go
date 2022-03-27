package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nparavinja/chiron-project/back-legs/db"
	"github.com/nparavinja/chiron-project/back-legs/services"
)

type server struct {
	router             *mux.Router
	userService        *services.UserService
	examinationService *services.ExaminationService
}

func StartServer(config Config) (*server, error) {
	// init db
	// test
	dbConnection, err := db.Connect(config.SQLConnectionString)
	if err != nil {
		// handle db errors
		// log
		log.Println(err)
		return nil, err
	}
	log.Println("Database connection enabled...")
	// init services
	s := &server{
		router:             mux.NewRouter(),
		userService:        &services.UserService{UserRepository: &db.UserRepository{DB: dbConnection}},
		examinationService: &services.ExaminationService{ExaminationRepository: &db.ExaminationRepository{DB: dbConnection}},
	}
	// init routes
	s.router.HandleFunc("/api/usr/", s.handleUser()).Methods("POST")
	s.router.HandleFunc("/api/exm/", s.handleExamination()).Methods("POST")
	s.router.HandleFunc("/admin/", s.handleAdmin())
	log.Println("Services and routers initialized...")
	return s, nil
}

func (s *server) middleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// some jwt check

		// next handler (w, r)
	}
}
