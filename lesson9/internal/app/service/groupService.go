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
)

type GroupService struct {
	group repository.GroupRepository
}

func NewGroupService(group repository.GroupRepository) *GroupService {
	return &GroupService{
		group: group,
	}
}

type CreateGroupRequest struct {
	Title       string `json:"title"`
	ParentGroup *int64 `json:"parentgroup"`
}

func (s *GroupService) Create(w http.ResponseWriter, r *http.Request) {
	req := new(CreateGroupRequest)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	if err := s.group.CreateGroup(r.Context(), model.Group{
		Title:       req.Title,
		ParentGroup: req.ParentGroup,
	}); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusCreated, nil)
	return
}

func (s *GroupService) Get(w http.ResponseWriter, r *http.Request) {
	log.Println("METHOD:", r.Method, "PATH:", r.URL.Path)
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	group, err := s.group.ReadGroup(r.Context(), int64(id))
	switch {
	case err == nil:
		response(w, http.StatusOK, group)
	case errors.Is(err, sql.ErrNoRows):
		responseError(w, http.StatusNotFound, err)
	default:
		responseError(w, http.StatusInternalServerError, err)
	}
}

type UpdateGroupResponse struct {
	Title       string `json:"title"`
	ParentGroup *int64 `json:"parentgroup"`
}

func (s *GroupService) Update(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	req := new(UpdateGroupResponse)
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}
	r.Body.Close()

	if err := s.group.UpdateGroup(r.Context(), model.Group{
		ID:          int64(id),
		Title:       req.Title,
		ParentGroup: req.ParentGroup,
	}); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

type DeleteGroupRequest struct {
	Id int64 `json:"id"`
}

func (s *GroupService) Delete(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	id, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	if err := s.group.DeleteGroup(r.Context(), int64(id)); err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusNoContent, nil)
}

func (s *GroupService) ListPersonLocal(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	groupID, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	persons, err := s.group.ListPersonLocal(r.Context(), int64(groupID))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, map[string][]model.Person{"results": persons})
}

func (s *GroupService) ListPersonAll(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	groupID, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	persons, err := s.group.ListPersonAll(r.Context(), int64(groupID))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, map[string][]model.Person{"results": persons})
}

func (s *GroupService) CountGroupLocal(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	groupID, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	count, err := s.group.CountGroupLocal(r.Context(), int64(groupID))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, map[string]int64{"count": count})
}

func (s *GroupService) CountGroupAll(w http.ResponseWriter, r *http.Request) {
	idString := r.PathValue("id")
	groupID, err := strconv.Atoi(idString)
	if err != nil {
		responseError(w, http.StatusBadRequest, err)
		return
	}

	count, err := s.group.CountGroupAll(r.Context(), int64(groupID))
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, map[string]int64{"count": count})
}

func (s *GroupService) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	log.Println("METHOD:", r.Method, "PATH:", r.URL.Path)

	groups, err := s.group.GetAllGroups(r.Context())
	if err != nil {
		responseError(w, http.StatusInternalServerError, err)
		return
	}

	response(w, http.StatusOK, map[string][]model.Group{"results": groups})
}
