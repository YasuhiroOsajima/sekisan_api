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

//func GetSekisanByEmployeeNum(id string) (s Sekisan, err error) {
//	s = Sekisan{}
//	err = db.QueryRowx(`SELECT employee_num, member.name, sekisan FROM sekisan
//                                    INNER JOIN member ON sekisan.employee_num = employee_num
//                                    WHERE employee_num=$1`, id).StructScan(&s)
//	if err != nil {
//		return
//	}
//	return
//}
