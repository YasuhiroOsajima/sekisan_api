package repository

type Sekisan struct {
	EmployeeNum int `json:"employee_num"`
	Hours       int `json:"hours"`
}

type sekisanRepository struct{}

func NewSekisanRepository() *sekisanRepository {
	return &sekisanRepository{}
}

func (r *sekisanRepository) GetAllSekisan() (sl []Sekisan, err error) {
	err = db.Select(&sl, "SELECT * FROM `sekisan`;")

	return
}
