package repository

type Admin struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Enabled int    `json:"enabled"`
}

type adminRepository struct{}

func NewAdminRepository() *adminRepository {
	return &adminRepository{}
}

func (r *adminRepository) GetAdmin(id int) (a Admin, err error) {
	err = db.Get(&a,
		"SELECT `id`, `name`, `enabled` FROM `admin` WHERE `id`=?;", id)

	return
}

func (r *adminRepository) GetAdminList() (al []Admin, err error) {
	err = db.Select(&al, "SELECT `id`, `name`, `enabled` FROM `admin`;")

	return
}

func (r *adminRepository) RegisterAdmin(name, passwd string, enabled int) (int64, error) {
	res, err := db.Exec(
		"INSERT INTO `admin`(`name`, `password`, `enabled`) VALUES (?, ?, ?);",
		name, passwd, enabled)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r *adminRepository) UpdateAdminName(id int, name string) (err error) {
	_, err = db.Exec(
		"UPDATE `admin` SET `name` = ? WHERE `id` = ?", name, id)

	return
}

func (r *adminRepository) UpdateAdminPassword(id int, password string) (err error) {
	_, err = db.Exec(
		"UPDATE `admin` SET `password` = ? WHERE `id` = ?", password, id)

	return
}

func (r *adminRepository) UpdateAdminEnabled(id int, enabled int) (err error) {
	_, err = db.Exec(
		"UPDATE `admin` SET `password` = ? WHERE `id` = ?", enabled, id)

	return
}
