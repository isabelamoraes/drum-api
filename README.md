<h1 align="center">
  Drum API
</h1>

<h4 align="center">
  Explorando os conceitos de Clean Arch com Go através de uma aplicação para gerenciar e avaliar kits de bateria.
</h4>

<p align="center">
  <a href="#arquitetura">Clean Arch com Go</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#features">Features</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#tecnologias">Tecnologias</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
  <a href="#setup">Setup</a>
  <a href="#referencias">Referências</a>&nbsp;&nbsp;&nbsp;|&nbsp;&nbsp;&nbsp;
</p>

## Clean Arch com Go

### Organização da aplicação em GO

Geralmente a comunidade Go organiza a aplicação seguindo a seguinte estrutura:
- cmd: arquivos que contém o código de entrada do programa;
- internal: armazena código restrito ao uso do módulo (data, models, handler, service, etc);
- pkg: pacotes externos;

### Clean Architecture

Proposto por Robert Martin, Clean Architecture, ou Arquitetura Limpa, é um padrão arquitetural que consistem em uma série de princípios para estruturar um projeto de forma coesa, testável e de fácil manutenção.

![image](https://github.com/user-attachments/assets/0d1e2aad-5f2a-419b-b87c-490c36aea413)

Referência: [Clean Code Blog](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

A estrutura se baseia nas seguintes camadas:
- Entidades - modelos e regras de negócio da empresa
- Casos de usos - regras de negócio da aplicação
- Adapters - mediação da interação entre as camadas externas centrais da aplicação
- Framework e Drivers - camada mais externa na qual temos a parte de frameworks, banco de dados, etc.

### Clean Arch com Go

Com base nos conceito de Clean Arch e na maneira como a comunidade organiza as aplicações em Go, o projeto respeita a seguinte estrutura:

```
├── cmd
├── internal
│   ├── adapter
|   |   ├── http
│   │   └── logger
│   ├── core
│   │   ├── domain
|   |   |   ├── applicationerror
│   │   |   └── usecase
│   │   └── dto
│   ├── di
│   ├── error
│   ├── infra
|   |   └── gorm
|   |   |   ├── model
|   |   |   └── repositories
│   └── scripts

```

## Features

- [x] CRUD de baterias
- [x] Buscar uma bateria e suas reviews
- [x] Adicionar review para determinada bateria
- [ ] Testes

## Tecnologias

Esse projeto foi desenvolvido utilizando:

-  **[Golang](https://go.dev/)** - Linguagem de programação open-source criada pelo Google;
-  **[Gin](https://gin-gonic.com/)** - Framework web para Golang;
-  **[GORM](https://gorm.io/)** - Biblioteca ORM para Golang.

## Setup

Para clonar e executar essa aplicação insira os comandos abaixo no terminal:

```bash
# Clone this repository
$ git clone https://github.com/isabelamoraes/drum-api.git drum-api

# Go into the repository
$ cd drum-api

# Install dependencies
$ go mod tidy

# Run docker-compose
$ docker-compose up

# Run migrations
$ go run internal/scripts/automigration/migrate.go

# Run api
$ go run cmd/main.go

```

## Referências

- Artigo - Construindo Sistemas com uma Arquitetura Limpa - [Link](https://engsoftmoderna.info/artigos/arquitetura-limpa.html)
- Cursos sobre Go na [Alura](https://www.alura.com.br/)
- Vídeo - Como organizar pastas e arquivo no Go - [Full Cycle](https://www.youtube.com/watch?v=OFud4iPuAH8)
- Artigo - Clear Arch com Golang  - [Link](https://dev.to/booscaaa/implementando-clean-architecture-com-golang-4n0a)
