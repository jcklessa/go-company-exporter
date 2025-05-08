package main

import (
	"go-code-export/config"
	"go-code-export/internal/repository"
	"go-code-export/internal/service"
	"go-code-export/internal/util"
	"log"
)

func main() {
	codes, err := util.ReadCodesFromFile("codes_all.txt")
	if err != nil {
		log.Fatalf("Erro ao ler o arquivo: %v", err)
	}

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}
	defer db.Close()

	repo := &repository.CompanyRepository{DB: db}
	service := service.NewCompanyService(repo)

	detected, notFound, err := service.ProcessCodes(codes)
	if err != nil {
		log.Fatalf("Erro ao processar códigos: %v", err)
	}

	err = util.GenerateExcel("companies.xlsx", detected, notFound)
	if err != nil {
		log.Fatalf("Erro ao gerar excel: %v", err)
	}

	log.Println("Arquivo gerado com sucesso: companies.xlsx")
}
