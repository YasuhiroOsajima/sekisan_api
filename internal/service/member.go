package service

import (
	"log"

	"sekisan_api/internal/repository"
)

type MemberList struct {
	Member []repository.Member
}

type memberRepository interface {
	GetMember(employeeNum int) (m repository.Member, err error)
	GetMemberList() (ml []repository.Member, err error)
	RegisterMember(employeeNum int, name string, enabled int) (err error)
	UpdateMemberName(employeeNum int, name string) (err error)
	UpdateMemberEnabled(employeeNum int, enabled int) (err error)
}

type memberService struct {
	mRepository memberRepository
}

func NewMemberService(mr memberRepository) *memberService {
	return &memberService{
		mRepository: mr,
	}
}

func (s *memberService) GetMemberList() (mList MemberList, err error) {
	ml, err := s.mRepository.GetMemberList()
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	mList = MemberList{ml}
	return
}

func (s *memberService) RegisterMember(employeeNum int, name string) (m repository.Member, err error) {
	mList, err := s.mRepository.GetMemberList()
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

	err = s.mRepository.RegisterMember(employeeNum, name, 1)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	m, err = s.mRepository.GetMember(employeeNum)

	return
}

func (s *memberService) UpdateMemberName(employeeNum int, name string) (m repository.Member, err error) {
	err = s.mRepository.UpdateMemberName(employeeNum, name)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	m, err = s.mRepository.GetMember(employeeNum)
	return
}

func (s *memberService) UpdateMemberEnabled(employeeNum, enabled int) (m repository.Member, err error) {
	err = s.mRepository.UpdateMemberEnabled(employeeNum, enabled)
	if err != nil {
		log.Printf("[INFO] sql is failed.")
		return
	}

	m, err = s.mRepository.GetMember(employeeNum)
	return
}
