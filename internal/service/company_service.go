package service

import (
	"go-code-export/internal/model"
	"go-code-export/internal/repository"
)

type CompanyService struct {
	repo *repository.CompanyRepository
}

func NewCompanyService(r *repository.CompanyRepository) *CompanyService {
	return &CompanyService{repo: r}
}

func (s *CompanyService) ProcessCodes(allCodes []string) ([]model.Company, []string, error) {
	var foundCompanies []model.Company
	foundCodes := map[string]bool{}

	const batchSize = 1000
	for i := 0; i < len(allCodes); i += batchSize {
		end := i + batchSize
		if end > len(allCodes) {
			end = len(allCodes)
		}
		batch := allCodes[i:end]

		results, err := s.repo.FetchByCodes(batch)
		if err != nil {
			return nil, nil, err
		}
		foundCompanies = append(foundCompanies, results...)
		for _, company := range results {
			foundCodes[company.Code] = true
		}
	}

	var notFound []string
	for _, code := range allCodes {
		if !foundCodes[code] {
			notFound = append(notFound, code)
		}
	}

	return foundCompanies, notFound, nil
}
