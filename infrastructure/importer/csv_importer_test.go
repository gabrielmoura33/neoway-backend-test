package importer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/gabrielmoura33/neoway-backend-test/domain"
	"github.com/gabrielmoura33/neoway-backend-test/infrastructure/importer"
	"github.com/gabrielmoura33/neoway-backend-test/usecase"
)

// Mock do ClientUseCase
type mockClientUseCase struct {
	clients []domain.Client
}

func (m *mockClientUseCase) CreateClient(c *domain.Client) error {
	m.clients = append(m.clients, *c)
	return nil
}
func (m *mockClientUseCase) GetAllClients(filterName string) ([]domain.Client, error) {
	return m.clients, nil
}
func (m *mockClientUseCase) GetClientByDocument(document string) (*domain.Client, error) {
	for _, cl := range m.clients {
		if cl.Document == document {
			return &cl, nil
		}
	}
	return nil, nil
}
func (m *mockClientUseCase) ClientExists(document string) (bool, error) {
	for _, cl := range m.clients {
		if cl.Document == document {
			return true, nil
		}
	}
	return false, nil
}

func TestImportCSV(t *testing.T) {
	uc := &mockClientUseCase{}
	csvPath := filepath.Join("testdata", "valid.csv")
	err := importer.ImportCSV(csvPath, uc)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(uc.clients) != 1 {
		t.Fatalf("Expected 1 client, got %d", len(uc.clients))
	}

	// Verifica se o cliente foi importado corretamente
	if uc.clients[0].Name != "Alphadale Curtis" {
		t.Errorf("Expected name 'Alphadale Curtis', got '%s'", uc.clients[0].Name)
	}
}
