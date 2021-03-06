package main

import (
	"io/ioutil"
	"log"

	"github.com/gorilla/mux"
	"github.com/nparavinja/chiron-project/back-legs/db"
	"github.com/nparavinja/chiron-project/back-legs/services"
)

type server struct {
	router             *mux.Router
	patientService     *services.PatientService
	doctorService      *services.DoctorService
	examinationService *services.ExaminationService
	adminService       *services.AdminService
}

func StartServer(config Config) (*server, error) {
	// read pass from db-secret
	binPass, err := ioutil.ReadFile("/run/secrets/db-password")
	if err != nil {
		return nil, err
	}
	dbConnection, err := db.Connect(string(binPass))
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
		patientService:     &services.PatientService{UserRepository: &db.UserRepository{DB: dbConnection}},
		doctorService:      &services.DoctorService{UserRepository: &db.UserRepository{DB: dbConnection}},
		examinationService: &services.ExaminationService{ExaminationRepository: &db.ExaminationRepository{DB: dbConnection}},
		adminService:       &services.AdminService{UserRepository: &db.UserRepository{DB: dbConnection}},
	}
	// init routes
	s.router.HandleFunc("/api/usr/", s.handleUser()).Methods("POST")
	s.router.HandleFunc("/api/exm/", s.handleExamination()).Methods("POST")
	s.router.HandleFunc("/admin/", s.handleAdmin())
	// s.router.HandleFunc("/admin/", logger(JWTMiddleware(s.handleAdmin())))
	log.Println("Services and routers initialized...")
	return s, nil
}
