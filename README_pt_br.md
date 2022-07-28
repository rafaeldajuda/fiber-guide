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

Para criar uma rota é preciso utlizar a instância do fiber (app) e chamar qualquer função que tenha o nome de algum método HTTP (POST, PUT, GET, PATCH e DELETE).

```go
app := fiber.New()

app.Get("/get", func(c *fiber.Ctx) error {
	return c.SendString("GET route")
})
```

Na função é preciso passar dois parâmetros. O primeiro é a rota desejada o segundo é uma função anônima, que recebe um contexto do fiber, que é onde fica a lógica da rota.

Na rota é sempre preciso retonar um valor, se não, a requisição não ira ser finalizada.

Exemplos das rotas:

```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/get", func(c *fiber.Ctx) error {
		return c.SendString("GET route")
	})

	app.Post("/post", func(c *fiber.Ctx) error {
		body := string(c.Body())
		return c.SendString(body)
	})

	app.Put("/put", func(c *fiber.Ctx) error {
		body := string(c.Body())
		return c.SendString(body)
	})

	app.Patch("/patch", func(c *fiber.Ctx) error {
		body := string(c.Body())
		return c.SendString(body)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return c.SendString("DELETE route")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}
	
}
```

Exemplo de chamada:
```cmd
curl --location --request GET 'localhost:3000/get'
```

## 3. Groups

É possível criar grupos de rotas. Com eles podemos personalizar e dividir melhor as chamadas das APIs.

Para criar um grupo é preciso chamar a função *Group()* a partir da instância do fiber (app). Na função passamos como parâmetro o caminho inicial das rotas que irão pertencer ao grupo.

```go
app := fiber.New()

// GROUP 1
v1 := app.Group("/v1")
```

Com o grupo criado iremos criar as rotas associadas a ele. As rotas terão que ser criadas a partir da instância do grupo (que nesse caso seria *v1*). Para criar a rota basta chamar qualquer função que tenha como nome métodos HTTP.

```go
v1.Get("/group", func(c *fiber.Ctx) error {
	return c.SendString("group 1")
})
```

Criação de um segundo grupo:

```go
// GROUP 2
v2 := app.Group("/v2")
v2.Get("/group", func(c *fiber.Ctx) error {
	return c.SendString("group 2")
})
```

Exemplo dos grupos:
```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// GROUP 1
	v1 := app.Group("/v1")
	v1.Get("/group", func(c *fiber.Ctx) error {
		return c.SendString("group 1")
	})

	// GROUP 2
	v2 := app.Group("/v2")
	v2.Get("/group", func(c *fiber.Ctx) error {
		return c.SendString("group 2")
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
```

Exemplos de chamadas:
```cmd
curl --location --request GET 'localhost:3000/v1/group'
```

```cmd
curl --location --request GET 'localhost:3000/v2/group'
```

## 4. URL Parameters

Existem várias formas de lidar com os parâmetros de rotas.

A primeira seria uma rota sem parâmetro.

```go
// WITHOUT PARAMETER
app.Get("/parameter", func(c *fiber.Ctx) error {
	return c.SendString("route without parameter")
})
```

A segunda seria uma rota com parâmetro. Nesse caso podemos capturar o valor do parâmetro ao utilizar a função *Params* passando o nome do parâmetro. Para nomear o parâmetro na rota precisamos passar um nome qualquer com um *:* (dois pontos) na frente. 

```go
// WITH PARAMETER
app.Get("/parameter/:item", func(c *fiber.Ctx) error {
	item := c.Params("item")
	return c.SendString(item)
})
```

Podemos também passar um parâmetro opcional na rota. Para isso precisamos utilizar um *?* (ponto de interrogação) no final do nome.

```go
// OPTIONAL PARAMETER
app.Get("/optionalParameter/:item?", func(c *fiber.Ctx) error {
	item := c.Params("item")
	return c.SendString(item)
})
```

Caso não queira passar nenhum nome específico, podemos utiliar o *\** (asterístico). Assim podemos capturar o valor sem especificar algum nome.

```go
// ANY ROUTE (greedy)
app.Get("/anyParameter/*", func(c *fiber.Ctx) error {
	item := c.Params("*")
	return c.SendString(item)
})
```

Podemos criar um parâmetro que recebe um valor na URL e capturar o valor dela por meio do nome do parâmetro.

