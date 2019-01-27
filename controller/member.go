package controller

import (
	"log"
	"sekisan_api/model"
)

type MemberList struct {
	Member []model.Member
}

func getMemberList() (MemberList, error) {
	mList, err := model.GetMemberList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		badRequest(w)
		return _, err
	}

	memberList := MemberList{mList}
	return memberList, err
}

func registerMember(employeeNum int, name string) (model.Admin, error) {
	mList, err := model.GetMemberList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	for _, m := range mList {
		if m.EmployeeNum == employeeNum {
			log.Printf("[INFO] Same employee number member already exists.")
			return _, err
		}
	}

	eId64, err := model.RegisterMember(employeeNum, name, 1)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	eId := int(eId64)
	e, err := model.GetAdmin(eId)

	return e, err
}

func updateMemberName(employeeNum int, name string) (model.Member, error) {
	err := model.UpdateMemberName(employeeNum, name)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	m, err := model.GetMember(employeeNum)
	return m, err
}

func updateMemberEnabled(employeeNum, enabled int) (model.Member, error) {
	err := model.UpdateMemberEnabled(employeeNum, enabled)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	m, err := model.GetMember(employeeNum)
	return m, err
}
