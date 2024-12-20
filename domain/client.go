package domain

import "time"

type ClientType string

const (
	ClientTypeIndividual ClientType = "PF" // Pessoa Física
	ClientTypeCompany    ClientType = "PJ" // Pessoa Jurídica
)

type Client struct {
	ID        uint       `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Document  string     `gorm:"uniqueIndex;not null"` 
	Name      string     `gorm:"not null"`
	Type      ClientType `gorm:"type:varchar(2);not null"`
	IsBlocked bool       `gorm:"not null"`
}
