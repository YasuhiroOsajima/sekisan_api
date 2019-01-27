package model

type Member struct {
	EmployeeNum int    `json:"employee_num"`
	Name        string `json:"name"`
	Enabled     int    `json:"enabled"`
}

func GetMember(employeeNum int) (Member, error) {
	var m Member
	err := db.Get(&m,
		"SELECT `employee_num`, `name`, `enabled` FROM `mwmber` WHERE `employee_num`=?;",
		employeeNum)

	return m, err
}

func GetMemberList() ([]Member, error) {
	var m []Member
	err := db.Select(&m,
		"SELECT `employee_num`, `name`, `enabled` FROM `member`;")

	return m, err
}

func RegisterMember(employeeNum int, name string, enabled int) error {
	tx := db.MustBegin()
	tx.MustExec(
		"INSERT INTO `member`(`employee_num`, `name`, `enabled`) VALUES (?, ?, ?);",
		employeeNum, name, enabled)
	tx.MustExec(
		`INSERT INTO sekisan(employee_num, hours) VALUES (?, ?)`,
		employeeNum, 0)
	err := tx.Commit()

	return err
}

func UpdateMemberName(employeeNum int, name string) error {
	_, err := db.Exec(
		"UPDATE `member` SET `name` = ? WHERE `employee_num` = ?",
		name, employeeNum)

	return err
}

func UpdateMemberEnabled(employeeNum int, enabled int) error {
	_, err := db.Exec(
		"UPDATE `member` SET `enabled` = ? WHERE id = ?",
		enabled, employeeNum)

	return err
}
