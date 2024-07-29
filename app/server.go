package app

import (
	"encoding/json"
	"net/http"
	"sam/pkg/user"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

const CTX_DEFAULT_REPO_KEY string = "repo"

type server struct {
	r  *mux.Router
	db *gorm.DB
	c  *config
}

func NewServer() *server {
	s := server{
		r: mux.NewRouter(),
		c: loadConfig(),
	}

	s.database()
	s.routes()

	return &s
}

func (s *server) Run() error {
	return http.ListenAndServe(s.c.Server.Host+":"+s.c.Server.Port, handlers.RecoveryHandler()(s.r))
}

func (s *server) Migrate() error {
	return s.db.AutoMigrate(user.User{})
}

func (s *server) writeResponse(w http.ResponseWriter, response interface{}) {
	w.WriteHeader(http.StatusOK)

	byteResponse, _ := json.Marshal(response)
	_, _ = w.Write(byteResponse)
}

func (s *server) writeErrorResponse(w http.ResponseWriter, err error, errorCode int) {
	response := make(map[string]interface{})

	response["error"] = err.Error()
	w.WriteHeader(errorCode)
	byteResponse, _ := json.Marshal(response)
	_, _ = w.Write(byteResponse)
}
