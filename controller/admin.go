package controller

import (
	"log"
	"sekisan_api/model"
)

type AdminList struct {
	Admin []model.Admin
}

func getAdminList() (AdminList, error) {
	aList, err := model.GetAdminList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	adminList := AdminList{aList}
	return adminList, err
}

func registerAdmin(name, password string) (model.Admin, error) {
	aList, err := model.GetAdminList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	for _, a := range aList {
		if a.Name == name {
			log.Printf("[INFO] Same name admin already exists.")
			return _, err
		}
	}

	aId64, err := model.RegisterAdmin(name, password, 1)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	aId := int(aId64)
	a, err := model.GetAdmin(aId)

	return a, err
}

func updateAdminName(id int, name string) (model.Admin, error) {
	err := model.UpdateAdminName(id, name)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	a, err := model.GetAdmin(id)
	return a, err
}

func updateAdminPassword(id int, password string) (model.Admin, error) {
	err := model.UpdateAdminPassword(id, password)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	a, err := model.GetAdmin(id)
	return a, err
}

func updateAdminEnabled(id int, enabled int) (model.Admin, error) {
	err := model.UpdateAdminEnabled(id, enabled)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return _, err
	}

	a, err := model.GetAdmin(id)
	return a, err
}
