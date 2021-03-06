package route

import (
	"github.com/gorilla/mux"

	"sekisan_api/internal/handler"
	"sekisan_api/internal/repository"
	"sekisan_api/internal/service"
)

func AddMemberRoute(r *mux.Router) {
	mr := repository.NewMemberRepository()
	ms := service.NewMemberService(mr)
	m := handler.NewMemberHandler(ms)
	r.HandleFunc("/sekisan_app/member", m.GetMemberList).Methods("GET")
	r.HandleFunc("/sekisan_app/member", m.RegisterMember).Methods("POST")
	r.HandleFunc("/sekisan_app/member", m.UpdateMemberName).Methods("POST")
	r.HandleFunc("/sekisan_app/member", m.UpdateMemberEnabled).Methods("POST")
}
