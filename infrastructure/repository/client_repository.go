package repository

import (
	"strings"

	"github.com/gabrielmoura33/neoway-backend-test/domain"
	"gorm.io/gorm"
)

type ClientRepository interface {
	Create(client *domain.Client) error
	GetAll(filterName string) ([]domain.Client, error)
	GetByDocument(doc string) (*domain.Client, error)
	Exists(doc string) (bool, error)
}

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) ClientRepository {
	return &clientRepository{db: db}
}

func (r *clientRepository) Create(client *domain.Client) error {
	return r.db.Create(client).Error
}

func (r *clientRepository) GetAll(filterName string) ([]domain.Client, error) {
	var clients []domain.Client
	query := r.db.Model(&domain.Client{}).Order("name ASC")
	if filterName != "" {
		query = query.Where("LOWER(name) LIKE ?", "%"+strings.ToLower(filterName)+"%")
	}
	if err := query.Find(&clients).Error; err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *clientRepository) GetByDocument(doc string) (*domain.Client, error) {
	var client domain.Client
	if err := r.db.Where("document = ?", doc).First(&client).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) Exists(doc string) (bool, error) {
	var count int64
	if err := r.db.Model(&domain.Client{}).Where("document = ?", doc).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
