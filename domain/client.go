package domain

import "time"

type ClientType string

const (
	ClientTypeIndividual ClientType = "PF" // Pessoa Física
	ClientTypeCompany    ClientType = "PJ" // Pessoa Jurídica
)

type Client struct {
	ID        uint       `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Document  string     `gorm:"uniqueIndex;not null"` // CPF ou CNPJ
	Name      string     `gorm:"not null"`
	Type      ClientType `gorm:"type:varchar(2);not null"` // PF ou PJ
	IsBlocked bool       `gorm:"not null"`
}
