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

func RegisterMember(employeeNum int, name string, enabled int) (int64, error) {
	res, err := db.Exec(
		"INSERT INTO `member`(`employee_num`, `name`, `enabled`) VALUES (?, ?, ?);",
		employeeNum, name, enabled)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
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
