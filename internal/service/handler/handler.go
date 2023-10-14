package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework-4/internal/service/repository"
	"homework-4/internal/service/repository/studentDB"
)

var (
	ErrWrongIdInQuery = errors.New("wrong id in query params")
)

const queryParamKey = "key"

type Server struct {
	Repo *studentDB.StudentsRepo
}

type StudentRequest struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Points int64  `json:"points"`
}

func writeResult(w http.ResponseWriter, result []byte) {
	_, wErr := w.Write(result)
	if wErr != nil {
		log.Printf("Write failed: %v", wErr)
	}
}

func writeError(w http.ResponseWriter, err error) {
	_, wErr := w.Write(errToBytes(err))
	if wErr != nil {
		log.Printf("Write failed: %v", wErr)
	}
}

func errToBytes(err error) []byte {
	return []byte(fmt.Sprintf("%v", err))
}

func CreateRouter(srv Server) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/student", srv.Create).Methods("POST")
	router.HandleFunc("/student", srv.Update).Methods("PUT")

	router.HandleFunc(fmt.Sprintf("/student/{%s:[0-9]*}", queryParamKey), srv.Get).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/student/{%s:[0-9]*}", queryParamKey), srv.Delete).Methods("DELETE")

	return router
}

func (s *Server) Create(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err)
		return
	}
	var unm StudentRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err)
		return
	}
	studentRepo := &repository.Student{
		ID:     unm.Id,
		Name:   unm.Name,
		Points: unm.Points,
	}
	err = s.Repo.Add(req.Context(), studentRepo)
	if err != nil {
		if errors.Is(err, repository.ErrConflict) {
			w.WriteHeader(http.StatusConflict)
			writeError(w, err)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err)
		return
	}
}

func (s *Server) Get(w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, ErrWrongIdInQuery)
		return
	}
	student, err := s.Repo.GetByID(req.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			w.WriteHeader(http.StatusNotFound)
			writeError(w, err)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err)
		return
	}
	studentJson, err := json.Marshal(student)
	if err != nil {
		writeError(w, err)
		return
	}
	writeResult(w, studentJson)
}

func (s *Server) Update(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err)
		return
	}
	var unm StudentRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err)
		return
	}
	studentRepo := &repository.Student{
		ID:     unm.Id,
		Name:   unm.Name,
		Points: unm.Points,
	}
	err = s.Repo.Update(req.Context(), studentRepo)

	if errors.Is(err, repository.ErrObjectNotFound) {
		w.WriteHeader(http.StatusNotFound)
		writeError(w, err)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err)
		return
	}
}
func (s *Server) Delete(w http.ResponseWriter, req *http.Request) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeError(w, ErrWrongIdInQuery)
		return
	}
	err = s.Repo.Delete(req.Context(), keyInt)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			w.WriteHeader(http.StatusNotFound)
			writeError(w, err)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		writeError(w, err)
		return
	}
}
