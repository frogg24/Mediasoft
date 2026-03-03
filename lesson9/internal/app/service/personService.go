package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"mediasoft/lesson9/internal/model"
	"mediasoft/lesson9/internal/repository"
	"net/http"
	"strconv"
	"time"
)

type PersonService struct {
	person repository.PersonRepository
}

func NewPersonService(person repository.PersonRepository) *PersonService {
	return &PersonService{
		person: person,
	}
}

type CreatePersonRequest struct {
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Birthdate time.Time `json:"birthday"`
	GroupID   int64     `json:"groupid"`
}

func (s *PersonService) Create(w http.ResponseWriter, r *http.Request) {
	req := new(CreatePersonRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	if err := s.person.CreatePerson(r.Context(), model.Person{
		Name:      req.Name,
		Lastname:  req.Lastname,
		Birthdate: req.Birthdate,
		GroupID:   req.GroupID,
	}); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusCreated, nil)
	return
}

func (s *PersonService) Get(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	person, err := s.person.ReadPerson(r.Context(), int64(id))
	switch {
	case err == nil:
		response(w, http.StatusOK, person)
	case errors.Is(err, sql.ErrNoRows):
		responseError(w, http.StatusNotFound, err)
	default:
		responseError(w, http.StatusInternalServerError, err)
	}
}

// type GetAllResponse struct {
// 	Results []GetResponse `json:"results"`
// }

// func (s *Service) GetAll(w http.ResponseWriter, r *http.Request) {
// 	employees, err := s.person.List(r.Context())
// 	if err != nil {
// 		responseError(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	result := make([]GetResponse, len(employees))
// 	for i, employee := range employees {
// 		result[i] = GetResponse{
// 			ID:       employee.ID,
// 			Name:     employee.Name,
// 			Surname:  employee.Surname,
// 			Position: employee.Position,
// 		}
// 	}
// 	response(w, http.StatusOK, GetAllResponse{
// 		Results: result,
// 	})
// }

type UpdatePersonResponse struct {
	Name      string    `json:"name"`
	Lastname  string    `json:"lastname"`
	Birthdate time.Time `json:"birthday"`
	GroupID   int64     `json:"groupid"`
}

func (s *PersonService) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	req := new(UpdatePersonResponse)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	if err := s.person.UpdatePerson(r.Context(), model.Person{
		ID:        int64(id),
		Name:      req.Name,
		Lastname:  req.Lastname,
		Birthdate: req.Birthdate,
		GroupID:   req.GroupID,
	}); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

type DeleteRequest struct {
	Id int64 `json:"id"`
}

func (s *PersonService) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.person.DeletePerson(r.Context(), int64(id)); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

func response(w http.ResponseWriter, code int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			log.Println(err)
		}
	}
}

func responseError(w http.ResponseWriter, code int, err error) {
	response(w, code, map[string]string{"error:": err.Error()})
}
