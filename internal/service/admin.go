package service

import (
	"log"

	"sekisan_api/internal/repository"
)

type AdminList struct {
	Admin []repository.Admin
}

type adminRepositoryI interface {
	GetAdmin(id int) (a repository.Admin, err error)
	GetAdminList() (al []repository.Admin, err error)
	RegisterAdmin(name, passwd string, enabled int) (int64, error)
	UpdateAdminName(id int, name string) (err error)
	UpdateAdminPassword(id int, password string) (err error)
	UpdateAdminEnabled(id int, enabled int) (err error)
}

type adminService struct {
	aRepository adminRepositoryI
}

func NewAdminService(ar adminRepositoryI) *adminService {
	return &adminService{
		aRepository: ar,
	}
}

func (s *adminService) GetAdminList() (aList AdminList, err error) {
	al, err := s.aRepository.GetAdminList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	aList = AdminList{al}
	return
}

func (s *adminService) RegisterAdmin(name, password string) (a repository.Admin, err error) {
	aList, err := s.aRepository.GetAdminList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	for _, a := range aList {
		foundAdmin := repository.Admin{}
		if a.Name == name {
			log.Printf("[INFO] Same name admin already exists.")
			foundAdmin = a
			break
		}
		if foundAdmin.Name != "" {
			return foundAdmin, nil
		}
	}

	aId64, err := s.aRepository.RegisterAdmin(name, password, 1)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	aId := int(aId64)
	a, err = s.aRepository.GetAdmin(aId)

	return
}

func (s *adminService) UpdateAdminName(id int, name string) (a repository.Admin, err error) {
	err = s.aRepository.UpdateAdminName(id, name)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	a, err = s.aRepository.GetAdmin(id)
	return
}

func (s *adminService) UpdateAdminPassword(id int, password string) (a repository.Admin, err error) {
	err = s.aRepository.UpdateAdminPassword(id, password)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	a, err = s.aRepository.GetAdmin(id)
	return
}

func (s *adminService) UpdateAdminEnabled(id int, enabled int) (a repository.Admin, err error) {
	err = s.aRepository.UpdateAdminEnabled(id, enabled)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	a, err = s.aRepository.GetAdmin(id)
	return
}
