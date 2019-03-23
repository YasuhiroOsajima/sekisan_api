package repository

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

type transactionRepository struct{}

func NewTransactionRepository() *transactionRepository {
	return &transactionRepository{}
}

func (r *transactionRepository) GetTransaction() (tl []Transaction, err error) {
	err = db.Get(&tl,
		"SELECT `id`, `employee_num`, `updated_date`, "+
			"`before`, `add`, `subtracted`, `after`, `reason` "+
			"FROM `transactions`")

	return
}

func (r *transactionRepository) AddTransaction(employeeNum, hour int, operation, reason string) (int64, error) {
	res, err := db.Exec(
		"INSERT INTO `transactions`(`employee_num`, `updated_date`, `before`, `added`, `subtracted`, `after`, `reason`) "+
			"VALUES (?, ?, ?, ?, ?, ?, ?);",
		employeeNum, hour, operation, reason)

	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}
