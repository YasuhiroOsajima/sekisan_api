package route

import (
	"github.com/gorilla/mux"
	"sekisan_api/handler"
	"sekisan_api/repository"
	"sekisan_api/service"
)

func AddSekisanRoute(r *mux.Router) {
	sr := repository.NewSekisanRepository()
	mr := repository.NewMemberRepository()
	ss := service.NewSekisanService(sr, mr)
	s := handler.NewSekisanHandler(ss)
	r.HandleFunc("/sekisan", s.GetAllSekisan).Methods("GET")
}
