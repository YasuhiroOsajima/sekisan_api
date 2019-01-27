package model

type Sekisan struct {
	EmployeeNum int `json:"employee_num"`
	Hours       int `json:"hours"`
}

func GetAllSekisan() ([]Sekisan, error) {
	var s []Sekisan
	err := db.Select(&s, "SELECT * FROM `sekisan`;")

	return s, err
}

