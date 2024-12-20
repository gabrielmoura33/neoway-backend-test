package validator

import (
	"regexp"
	"strconv"
)

func IsValidCPF(document string) bool {
	cpf := onlyDigits(document)
	if len(cpf) != 11 {
		return false
	}

	invalids := []string{
		"00000000000", "11111111111", "22222222222",
		"33333333333", "44444444444", "55555555555",
		"66666666666", "77777777777", "88888888888", "99999999999",
	}
	for _, inv := range invalids {
		if cpf == inv {
			return false
		}
	}

	for i := 9; i < 11; i++ {
		sum := 0
		for j := 0; j < i; j++ {
			num, _ := strconv.Atoi(string(cpf[j]))
			sum += num * (i + 1 - j)
		}
		rest := sum % 11
		if rest < 2 {
			rest = 0
		} else {
			rest = 11 - rest
		}
		if rest != atoi(string(cpf[i])) {
			return false
		}
	}

	return true
}

func IsValidCNPJ(document string) bool {
	cnpj := onlyDigits(document)
	if len(cnpj) != 14 {
		return false
	}

	invalids := []string{
		"00000000000000", "11111111111111", "22222222222222",
		"33333333333333", "44444444444444", "55555555555555",
		"66666666666666", "77777777777777", "88888888888888", "99999999999999",
	}
	for _, inv := range invalids {
		if cnpj == inv {
			return false
		}
	}

	var calcDigits = func(c string) bool {
		length := len(c)
		sum := 0
		pos := length - 7
		for i := length - 1; i >= 0; i-- {
			num := atoi(string(c[length-1-i]))
			sum += num * pos
			pos--
			if pos < 2 {
				pos = 9
			}
		}
		result := sum % 11
		if result < 2 {
			result = 0
		} else {
			result = 11 - result
		}
		return result == atoi(string(c[len(c)-1]))
	}

	if !calcDigits(cnpj[:12]) {
		return false
	}
	if !calcDigits(cnpj[:13]) {
		return false
	}

	return true
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func onlyDigits(s string) string {
	re := regexp.MustCompile(`\D`)
	return re.ReplaceAllString(s, "")
}
