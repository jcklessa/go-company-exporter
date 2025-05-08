package repository

import (
	"database/sql"
	"fmt"
	"go-code-export/internal/model"
	"strings"
)

type CompanyRepository struct {
	DB *sql.DB
}

func (r *CompanyRepository) FetchByCodes(codes []string) ([]model.Company, error) {
	if len(codes) == 0 {
		return nil, nil
	}

	placeholders := make([]string, len(codes))
	args := make([]interface{}, len(codes))
	for i, code := range codes {
		placeholders[i] = fmt.Sprintf("$%d", i+1)
		args[i] = code
	}

	query := fmt.Sprintf("SELECT id, name, code, cnpj FROM company WHERE code IN (%s)", strings.Join(placeholders, ","))
	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []model.Company
	for rows.Next() {
		var c model.Company
		if err := rows.Scan(&c.ID, &c.Name, &c.Code, &c.CNPJ); err != nil {
			return nil, err
		}
		companies = append(companies, c)
	}
	return companies, nil
}
