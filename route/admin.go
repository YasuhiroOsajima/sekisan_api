package route

import (
	"github.com/gorilla/mux"

	"sekisan_api/handler"
	"sekisan_api/repository"
	"sekisan_api/service"
)

func AddAdminRoute(r *mux.Router) {
	ar := repository.NewAdminRepository()
	as := service.NewAdminService(ar)
	a := handler.NewAdminHandler(as)

	r.HandleFunc("/sekisan_app/admin", a.GetAdminList).Methods("GET")
	r.HandleFunc("/sekisan_app/admin", a.RegisterAdmin).Methods("POST")
	r.HandleFunc("/sekisan_app/{admin_id[0-9]+}/name", a.UpdateAdminName).Methods("POST")
	r.HandleFunc("/sekisan_app/{admin_id[0-9]+}/name", a.UpdateAdminPassword).Methods("POST")
	r.HandleFunc("/sekisan_app/{admin_id[0-9]+}/name", a.UpdateAdminEnabled).Methods("POST")
}
