package main

import (
	"log"
	"os"
	"time"

	"github.com/gabrielmoura33/neoway-backend-test/config"
	"github.com/gabrielmoura33/neoway-backend-test/infrastructure/database"
	"github.com/gabrielmoura33/neoway-backend-test/infrastructure/importer"
	"github.com/gabrielmoura33/neoway-backend-test/infrastructure/repository"
	"github.com/gabrielmoura33/neoway-backend-test/interface/handler"
	"github.com/gabrielmoura33/neoway-backend-test/interface/router"
	"github.com/gabrielmoura33/neoway-backend-test/usecase"
)

func main() {
	cfg := config.LoadConfig()
	db := database.ConnectDB(cfg)

	clientRepo := repository.NewClientRepository(db)
	clientUC := usecase.NewClientUseCase(clientRepo)
	clientHandler := handler.NewClientHandler(clientUC)
	r := router.SetupRouter(clientHandler)

	// Importar dados do CSV (se necessário)
	csvPath := os.Getenv("CSV_PATH")
	if csvPath == "" {
		csvPath = "Base_Dados_Teste.csv"
	}

	start := time.Now()
	log.Println("Iniciando importação CSV...")
	err := importer.ImportCSV(csvPath, clientUC)
	if err != nil {
		log.Printf("Erro ao importar CSV: %v", err)
	} else {
		log.Printf("Importação concluída em %s", time.Since(start))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Iniciando servidor na porta %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
