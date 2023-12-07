package services

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtService struct { // Objeto do Jwt que ira contar a chave e o issure
	secretKey string
	issure    string
}

func NewJWTService() *jwtService { // Aqui estou criando uma nova assinatura
	return &jwtService{
		secretKey: "secret-key",
		issure:    "aluno-api",
	}
}

type Claim struct { // Aqui eu crio um estrutura para criar um json de claim
	Sum  uint   `json:"sum"`
	Nome string `json:"f1"`
	jwt.StandardClaims
}

func (s *jwtService) GeracaoDeToken(id uint, nome string) (string, error) { // Metodo para gerar o token
	claim := &Claim{ // Aqui estou criando os dados da claim
		id,   // id do usuario
		nome, // Nome do usaurio
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 2).Unix(), // expiraçãdo token
			Issuer:    s.issure,                             // issure
			IssuedAt:  time.Now().Unix(),                    // data de criação do token
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // crio o token com converto ele para HS256

	t, err := token.SignedString([]byte(s.secretKey)) // realizo a assinatura.
	if err != nil {
		return "", err
	}

	return t, nil
}

func (s *jwtService) ValidacaoDoToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, isValid := t.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Token invalid %v", token)
		}
		return []byte(s.secretKey), nil
	})

	return err == nil
}
