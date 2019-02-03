package route

import (
	"github.com/gorilla/mux"

	"sekisan_api/handler"
	"sekisan_api/repository"
	"sekisan_api/service"
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
