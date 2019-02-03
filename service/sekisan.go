package service

import (
	"log"
	"sekisan_api/repository"
)

type sekisan struct {
	EmployeeNum int    `json:"employee_num"`
	Name        string `json:"name"`
	Hours       int    `json:"hours"`
}

type SekisanList struct {
	Sekisan []sekisan
}

type sekisanRepository interface {
	GetAllSekisan() (sl []repository.Sekisan, err error)
}

type sekisanService struct {
	sRepository sekisanRepository
	mRepository memberRepository
}

func NewSekisanService(sr sekisanRepository, mr memberRepository) *sekisanService {
	return &sekisanService{
		sRepository: sr,
		mRepository: mr,
	}
}

func (s *sekisanService) GetSekisanList() (sList SekisanList, err error) {
	hList, err := s.sRepository.GetAllSekisan()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	mList, err := s.mRepository.GetMemberList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	var sl []sekisan
	for _, h := range hList {
		for _, m := range mList {
			if h.EmployeeNum == m.EmployeeNum {
				sek := sekisan{
					EmployeeNum: h.EmployeeNum,
					Name:        m.Name,
					Hours:       h.Hours,
				}

				sl = append(sl, sek)
			}
		}
	}
	sList = SekisanList{sl}
	return
}
