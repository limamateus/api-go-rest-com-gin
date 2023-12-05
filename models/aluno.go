package models

import (
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type Aluno struct {
	gorm.Model `json:"-"`
	Nome       string `json:"nome" validate:"required"`
	CPF        string `json:"cpf" validate:"required,len=11"`
}

func ValidacaoDoAluno(aluno *Aluno) error {
	validate := validator.New()
	if err := validate.Struct(aluno); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err
		}

		var fieldErrors []string
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "required":
				fieldErrors = append(fieldErrors, "O campo '"+err.Field()+"' é obrigatório")
			case "len":
				fieldErrors = append(fieldErrors, "O campo '"+err.Field()+"' deve ter 11 dígitos")
			default:
				fieldErrors = append(fieldErrors, "Erro de validação no campo '"+err.Field()+"'")
			}
		}

		return errors.New("Erros de validação: " + strings.Join(fieldErrors, ", "))
	}
	return nil
}
