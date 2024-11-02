# Golang Project

Este é um projeto de API construído em Go utilizando o framework Gin e banco de dados PostgreSQL.

## Pré-requisitos

Antes de começar, verifique se você tem as seguintes ferramentas instaladas:

- [Go](https://golang.org/dl/) (versão recomendada: 1.XX)

## Instalando Dependências

Navegue até o diretório do projeto e instale as dependências:

```prompt
go mod tidy
```


## Executando o Projeto

Após configurar tudo, execute o projeto com o seguinte comando:

```prompt
go run ./cmd/api/main.go
```

## Swagger

Após a compilação do projeto, é possível realizar o acesso do Swagger para visualizar os endpoints

```
http://localhost:8080/swagger/index.html#/
```

Ao implementar novos endpoints, deve ser realizado a documentação dos métodos dentro do código, assim como o modelo abaixo:

```
// NomeDoMetodo godoc
// @Summary Resumo do método
// @Description Descrição do Método
// @Tags nome da tag, exemplo (usuários)
// @Accept tipo de entrada aceito (ex: json)
// @Produce tipo de retorno aceito (ex: json)
// @Success 200 {tipo} ({array}, {bool}, {object}) string
// @Failure 500 {tipo} ({array}, {bool}, {object}) string
// @Router /nomeDaRota [tipo HTTP] ([get], [post], [put], [delete], [patch], [options], [trace], [connect])
func NomeDoMetodo(c *gin.Context){
    c.JSON(200, true)
}
```

Se você realizar alterações nos endpoints, para atualizar a documentação do swagger, rode o seguinte comando:

```
swag init -g /cmd/api/main.go
```