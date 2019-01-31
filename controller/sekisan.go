package controller

import (
	"log"
	"sekisan_api/model"
)

type sekisan struct {
	EmployeeNum int    `json:"employee_num"`
	Name        string `json:"name"`
	Hours       int    `json:"hours"`
}

type sekisanList struct {
	Sekisan []sekisan
}

func getSekisanList() (sList sekisanList, err error) {
	hList, err := model.GetAllSekisan()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	mList, err := model.GetMemberList()
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
	sList = sekisanList{sl}
	return
}
