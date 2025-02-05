package auth

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CriarToken(id int) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioID"] = id
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)

	return token.SignedString([]byte(config.SecretKey))
}

func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}
	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil

	}
	return fmt.Errorf("Token inválido")
}

func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura inesperado! %v", token.Header["alg"])
	}
	return []byte(config.SecretKey), nil
}

func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return " "

}

func ExtrairUsuarioId(r *http.Request) (int, error) {
	tokenString := extrairToken(r)
	token, err := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if err != nil {
		return 0, err
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioId, err := strconv.ParseInt(fmt.Sprintf("%.0f", permissoes["usuarioID"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return int(usuarioId), nil
	}

	return 0, errors.New("token inválido")
}
