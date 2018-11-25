package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"sekisan_api/model"
	"time"
)

const (
	SessionName = "test_session"
)

var BaseTime time.Time
var db *sql.DB

type Handler struct {
	db    *sql.DB
	store sessions.Store
}

//type Handler struct {
//	store sessions.Store
//}

// func NewHandler(store sessions.Store) *Handler {
// 	BaseTime = time.Date(2018, 10, 16, 10, 0, 0, 0, time.Local)
// 	return &Handler{
// 		store: store,
// 	}
// }

func NewHandler(db *sql.DB, store sessions.Store) *Handler {
	BaseTime = time.Date(2018, 10, 16, 10, 0, 0, 0, time.Local)
	return &Handler{
		db:    db,
		store: store,
	}
}

func (h *Handler) GetSekisan(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sekisan_id := vars["id"]
	if sekisan_id == "" {
		log.Printf("[INFO] sekisan is null.")
		log.Printf(sekisan_id)
		badRequest(w)
		return
	}

	sek, err_r := model.GetSekisanByEmployeeNum(h.db, sekisan_id)
	if err_r != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}
	log.Printf("?", sek)

	w.Header().Set("Content-Type", "application/json")

	sekisan := model.Sekisan{
		ID: sek.ID,
		EmployeeNum: sek.EmployeeNum,
		Sekisan: sek.Sekisan,
	}

	if err := json.NewEncoder(w).Encode(sekisan); err != nil {
		panic(err)
	}
}
