package usecase

import (
	"errors"
	"strings"

	"github.com/gabrielmoura33/neoway-backend-test/domain"
	"github.com/gabrielmoura33/neoway-backend-test/infrastructure/repository"
)

type ClientUseCase interface {
	CreateClient(c *domain.Client) error
	GetAllClients(filterName string) ([]domain.Client, error)
	GetClientByDocument(document string) (*domain.Client, error)
	ClientExists(document string) (bool, error)
}

type clientUseCase struct {
	repo repository.ClientRepository
}

func NewClientUseCase(repo repository.ClientRepository) ClientUseCase {
	return &clientUseCase{repo: repo}
}

func (uc *clientUseCase) CreateClient(c *domain.Client) error {
	cleanDoc := removeNonDigits(c.Document)

	// Validação simples: apenas checar tamanho
	if c.Type == domain.ClientTypeIndividual {
		// Espera-se 11 dígitos para CPF
		if len(cleanDoc) != 11 {
			return errors.New("invalid CPF length")
		}
	} else if c.Type == domain.ClientTypeCompany {
		// Espera-se 14 dígitos para CNPJ
		if len(cleanDoc) != 14 {
			return errors.New("invalid CNPJ length")
		}
	} else {
		return errors.New("unknown client type")
	}

	// Verifica se o documento já existe (utilizando a versão sem pontuações)
	exists, err := uc.repo.Exists(cleanDoc)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("document already exists")
	}

	return uc.repo.Create(&domain.Client{
		Document:  c.Document, // Armazena com pontuações
		Name:      c.Name,
		Type:      c.Type,
		IsBlocked: c.IsBlocked,
	})
}

func (uc *clientUseCase) GetAllClients(filterName string) ([]domain.Client, error) {
	return uc.repo.GetAll(filterName)
}

func (uc *clientUseCase) GetClientByDocument(document string) (*domain.Client, error) {
	cleanDoc := removeNonDigits(document)

	if len(cleanDoc) == 11 {
		// CPF - mas não verificamos dígito verificador
	} else if len(cleanDoc) == 14 {
		// CNPJ - mas não verificamos dígito verificador
	} else {
		return nil, errors.New("invalid document format")
	}
	return uc.repo.GetByDocument(document)
}

func (uc *clientUseCase) ClientExists(document string) (bool, error) {
	cleanDoc := removeNonDigits(document)

	if len(cleanDoc) == 11 {
		// CPF de 11 dígitos
	} else if len(cleanDoc) == 14 {
		// CNPJ de 14 dígitos
	} else {
		return false, errors.New("invalid document format")
	}
	return uc.repo.Exists(cleanDoc)
}

func removeNonDigits(s string) string {
	r := strings.NewReplacer(".", "", "-", "", "/", "")
	return r.Replace(s)
}
