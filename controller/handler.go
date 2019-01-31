package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"strconv"
)

const (
	SessionSecret = "testtest"
	SessionName   = "test_session"
)

//var BaseTime time.Time
var (
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	store = sessions.NewCookieStore([]byte(SessionSecret))
)

// Correct response handlers.
type Handler struct {
	store sessions.Store
}

func NewHandler() *Handler {
	//BaseTime = time.Date(2018, 10, 16, 10, 0, 0, 0, time.Local)
	return &Handler{
		store: store,
	}
}

// 404 handler
func (h *Handler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = w.Write([]byte("Not Found\n"))
}

// Admin handler
func (h *Handler) GetAdminList(w http.ResponseWriter, r *http.Request) {
	adminList, err := getAdminList()
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

func (h *Handler) RegisterAdmin(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	name := v["name"]
	password := v["password"]

	a, err := registerAdmin(name, password)
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

func (h *Handler) UpdateAdminName(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	idStr := v["admin_id"]
	name := v["name"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + idStr + "'.")
		badRequest(w)
		return
	}

	a, err := updateAdminName(id, name)
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

func (h *Handler) UpdateAdminPassword(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	idStr := v["admin_id"]
	password := v["password"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + idStr + "'.")
		badRequest(w)
		return
	}

	a, err := updateAdminPassword(id, password)
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

func (h *Handler) UpdateAdminEnabled(w http.ResponseWriter, r *http.Request) {
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

	a, err := updateAdminEnabled(id, enabled)
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

// Member handler
func (h *Handler) GetMemberList(w http.ResponseWriter, r *http.Request) {
	memberList, err := getMemberList()
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

func (h *Handler) RegisterMember(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	employeeNumStr := v["employee_num"]
	name := v["name"]

	employeeNum, err := strconv.Atoi(employeeNumStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + employeeNumStr + "'.")
		badRequest(w)
		return
	}

	m, err := registerMember(employeeNum, name)
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

func (h *Handler) UpdateMemberName(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	employeeNumStr := v["employee_num"]
	name := v["name"]

	employeeNum, err := strconv.Atoi(employeeNumStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + employeeNumStr + "'.")
		badRequest(w)
		return
	}

	m, err := updateMemberName(employeeNum, name)
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

func (h *Handler) UpdateMemberEnabled(w http.ResponseWriter, r *http.Request) {
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

	a, err := updateMemberEnabled(employeeNum, enabled)
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

// Sekisan handler
func (h *Handler) GetAllSekisan(w http.ResponseWriter, r *http.Request) {
	sekisanRes, err := getSekisanList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sekisanRes); err != nil {
		panic(err)
	}
}

// Transaction handler
func (h *Handler) GetTransactionList(w http.ResponseWriter, r *http.Request) {
	transactionList, err := getTransactionList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(transactionList); err != nil {
		panic(err)
	}
}

func (h *Handler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	employeeNumStr := v["employee_num"]
	hourStr := v["hour"]
	operation := v["operation"]
	reason := v["reason"]

	employeeNum, err := strconv.Atoi(employeeNumStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + employeeNumStr + "'.")
		badRequest(w)
		return
	}

	hour, err := strconv.Atoi(hourStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + employeeNumStr + "'.")
		badRequest(w)
		return
	}

	a, err := addTransaction(employeeNum, hour, operation, reason)
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
