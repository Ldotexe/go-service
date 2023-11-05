package handler

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"homework-6/internal/message_broker"
	"homework-6/internal/service/repository"
)

var (
	ErrWrongIdInQuery = errors.New("wrong id in query params")
)

const queryParamKey = "key"

type Server struct {
	Repo   repository.StudentRepo
	Sender message_broker.Sender
}

func NewServer(repo repository.StudentRepo, sender message_broker.Sender) *Server {
	return &Server{
		Repo:   repo,
		Sender: sender,
	}
}

type StudentRequest struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Points int64  `json:"points"`
}

func writeResult(w http.ResponseWriter, result []byte) {
	_, wErr := w.Write(result)
	if wErr != nil {
		log.Printf("Write failed: %v", wErr)
	}
}

func writeResponse(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	if err != nil {
		_, wErr := w.Write(errToBytes(err))
		if wErr != nil {
			log.Printf("Write failed: %v", wErr)
		}
	}
}

func errToBytes(err error) []byte {
	return []byte(fmt.Sprintf("%v", err))
}

func CreateRouter(srv Server) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/student", srv.RunCreate).Methods("POST")
	router.HandleFunc("/student", srv.RunUpdate).Methods("PUT")

	router.HandleFunc(fmt.Sprintf("/student/{%s:[0-9]*}", queryParamKey), srv.RunGet).Methods("GET")
	router.HandleFunc(fmt.Sprintf("/student/{%s:[0-9]*}", queryParamKey), srv.RunDelete).Methods("DELETE")

	return router
}

func parseRequestKey(req *http.Request) (int64, int, error) {
	key, ok := mux.Vars(req)[queryParamKey]
	if !ok {
		return 0, http.StatusBadRequest, repository.ErrBadRequest
	}
	keyInt, err := strconv.ParseInt(key, 10, 64)
	if err != nil {
		return 0, http.StatusBadRequest, ErrWrongIdInQuery
	}
	return keyInt, http.StatusOK, nil
}

func parseRequestBody(req *http.Request) (*StudentRequest, int, error) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	var unm StudentRequest
	if err = json.Unmarshal(body, &unm); err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return &unm, http.StatusOK, nil
}

func (s *Server) Create(ctx context.Context, student StudentRequest) (int, error) {
	studentRepo := &repository.Student{
		ID:     student.ID,
		Name:   student.Name,
		Points: student.Points,
	}
	err := s.Repo.Add(ctx, studentRepo)
	if err != nil {
		if errors.Is(err, repository.ErrConflict) {
			return http.StatusConflict, err
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *Server) RunCreate(w http.ResponseWriter, req *http.Request) {
	student, status, err := parseRequestBody(req)
	if status != http.StatusOK {
		writeResponse(w, status, err)
		s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
		return
	}
	status, err = s.Create(req.Context(), *student)
	if status != http.StatusOK {
		writeResponse(w, status, err)
		s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
		return
	}
	s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
}

func (s *Server) Get(ctx context.Context, key int64) ([]byte, int, error) {
	student, err := s.Repo.GetByID(ctx, key)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return nil, http.StatusNotFound, err
		}
		return nil, http.StatusInternalServerError, err
	}
	studentJson, err := json.Marshal(student)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return studentJson, http.StatusOK, nil
}

func (s *Server) RunGet(w http.ResponseWriter, req *http.Request) {
	key, status, err := parseRequestKey(req)
	if status != http.StatusOK {
		writeResponse(w, status, err)
		s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
		return
	}
	result, status, err := s.Get(req.Context(), key)
	if status != http.StatusOK {
		writeResponse(w, status, err)
		s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
		return
	}
	writeResult(w, result)
	s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
}

func (s *Server) Update(ctx context.Context, student StudentRequest) (int, error) {
	studentRepo := &repository.Student{
		ID:     student.ID,
		Name:   student.Name,
		Points: student.Points,
	}
	err := s.Repo.Update(ctx, studentRepo)

	if errors.Is(err, repository.ErrObjectNotFound) {
		return http.StatusNotFound, err
	}
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *Server) RunUpdate(w http.ResponseWriter, req *http.Request) {
	student, status, err := parseRequestBody(req)
	if status != http.StatusOK {
		writeResponse(w, status, err)
		s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
		return
	}
	status, err = s.Update(req.Context(), *student)
	if status != http.StatusOK {
		writeResponse(w, status, err)
		s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
		return
	}
	s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
}

func (s *Server) Delete(ctx context.Context, key int64) (int, error) {
	err := s.Repo.Delete(ctx, key)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return http.StatusNotFound, err
		}
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func (s *Server) RunDelete(w http.ResponseWriter, req *http.Request) {
	key, status, err := parseRequestKey(req)
	if status != http.StatusOK {
		writeResponse(w, status, err)
		s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
		return
	}
	status, err = s.Delete(req.Context(), key)
	if status != http.StatusOK {
		writeResponse(w, status, err)
		s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
		return
	}
	s.Sender.SendMessage(message_broker.NewMessageInfo(req, status))
}
