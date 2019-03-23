package route

import (
	"sekisan_api/modules/pkg/mod/github.com/gorilla/mux@v1.6.2"
	"sekisan_api/internal/handler"
	"sekisan_api/internal/repository"
	"sekisan_api/internal/service"
)

func AddSekisanRoute(r *mux.Router) {
	sr := repository.NewSekisanRepository()
	mr := repository.NewMemberRepository()
	ss := service.NewSekisanService(sr, mr)
	s := handler.NewSekisanHandler(ss)
	r.HandleFunc("/sekisan", s.GetAllSekisan).Methods("GET")
}
