package model

type Transaction struct {
	Id          int    `json:"id"`
	EmployeeNum int    `json:"employee_num"`
	UpdatedDate string `json:"updated_date"`
	Before      int    `json:"before"`
	Added       int    `json:"added"`
	Subtracted  int    `json:"subtracted"`
	After       int    `json:"after"`
	Reason      string `json:"reason"`
}

func GetTransaction(employeeNum int) (Member, error) {
	var m Member
	err := db.Get(&m,
		"SELECT `employee_num`, `name`, `enabled` FROM `mwmber` WHERE `employee_num`=?;",
		employeeNum)

	return m, err
}