```go
// OTHER
app.Get("/parameterColor/color::color", func(c *fiber.Ctx) error {
	item := c.Params("verde")
	return c.SendString(item)
})
```

Nesse caso ao chamar a URL *http://localhost:<port>/parameterColor/color:amarelo*, o valor do parâmetro *color* será "amarelo".

Caso o parâmetro passado seja um número podemos capturar esse valor já como um inteiro utilizando a funçao *ParamsInt*.

```go
// OTHER
app.Get("/int/:number", func(c *fiber.Ctx) error {
	number, err := c.ParamsInt("number", 0)
	if err != nil {
		log.Fatal(err.Error())
	}
	return c.SendString(strconv.Itoa(number))
})
```

Exemplo completo:

```go
package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	// WITHOUT PARAMETER
	app.Get("/parameter", func(c *fiber.Ctx) error {
		return c.SendString("route without parameter")
	})

	// WITH PARAMETER
	app.Get("/parameter/:item", func(c *fiber.Ctx) error {
		item := c.Params("item")
		return c.SendString(item)
	})

	// OPTIONAL PARAMETER
	app.Get("/optionalParameter/:item?", func(c *fiber.Ctx) error {
		item := c.Params("item")
		return c.SendString(item)
	})

	// ANY ROUTE (greedy)
	app.Get("/anyParameter/*", func(c *fiber.Ctx) error {
		item := c.Params("*")
		return c.SendString(item)
	})

	// OTHER
	app.Get("/parameterColor/color::color", func(c *fiber.Ctx) error {
		item := c.Params("color")
		return c.SendString(item)
	})

	// OTHER
	app.Get("/int/:number", func(c *fiber.Ctx) error {
		number, err := c.ParamsInt("number", 0)
		if err != nil {
			log.Fatal(err.Error())
		}

		return c.SendString(strconv.Itoa(number))
	})

	// START SERVER
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
```

Exemplos de algumas chamadas:

```cmd
curl --location --request GET 'localhost:3000/parameterColor/color:amarelo'
```

```cmd
curl --location --request GET 'localhost:3000/parameter'
```

```cmd
curl --location --request GET 'localhost:3000/parameter/teste'
```

## 5. Query String

Para capturar as query strings das requisições utilizar a função *c.Query* passando o nome da query como parâmetro. Caso a query não exista será retornado uma string vazia.

```go
// QUERY STRING
app.Get("/queryString", func(c *fiber.Ctx) error {
	item := c.Query("item")
	item2 := c.Query("item2")
	return c.SendString(item + " - " + item2)
})
```

É possível traformar querys strings em estruturas. Para isso primeiro é preciso criar uma estrura com os nomes das querys que irão ser passadas.

```go
// Field names should start with an uppercase letter
type Person struct {
	Name     string   `query:"name"`
	Pass     string   `query:"pass"`
	Products []string `query:"products"`
}
```

Para transformar as querys em estruturas utilizar a função *c.QueryParser* passando como parâmetro uma váriavel do tipo da estrutura da query.

```go
p := new(Person)

if err := c.QueryParser(p); err != nil {
	return err
}
```

Exemplo completo:

```go
package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// Field names should start with an uppercase letter
	type Person struct {
		Name     string   `query:"name"`
		Pass     string   `query:"pass"`
		Products []string `query:"products"`
	}

	app := fiber.New()

	// QUERY STRING
	app.Get("/queryString", func(c *fiber.Ctx) error {
		item := c.Query("item")
		item2 := c.Query("item2")
		return c.SendString(item + " - " + item2)
	})

	// QUERY STRING TO STRUCT
	app.Get("/queryToStruct", func(c *fiber.Ctx) error {
		p := new(Person)

		if err := c.QueryParser(p); err != nil {
			return err
		}

		log.Println(p.Name)
		log.Println(p.Pass)
		log.Println(p.Products)

		return c.Status(fiber.StatusOK).JSON(p)
	})

	err := app.Listen(":3000")
	if err != nil {
		log.Fatal(err.Error())
	}

}
```

Exemplos de chamadas:

```curl
curl --location --request GET 'http://localhost:3000/queryString?item=123&item2=456'
```

```curl
curl --location --request GET 'http://localhost:3000/queryToStruct?name=Rafael&pass=123&products=banana,boots'
```