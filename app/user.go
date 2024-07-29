package app

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"sam/pkg/repository"
	"sam/pkg/user"
	"sam/pkg/util"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *server) getUserHandler() func(w http.ResponseWriter, r *http.Request) {
	userRepository := repository.GetRepository[user.User](s.db)

	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		if _, ok := vars["id"]; !ok {
			s.writeErrorResponse(w, errors.New(util.ID_REQUIRED_ERROR_MESSGE), http.StatusBadRequest)
			return
		}

		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			s.writeErrorResponse(w, errors.New(util.ID_TYPE_ERROR_MESSGE), http.StatusBadRequest)
			return
		}

		entity, err := user.GetUser(id, userRepository)

		if err != nil {
			s.writeErrorResponse(w, err, http.StatusNotFound)
			return
		}

		s.writeResponse(w, entity)
	}
}

func (s *server) registerUserHandler() func(w http.ResponseWriter, r *http.Request) {
	userRepository := repository.GetRepository[user.User](s.db)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			s.writeErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		userRegisterDto := &user.UserRegisterDTO{}
		if err != json.Unmarshal(body, userRegisterDto) {
			s.writeErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		entity, err := user.CreateUser(userRegisterDto, userRepository)

		if err != nil {
			s.writeErrorResponse(w, err, http.StatusInternalServerError)
		}

		s.writeResponse(w, entity)
	}
}

func (s *server) loginUserHandler() func(w http.ResponseWriter, r *http.Request) {
	userRepository := repository.GetRepository[user.User](s.db)

	return func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			s.writeErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		userLoginDTO := &user.UserLoginDTO{}
		if err != json.Unmarshal(body, userLoginDTO) {
			s.writeErrorResponse(w, err, http.StatusBadRequest)
			return
		}

		// validate request
		if userLoginDTO.Username == "" || userLoginDTO.Password == "" {
			s.writeErrorResponse(w, errors.New(util.LOGIN_REQUIRED_FIELDS_ERROR_MESSAGE), http.StatusBadRequest)
			return
		}

		// get user
		u, err := user.GetUserByUsernameAndPlainPassword(
			userLoginDTO.Username,
			userLoginDTO.Password,
			userRepository,
		)

		if err != nil {
			s.writeErrorResponse(w, err, http.StatusNotFound)
			return
		}

		token, err := user.GenerateTokenAndStoreUser(u, s.c.Jwt.Secret, userRepository)

		if err != nil {
			s.writeErrorResponse(w, err, http.StatusInternalServerError)
		}

		s.writeResponse(w, token)
	}
}

func (s *server) refreshTokenHandler() func(w http.ResponseWriter, r *http.Request) {
	// @todo implement

	return func(w http.ResponseWriter, r *http.Request) {
	}
}
