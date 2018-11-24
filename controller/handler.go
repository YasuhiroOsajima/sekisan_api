package controller

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"net/http"
	"time"
)

const (
	SessionName = "test_session"
)

var BaseTime time.Time

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

//func (h *Handler) userByRequest(r *http.Request) (*model.User, error) {
//	v := r.Context().Value("user_id")
//	if id, ok := v.(int64); ok {
//		return model.GetUserByID(h.db, id)
//	}
//	return nil, errors.New("Not authenticated")
//}

// SELECT * FROM sekisan WHERE `employee_num`=2001;

func (h *Handler) ShowId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "My number is : %v\n", vars["id"])
}
