package controller

import (
	"log"
	"sekisan_api/model"
)

type adminList struct {
	Admin []model.Admin
}

func getAdminList() (aList adminList, err error) {
	al, err := model.GetAdminList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	aList = adminList{al}
	return
}

func registerAdmin(name, password string) (a model.Admin, err error) {
	aList, err := model.GetAdminList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	for _, a := range aList {
		if a.Name == name {
			log.Printf("[INFO] Same name admin already exists.")
			return
		}
	}

	aId64, err := model.RegisterAdmin(name, password, 1)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	aId := int(aId64)
	a, err = model.GetAdmin(aId)

	return
}

func updateAdminName(id int, name string) (a model.Admin, err error) {
	err = model.UpdateAdminName(id, name)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	a, err = model.GetAdmin(id)
	return
}

func updateAdminPassword(id int, password string) (a model.Admin, err error) {
	err = model.UpdateAdminPassword(id, password)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	a, err = model.GetAdmin(id)
	return
}

func updateAdminEnabled(id int, enabled int) (a model.Admin, err error) {
	err = model.UpdateAdminEnabled(id, enabled)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	a, err = model.GetAdmin(id)
	return
}
