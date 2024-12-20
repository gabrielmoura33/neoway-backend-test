package validator_test

import (
	"testing"

	"github.com/gabrielmoura33/neoway-backend-test/infrastructure/validator"
)

func TestIsValidCPF(t *testing.T) {
	validCPF := "12345678909" // CPF válido de exemplo
	if !validator.IsValidCPF(validCPF) {
		t.Errorf("Expected CPF %s to be valid", validCPF)
	}

	invalidCPF := "11111111111"
	if validator.IsValidCPF(invalidCPF) {
		t.Errorf("Expected CPF %s to be invalid", invalidCPF)
	}
}

func TestIsValidCNPJ(t *testing.T) {
	// CNPJ válido de exemplo (ex: 30607887000141)
	validCNPJ := "30.607.887/0001-41"
	if !validator.IsValidCNPJ(validCNPJ) {
		t.Errorf("Expected CNPJ %s to be valid", validCNPJ)
	}

	invalidCNPJ := "00.000.000/0000-00"
	if validator.IsValidCNPJ(invalidCNPJ) {
		t.Errorf("Expected CNPJ %s to be invalid", invalidCNPJ)
	}
}
