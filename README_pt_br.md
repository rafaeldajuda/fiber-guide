[english version](./README.md)

# MANUAL FIBER

Este manual contém o básico necessário para desenvolver uma API com o fiber.

Fiber para quem não sabe é um framework em GO (golang) focado em desenvolvimento de APIs. Ele é extramente simples e muito parecido com o express do NodeJS.
<br/>
___
**OBS:** Para conseguir entender o manual é necessário saber o básico de GO.
___
<br/>

## Sumário

0. [Fiber](#0-fiber)
1. [Start Server](#1-start-server)
2. [Routes](#2-routes)
3. [Groups](#3-groups)
4. [URL Parameters](#4-url-parameters)
5. [Query String](#5-query-string)
6. [Headers](#6-headers)
7. [JSON Request Response](#7-json-request-response)
8. [Middlewares](#-8middlewares)
9. [Config Server](#-9config-server)
10. [Basic Auth](#10-basic-auth)
11. [JWT HS256](#11-jwt-hs256)
12. [JWT RS256](#12-jwt-rs256)


## 0. Fiber

Como foi dito no início do manual, fiber é um framework em GO com o objetivo de desenvolver APIs. Seu desenvolvimento foi inspirado no Express do NodeJS e tem como características ser muito rápida e ter facíl desenvolvimento.

Iremos utilizar a versão 2 do fiber para este manual.

O documentação oficial do fiber pode ser acessada pelo seguinte link:<br/>
https://docs.gofiber.io/

## 1. Start Server

Antes de iniciar o server é preciso baixar o fiber.

```go
go get github.com/gofiber/fiber/v2
```

Agora iremos iniciar a estrutura básica de um arquivo main em GO já com o fiber importado.

```go
package main

import "github.com/gofiber/fiber/v2"

func main() {

}
```

Dentro da função **main()** será criada uma variável chamada **app** (ou o nome que quiser) que receberá uma instância do fiber.

```go
// CONFIG APP
app := fiber.New()
```

Abaixo do app iremos iniciar uma rota para que possamos testar o fiber.

```go
// ROUTE
app.Get("/", func(c *fiber.Ctx) error {
	return c.SendString("server ok")
})
```

Abaixo da rota iremos iniciar o servidor para que possamos chamar as nossas rotas.

```go
// START SERVER
err := app.Listen(":3000")
if err != nil {
	log.Fatal(err.Error())
}
```

A função **app.Listen()** recebe um valor como parâmetro, que no caso será a porta onde vamos bater para chamar as APIs. A porta deve ser passada como string com um **:** (dois pontos) na frente. Na porta é possível passar **:00** (dois pontos e zero e zero) fazendo que o fiber utilize a primeira porta dispoível que encontrar.

O código final ficará da seguinte forma:

```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// CONFIG APP
	app := fiber.New()

	// ROUTE
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("server ok")
	})

	// START SERVER
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
```

Para subir o servidor bastar rodar o arquivo main do GO.

```go
go run main.go 
```

Ao iniciar um servidor fiber irá aparecer no terminal a seguinte mensagem.

```text
 ┌───────────────────────────────────────────────────┐ 
 │                   Fiber v2.34.0                   │ 
 │               http://127.0.0.1:3000               │ 
 │       (bound on host 0.0.0.0 and port 3000)       │ 
 │                                                   │ 
 │ Handlers ............. 2  Processes ........... 1 │ 
 │ Prefork ....... Disabled  PID ............. 11347 │ 
 └───────────────────────────────────────────────────┘ 
```

Com isso o nosso servidor fiber já está funcionando.

Para testar iremos chamar a rota que criamos, ela pode ser chamada pelo Postman ou diretamente pelo browser. Nesse caso irei utilizar o CURL para testar está rota e as outras que serão utilizadas nesse manual.

```cmd
curl --location --request GET 'localhost:3000/'
```

Iremos receber a seguinte resposta.

```text
server ok
```
## 2. Routes