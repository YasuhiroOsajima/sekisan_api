package model

type Admin struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Enabled int    `json:"enabled"`
}

func GetAdmin(id int) (Admin, error) {
	var a Admin
	err := db.Get(&a,
		"SELECT `id`, `name`, `enabled` FROM `admin` WHERE `id`=?;", id)

	return a, err
}

func GetAdminList() ([]Admin, error) {
	var a []Admin
	err := db.Select(&a, "SELECT `id`, `name`, `enabled` FROM `admin`;")

	return a, err
}

func RegisterAdmin(name, passwd string, enabled int) (int64, error) {
	res, err := db.Exec(
		"INSERT INTO `admin`(`name`, `password`, `enabled`) VALUES (?, ?, ?);",
		name, passwd, enabled)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func UpdateAdminName(id int, name string) error {
	_, err := db.Exec(
		"UPDATE `admin` SET `name` = ? WHERE `id` = ?", name, id)

	return err
}

func UpdateAdminPassword(id int, password string) error {
	_, err := db.Exec(
		"UPDATE `admin` SET `password` = ? WHERE `id` = ?", password, id)

	return err
}

func UpdateAdminEnabled(id int, enabled int) error {
	_, err := db.Exec(
		"UPDATE `admin` SET `password` = ? WHERE `id` = ?", enabled, id)

	return err
}
