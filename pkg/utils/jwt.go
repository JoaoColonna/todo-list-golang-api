package utils

import (
	"fmt"
	"time"
	"github.com/golang-jwt/jwt/v4"
)

var secretKey = []byte("secretpassword")

/*
GenerateToken gera um token JWT (JSON Web Token) válido para o usuário, com base no ID do usuário fornecido.

Esta função cria um token JWT utilizando a biblioteca `jwt` e inclui no token as seguintes informações:
- `user_id`: O ID do usuário fornecido como argumento.
- `exp`: A data de expiração do token, que é definida para 1 hora a partir do momento atual.

O token gerado é assinado com uma chave secreta (definida pela variável `secretKey`) e usa o algoritmo de assinatura HS256.

Retorno:
- Um token JWT assinado em formato de string, ou um erro caso a assinatura falhe.

Exemplo de uso:
    token, err := GenerateToken(userID)
    if err != nil {
        // Tratar erro
    }
    fmt.Println("Token gerado:", token)
*/
func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

/*
VerifyToken verifica a validade de um token JWT, decodificando e validando sua assinatura.

Esta função recebe um token JWT como string, tenta parseá-lo e verifica se a assinatura é válida usando o método de assinatura HS256 com a chave secreta (`secretKey`). Se o token for válido, ela retorna os claims (dados) presentes no token. Caso contrário, ela retorna um erro.

Fluxo:
1. A função tenta parsear o token JWT com a chave secreta e verifica se o método de assinatura é o esperado (HS256).
2. Se o token for válido, os claims (informações do token) são extraídos e retornados.
3. Se o token não for válido ou a assinatura não for reconhecida, a função retorna um erro.

Retorno:
- `jwt.MapClaims`: Os claims (dados) presentes no token JWT, caso o token seja válido.
- `error`: Um erro caso o token seja inválido, esteja mal formatado ou a assinatura não seja válida.

Exemplo de uso:
    claims, err := VerifyToken(tokenString)
    if err != nil {
        // Tratar erro, token inválido
    }
    fmt.Println("Claims do token:", claims)
*/
func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if token.Method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("invalid signing method")
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}