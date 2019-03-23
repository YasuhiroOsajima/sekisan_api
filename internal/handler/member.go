package handler

import (
	"encoding/json"
	"sekisan_api/modules/pkg/mod/github.com/gorilla/mux@v1.6.2"
	"log"
	"net/http"
	"sekisan_api/internal/repository"
	"strconv"

	"sekisan_api/internal/service"
)

type memberService interface {
	GetMemberList() (mList service.MemberList, err error)
	RegisterMember(employeeNum int, name string) (m repository.Member, err error)
	UpdateMemberName(employeeNum int, name string) (m repository.Member, err error)
	UpdateMemberEnabled(employeeNum, enabled int) (m repository.Member, err error)
}

type memberHandler struct {
	service memberService
}

func NewMemberHandler(memberService memberService) *memberHandler {
	return &memberHandler{
		service: memberService,
	}
}

// Member handler
func (h *memberHandler) GetMemberList(w http.ResponseWriter, r *http.Request) {
	memberList, err := h.service.GetMemberList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(memberList); err != nil {
		panic(err)
	}
}

func (h *memberHandler) RegisterMember(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	employeeNumStr := v["employee_num"]
	name := v["name"]

	employeeNum, err := strconv.Atoi(employeeNumStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + employeeNumStr + "'.")
		badRequest(w)
		return
	}

	m, err := h.service.RegisterMember(employeeNum, name)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func (h *memberHandler) UpdateMemberName(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	employeeNumStr := v["employee_num"]
	name := v["name"]

	employeeNum, err := strconv.Atoi(employeeNumStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + employeeNumStr + "'.")
		badRequest(w)
		return
	}

	m, err := h.service.UpdateMemberName(employeeNum, name)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(m); err != nil {
		panic(err)
	}
}

func (h *memberHandler) UpdateMemberEnabled(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	employeeNumStr := v["employee_num"]
	enabledStr := v["enabledStr"]

	employeeNum, err := strconv.Atoi(employeeNumStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + employeeNumStr + "'.")
		badRequest(w)
		return
	}

	enabled, err := strconv.Atoi(enabledStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + enabledStr + "'.")
		badRequest(w)
		return
	}

	a, err := h.service.UpdateMemberEnabled(employeeNum, enabled)
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
