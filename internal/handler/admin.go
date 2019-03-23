package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"sekisan_api/modules/pkg/mod/github.com/gorilla/mux@v1.6.2"

	"sekisan_api/internal/repository"
	"sekisan_api/internal/service"
)

type adminService interface {
	GetAdminList() (aList service.AdminList, err error)
	RegisterAdmin(name, password string) (a repository.Admin, err error)
	UpdateAdminName(id int, name string) (a repository.Admin, err error)
	UpdateAdminPassword(id int, password string) (a repository.Admin, err error)
	UpdateAdminEnabled(id int, enabled int) (a repository.Admin, err error)
}

// Admin handler
type adminHandler struct {
	service adminService
}

func NewAdminHandler(adminService adminService) *adminHandler {
	return &adminHandler{
		service: adminService,
	}
}

func (h *adminHandler) GetAdminList(w http.ResponseWriter, r *http.Request) {
	adminList, err := h.service.GetAdminList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(adminList); err != nil {
		panic(err)
	}
}

func (h *adminHandler) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	name := v["name"]
	password := v["password"]

	a, err := h.service.RegisterAdmin(name, password)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(a); err != nil {
		panic(err)
	}
}

func (h *adminHandler) UpdateAdminName(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	idStr := v["admin_id"]
	name := v["name"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + idStr + "'.")
		badRequest(w)
		return
	}

	a, err := h.service.UpdateAdminName(id, name)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(a); err != nil {
		panic(err)
	}
}

func (h *adminHandler) UpdateAdminPassword(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	idStr := v["admin_id"]
	password := v["password"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + idStr + "'.")
		badRequest(w)
		return
	}

	a, err := h.service.UpdateAdminPassword(id, password)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(a); err != nil {
		panic(err)
	}
}

func (h *adminHandler) UpdateAdminEnabled(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	idStr := v["admin_id"]
	eStr := v["enabled"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + idStr + "'.")
		badRequest(w)
		return
	}

	enabled, err := strconv.Atoi(eStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + idStr + "'.")
		badRequest(w)
		return
	}

	a, err := h.service.UpdateAdminEnabled(id, enabled)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(a); err != nil {
		panic(err)
	}
}
