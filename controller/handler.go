package controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"sekisan_api/model"
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
	enabled := v["enabled"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Printf("[INFO] Can't cast to int '" + idStr + "'.")
		badRequest(w)
		return
	}

	a, err := updateAdminPassword(id, enabled)
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


// Sekisan handler
func (h *Handler) GetSekisan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sekisan_id := vars["emp_id"]
	if sekisan_id == "" {
		log.Printf("[INFO] sekisan is null. ?", sekisan_id)
		badRequest(w)
		return
	}

	sek, err_r := model.GetSekisanByEmployeeNum(h.db, sekisan_id)
	if err_r != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	sekisan := model.Sekisan{
		ID:          sek.ID,
		EmployeeNum: sek.EmployeeNum,
		Sekisan:     sek.Sekisan,
	}

	if err := json.NewEncoder(w).Encode(sekisan); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllSekisan(w http.ResponseWriter, r *http.Request) {
	sek_all, err_r := model.GetAllSekisan(h.db)
	if err_r != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	var sekisan_list []model.Sekisan
	for _, sek := range sek_all {
		sek_i := model.Sekisan{
			ID:          sek.ID,
			EmployeeNum: sek.EmployeeNum,
			Sekisan:     sek.Sekisan,
		}
		sekisan_list = append(sekisan_list, sek_i)
	}

	type AllSekisan struct {
		Sekisan []model.Sekisan
	}
	all_sekisan := AllSekisan{sekisan_list}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(all_sekisan); err != nil {
		panic(err)
	}
}
