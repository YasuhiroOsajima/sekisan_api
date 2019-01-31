package model

type Admin struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Enabled int    `json:"enabled"`
}

func GetAdmin(id int) (a Admin, err error) {
	err = db.Get(&a,
		"SELECT `id`, `name`, `enabled` FROM `admin` WHERE `id`=?;", id)

	return
}

func GetAdminList() (al []Admin, err error) {
	err = db.Select(&al, "SELECT `id`, `name`, `enabled` FROM `admin`;")

	return
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

func UpdateAdminName(id int, name string) (err error) {
	_, err = db.Exec(
		"UPDATE `admin` SET `name` = ? WHERE `id` = ?", name, id)

	return
}

func UpdateAdminPassword(id int, password string) (err error) {
	_, err = db.Exec(
		"UPDATE `admin` SET `password` = ? WHERE `id` = ?", password, id)

	return
}

func UpdateAdminEnabled(id int, enabled int) (err error) {
	_, err = db.Exec(
		"UPDATE `admin` SET `password` = ? WHERE `id` = ?", enabled, id)

	return
}
