package controller

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"log"
	"net/http"
	"sekisan_api/model"
)

const (
	SessionName = "test_session"
)

//var BaseTime time.Time

type Handler struct {
	db    *sql.DB
	store sessions.Store
}

func NewHandler(db *sql.DB, store sessions.Store) *Handler {
	//BaseTime = time.Date(2018, 10, 16, 10, 0, 0, 0, time.Local)
	return &Handler{
		db:    db,
		store: store,
	}
}

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
