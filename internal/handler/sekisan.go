package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"sekisan_api/internal/service"
)

type sekisanService interface {
	GetSekisanList() (sList service.SekisanList, err error)
}

type sekisanHandler struct {
	service sekisanService
}

func NewSekisanHandler(sekisanService sekisanService) *sekisanHandler{
	return &sekisanHandler{
		service: sekisanService,
	}
}

// Sekisan handler
func (h *sekisanHandler) GetAllSekisan(w http.ResponseWriter, r *http.Request) {
	sekisanRes, err := h.service.GetSekisanList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(sekisanRes); err != nil {
		panic(err)
	}
}
