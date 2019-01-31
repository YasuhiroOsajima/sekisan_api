package controller

import (
	"log"
	"sekisan_api/model"
)

type memberList struct {
	Member []model.Member
}

func getMemberList() (mList memberList, err error) {
	ml, err := model.GetMemberList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	mList = memberList{ml}
	return
}

func registerMember(employeeNum int, name string) (m model.Member, err error) {
	mList, err := model.GetMemberList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	for _, m := range mList {
		if m.EmployeeNum == employeeNum {
			log.Printf("[INFO] Same employee number member already exists.")
			return
		}
	}

	err = model.RegisterMember(employeeNum, name, 1)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	m, err = model.GetMember(employeeNum)

	return
}

func updateMemberName(employeeNum int, name string) (m model.Member, err error) {
	err = model.UpdateMemberName(employeeNum, name)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	m, err = model.GetMember(employeeNum)
	return
}

func updateMemberEnabled(employeeNum, enabled int) (m model.Member, err error) {
	err = model.UpdateMemberEnabled(employeeNum, enabled)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	m, err = model.GetMember(employeeNum)
	return
}
