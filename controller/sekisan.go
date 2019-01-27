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

func getSekisanList() (sekisanList, error) {
	hourList, err := model.GetAllSekisan()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	memberList, err := model.GetMemberList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	var sList []sekisan
	for _, h := range hourList {
		for _, m := range memberList {
			if h.EmployeeNum == m.EmployeeNum {
				sek := sekisan{
					EmployeeNum: h.EmployeeNum,
					Name:        m.Name,
					Hours:       h.Hours,
				}

				sList = append(sList, sek)
			}
		}
	}
	sekisanRes := sekisanList{sList}
	return sekisanRes, err
}
