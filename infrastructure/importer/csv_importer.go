package importer

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/gabrielmoura33/neoway-backend-test/domain"
	"github.com/gabrielmoura33/neoway-backend-test/infrastructure/validator"
	"github.com/gabrielmoura33/neoway-backend-test/usecase"
)

func ImportCSV(filePath string, uc usecase.ClientUseCase) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	// Espera-se colunas: DOCUMENTO, NOME/RAZAO_SOCIAL
	// A primeira linha é o header
	for i, line := range records {
		if i == 0 {
			continue // pular header
		}
		if len(line) < 2 {
			continue // linha inválida
		}
		document := strings.TrimSpace(line[0])
		name := strings.TrimSpace(line[1])

		// Determinar o tipo PF ou PJ
		var clientType domain.ClientType
		cleanDoc := removeNonDigits(document)
		if len(cleanDoc) == 11 {
			clientType = domain.ClientTypeIndividual
		} else if len(cleanDoc) == 14 {
			clientType = domain.ClientTypeCompany
		} else {
			// Documento inválido, pular
			log.Printf("Documento inválido linha %d: %s", i+1, document)
			continue
		}

		client := domain.Client{
			Document:  document,
			Name:      name,
			Type:      clientType,
			IsBlocked: false, // não há essa info no CSV, assumiremos false
		}

		if err := uc.CreateClient(&client); err != nil {
			log.Printf("Erro ao importar cliente na linha %d: %v", i+1, err)
		}
	}
	return nil
}

func removeNonDigits(s string) string {
	r := strings.NewReplacer(".", "", "-", "", "/", "")
	return r.Replace(s)
}
