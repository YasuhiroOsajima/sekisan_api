package model

type Member struct {
	EmployeeNum int    `json:"employee_num"`
	Name        string `json:"name"`
	Enabled     int    `json:"enabled"`
}

func GetMember(employeeNum int) (m Member, err error) {
	err = db.Get(&m,
		"SELECT `employee_num`, `name`, `enabled` FROM `mwmber` WHERE `employee_num`=?;",
		employeeNum)

	return
}

func GetMemberList() (ml []Member, err error) {
	err = db.Select(&ml,
		"SELECT `employee_num`, `name`, `enabled` FROM `member`;")

	return
}

func RegisterMember(employeeNum int, name string, enabled int) (err error) {
	tx := db.MustBegin()
	tx.MustExec(
		"INSERT INTO `member`(`employee_num`, `name`, `enabled`) VALUES (?, ?, ?);",
		employeeNum, name, enabled)
	tx.MustExec(
		`INSERT INTO sekisan(employee_num, hours) VALUES (?, ?)`,
		employeeNum, 0)
	err = tx.Commit()

	return
}

func UpdateMemberName(employeeNum int, name string) (err error) {
	_, err = db.Exec(
		"UPDATE `member` SET `name` = ? WHERE `employee_num` = ?",
		name, employeeNum)

	return
}

func UpdateMemberEnabled(employeeNum int, enabled int) (err error) {
	_, err = db.Exec(
		"UPDATE `member` SET `enabled` = ? WHERE id = ?",
		enabled, employeeNum)

	return
}
