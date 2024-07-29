package app

import (
	"github.com/gorilla/handlers"
)

func (s *server) routes() {
	s.r.Use(s.jsonMiddleware)
	s.r.Use(handlers.RecoveryHandler())

	s.r.HandleFunc("/user/{id}", s.getUserHandler()).Methods("GET")
	s.r.HandleFunc("/register", s.registerUserHandler()).Methods("POST")
	s.r.HandleFunc("/login", s.loginUserHandler()).Methods("POST")
	s.r.HandleFunc("/refresh-token", s.refreshTokenHandler()).Methods("POST")
}
