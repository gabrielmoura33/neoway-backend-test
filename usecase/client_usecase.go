package usecase

import (
	"errors"

	"github.com/gabrielmoura33/neoway-backend-test/domain"
	"github.com/gabrielmoura33/neoway-backend-test/infrastructure/repository"
	"github.com/gabrielmoura33/neoway-backend-test/infrastructure/validator"
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
	if c.Type == domain.ClientTypeIndividual {
		if !validator.IsValidCPF(c.Document) {
			return errors.New("invalid CPF")
		}
	} else if c.Type == domain.ClientTypeCompany {
		if !validator.IsValidCNPJ(c.Document) {
			return errors.New("invalid CNPJ")
		}
	} else {
		return errors.New("unknown client type")
	}

	// Checar existência
	exists, err := uc.repo.Exists(c.Document)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("document already exists")
	}

	return uc.repo.Create(c)
}

func (uc *clientUseCase) GetAllClients(filterName string) ([]domain.Client, error) {
	return uc.repo.GetAll(filterName)
}

func (uc *clientUseCase) GetClientByDocument(document string) (*domain.Client, error) {
	// Verificar CPF/CNPJ válido
	if validator.IsCPF(document) {
		if !validator.IsValidCPF(document) {
			return nil, errors.New("invalid CPF")
		}
	} else if validator.IsCNPJ(document) {
		if !validator.IsValidCNPJ(document) {
			return nil, errors.New("invalid CNPJ")
		}
	} else {
		return nil, errors.New("invalid document format")
	}
	return uc.repo.GetByDocument(document)
}

func (uc *clientUseCase) ClientExists(document string) (bool, error) {
	// Apenas verifica se é válido
	if validator.IsCPF(document) {
		if !validator.IsValidCPF(document) {
			return false, errors.New("invalid CPF")
		}
	} else if validator.IsCNPJ(document) {
		if !validator.IsValidCNPJ(document) {
			return false, errors.New("invalid CNPJ")
		}
	} else {
		return false, errors.New("invalid document format")
	}
	return uc.repo.Exists(document)
}
