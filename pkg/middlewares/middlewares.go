package middlewares

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"

	"golang_api/pkg/utils"
	"golang_api/pkg/models"
)

/*
AuthenticationMiddleware é um middleware que valida o token de autenticação fornecido no cabeçalho da requisição.

Este middleware verifica a presença e a validade de um token JWT no cabeçalho `Authorization`. Se o token estiver ausente ou inválido, ele retorna uma resposta de erro com o status HTTP 401 (não autorizado) e encerra a requisição. Caso o token seja válido, o middleware extrai o `user_id` do token e o define no contexto da requisição, permitindo o acesso a esse valor nos próximos handlers.

Fluxo de funcionamento:
1. Verifica se o cabeçalho `Authorization` está presente na requisição.
2. Se o token estiver ausente ou estiver mal formatado, retorna erro 401.
3. Verifica se o token é válido utilizando a função `VerifyToken`.
4. Se o token for válido, extrai o `user_id` dos claims e o armazena no contexto da requisição.
5. Chama o próximo middleware ou handler.

Retorno: gin.HandlerFunc, que é um handler para o middleware.

Exemplo de uso:
    r := gin.Default()
    r.Use(AuthenticationMiddleware()) // Aplica o middleware de autenticação
*/
func AuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "missing authentication token"})
			c.Abort()
			return
		}

		tokenParts := strings.Split(tokenString, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "invalid authentication token"})
			c.Abort()
			return
		}

		tokenString = tokenParts[1]

		claims, err := utils.VerifyToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, models.ErrorResponse{Error: "invalid authentication token"})
			c.Abort()
			return
		}

		c.Set("user_id", claims["user_id"])
		c.Next()
	}
}