package service

import (
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"mediasoft/lesson8/internal/rest/model"
	"mediasoft/lesson8/internal/rest/repository"
	"net/http"
	"strconv"
)

type Service struct {
	employee repository.EmployeeRepository
}

func NewService(employee repository.EmployeeRepository) *Service {
	return &Service{
		employee: employee,
	}
}

type CreateRequest struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Position string `json:"position"`
}

func (s *Service) Create(w http.ResponseWriter, r *http.Request) {
	req := new(CreateRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	if err := s.employee.Create(r.Context(), model.Employee{
		Name:     req.Name,
		Surname:  req.Surname,
		Position: req.Position,
	}); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusCreated, nil)
	return
}

type GetResponse struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Position string `json:"position"`
}

func (s *Service) Get(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	employee, err := s.employee.Read(r.Context(), int64(id))
	switch {
	case err == nil:
		response(w, http.StatusOK, employee)
	case errors.Is(err, sql.ErrNoRows):
		responseError(w, http.StatusNotFound, err)
	default:
		responseError(w, http.StatusInternalServerError, err)
	}
}

type GetAllResponse struct {
	Results []GetResponse `json:"results"`
}

func (s *Service) GetAll(w http.ResponseWriter, r *http.Request) {
	employees, err := s.employee.List(r.Context())
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}
	result := make([]GetResponse, len(employees))
	for i, employee := range employees {
		result[i] = GetResponse{
			ID:       employee.ID,
			Name:     employee.Name,
			Surname:  employee.Surname,
			Position: employee.Position,
		}
	}
	response(w, http.StatusOK, GetAllResponse{
		Results: result,
	})
}

type UpdateResponse struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Position string `json:"position"`
}

func (s *Service) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	req := new(UpdateResponse)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	if err := s.employee.Update(r.Context(), model.Employee{
		ID:       int64(id),
		Name:     req.Name,
		Surname:  req.Surname,
		Position: req.Position,
	}); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

type DeleteRequest struct {
	Id int64 `json:"id"`
}

func (s *Service) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.employee.Delete(r.Context(), int64(id)); err != nil {
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
