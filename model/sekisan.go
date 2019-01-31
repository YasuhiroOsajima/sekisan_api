package model

type Sekisan struct {
	EmployeeNum int `json:"employee_num"`
	Hours       int `json:"hours"`
}

func GetAllSekisan() (sl []Sekisan, err error) {
	err = db.Select(&sl, "SELECT * FROM `sekisan`;")

	return
}

