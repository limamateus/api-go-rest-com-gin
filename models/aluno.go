package models

import (
	"gopkg.in/validator.v2"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model
	Nome string `json:"nome" validate:"nonzero" `
	CPF  string `json:"cpf" validate:"nonzero, len=11, regexp=^[0-9]*$"`
}

func ValidacaoDoAluno(aluno *Aluno) error {
	if err := validator.Validate(aluno); err != nil { // Estou usando a bilbiote validator para validar as propriedade e retornar ume erro com base nas proriedades.
		return err
	}
	return nil

}
