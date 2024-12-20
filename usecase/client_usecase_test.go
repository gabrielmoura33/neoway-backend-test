package usecase_test

import (
	"errors"
	"testing"

	"github.com/gabrielmoura33/neoway-backend-test/domain"
	"github.com/gabrielmoura33/neoway-backend-test/usecase"
)

// Mock Repository
type mockRepo struct {
	clients []domain.Client
}

func (m *mockRepo) Create(c *domain.Client) error {
	m.clients = append(m.clients, *c)
	return nil
}
func (m *mockRepo) GetAll(filterName string) ([]domain.Client, error) {
	return m.clients, nil
}
func (m *mockRepo) GetByDocument(doc string) (*domain.Client, error) {
	for _, cl := range m.clients {
		if cl.Document == doc {
			return &cl, nil
		}
	}
	return nil, nil
}
func (m *mockRepo) Exists(doc string) (bool, error) {
	for _, cl := range m.clients {
		if cl.Document == doc {
			return true, nil
		}
	}
	return false, nil
}

func TestCreateClient(t *testing.T) {
	repo := &mockRepo{}
	uc := usecase.NewClientUseCase(repo)

	client := &domain.Client{
		Document:  "30.607.887/0001-41",
		Name:      "Test Company",
		Type:      domain.ClientTypeCompany,
		IsBlocked: false,
	}

	err := uc.CreateClient(client)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verifica se o cliente foi inserido
	if len(repo.clients) != 1 {
		t.Errorf("Expected 1 client, got %d", len(repo.clients))
	}
}

func TestCreateClientExists(t *testing.T) {
	repo := &mockRepo{
		clients: []domain.Client{
			{Document: "30.607.887/0001-41", Name: "Existing Company", Type: domain.ClientTypeCompany},
		},
	}
	uc := usecase.NewClientUseCase(repo)

	client := &domain.Client{
		Document: "30.607.887/0001-41",
		Name:     "Another Company",
		Type:     domain.ClientTypeCompany,
	}

	err := uc.CreateClient(client)
	if err == nil {
		t.Fatalf("Expected error due to existing document, got nil")
	}

	if !errors.Is(err, errors.New("document already exists")) {
		t.Errorf("Expected 'document already exists' error, got %v", err)
	}
}
