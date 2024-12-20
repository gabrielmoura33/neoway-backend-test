package importer

import (
	"encoding/csv"
	"log"
	"os"
	"strings"

	"github.com/gabrielmoura33/neoway-backend-test/domain"
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

	var successCount, failCount int
	for i, line := range records {
		if i == 0 {
			// Ignora o header
			continue
		}
		if len(line) < 2 {
			log.Printf("Linha %d inválida: %v\n", i+1, line)
			failCount++
			continue
		}

		document := strings.TrimSpace(line[0])
		name := strings.TrimSpace(line[1])

		// O tipo será determinado pelo tamanho do documento após remover dígitos não numéricos.
		cleanDoc := removeNonDigits(document)
		var clientType domain.ClientType
		if len(cleanDoc) == 11 {
			clientType = domain.ClientTypeIndividual
		} else if len(cleanDoc) == 14 {
			clientType = domain.ClientTypeCompany
		} else {
			log.Printf("Linha %d: Documento inválido (%s)", i+1, document)
			failCount++
			continue
		}

		client := &domain.Client{
			Document:  document, // Armazena com pontuação
			Name:      name,
			Type:      clientType,
			IsBlocked: false,
		}

		if err := uc.CreateClient(client); err != nil {
			log.Printf("Erro ao importar cliente na linha %d (Doc: %s): %v", i+1, document, err)
			failCount++
		} else {
			successCount++
		}
	}

	log.Printf("Importação concluída. Sucesso: %d, Falhas: %d\n", successCount, failCount)
	return nil
}

func removeNonDigits(s string) string {
	r := strings.NewReplacer(".", "", "-", "", "/", "")
	return r.Replace(s)
}
