package domain

import (
	"database/sql"
	"errors"
	"log"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const TOKEN_DURATION = time.Hour

type Login struct {
	Username   string         `db:"username"`
	CustomerId sql.NullString `db:"customer_id"`
	Accounts   sql.NullString `db:"account_numbers"`
	Role       string         `db:"role"`
}

//GenerateToken gera o token a partir dos dados do usuário
func (l Login) GenerateToken() (*string, error) {

	var claims jwt.MapClaims

	//if o atributo account e customer não foram nulos criar role de usuário, caso seja, role de admin
	if l.Accounts.Valid && l.CustomerId.Valid {
		claims = l.claimsForUser()
	} else {
		claims = l.claimsForAdmin()
	}

	// 4º É invocado o método jwt.NewWithClaims pra criação das claims com dados do usuário
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 5º É retornado a numeração do token no formato String por base nas clams geradas
	signedTokenAsString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))

	if err != nil {
		log.Println("Failed while signing token: " + err.Error())
		return nil, errors.New("cannot generate token")
	}
	return &signedTokenAsString, nil
}

// gera token pra usário com permissão do tipo user
func (l Login) claimsForUser() jwt.MapClaims {
	accounts := strings.Split(l.Accounts.String, ",")
	return jwt.MapClaims{
		"customer_id": l.CustomerId.String,
		"role":        l.Role,
		"username":    l.Username,
		"accounts":    accounts,
		"exp":         time.Now().Add(TOKEN_DURATION).Unix(),
	}
}

// gera token pra usário com permissão do tipo admin
func (l Login) claimsForAdmin() jwt.MapClaims {
	return jwt.MapClaims{
		"role":     l.Role,
		"username": l.Username,
		"exp":      time.Now().Add(TOKEN_DURATION).Unix(),
	}
}
