package utils

import (
	"gin/models"
	"unicode"
)

func toInt(r rune) int {
	return int(r - '0')
}

func allDigit(doc string) bool {
	for _, r := range doc {
		if !unicode.IsDigit(r) {
			return false
		}
	}

	return true
}

func ValidateDocsDriver(user *models.User, endereco *models.Endereco) (bool, string) {

	validateCPF := IsCPF(user.CPF)

	if !validateCPF {

		return false, "cpf invalid, type and try again."

	}

	validateCNH := IsCNH(user.CNH)

	if !validateCNH {

		return false, "cnh invalid, type and try again."
	}

	validateCEP := IsCEP(endereco.CEP)

	if !validateCEP {

		return false, "cep invalid, type and try again."
	}

	return true, "Ok!"

}
