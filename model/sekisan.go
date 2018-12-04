package model

import (
	"database/sql"
)

type Sekisan struct {
	ID          int `json:"id"`
	EmployeeNum int `json:"employee_num"`
	Sekisan     int `json:"sekisan"`
}

func querySekisan(rows *sql.Rows, e error) (sekisan []*Sekisan, err error) {
	if e != nil {
		return nil, e
	}

	defer func() {
		err = rows.Close()
	}()

	sekisan = []*Sekisan{}
	for rows.Next() {
		var s Sekisan
		if err = rows.Scan(&s.ID, &s.EmployeeNum, &s.Sekisan); err != nil {
			return
		}
		sekisan = append(sekisan, &s)
	}
	err = rows.Err()
	return
}

func GetSekisanByEmployeeNum(d QueryExecutor, id string) (*Sekisan, error) {
	s, err := querySekisan(d.Query(`SELECT employee_num, member.name, sekisan FROM sekisan 
                                    INNER JOIN member ON sekisan.employee_num = employee_num 
                                    WHERE employee_num = ?`, id))
	if err != nil {
		return nil, err
	}
	return s[0], nil
}

func GetAllSekisan(d QueryExecutor) ([]*Sekisan, error) {
	s, err := querySekisan(d.Query(`SELECT * FROM sekisan`))
	if err != nil {
		return nil, err
	}
	return s, nil
}

func SetSekisan(d QueryExecutor, e, s int) error {
	_, err := d.Exec(`INSERT INTO sekisan(employee_num, sekisan) VALUES (?, ?)`, e, s)
	return err
}
