package route

import (
	"github.com/gorilla/mux@v1.6.2"

	"sekisan_api/internal/handler"
	"sekisan_api/internal/repository"
	"sekisan_api/internal/service"
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
