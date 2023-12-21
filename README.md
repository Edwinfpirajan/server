# Biblioteca Base para Servidores

La Biblioteca Base para Servidores es un conjunto de herramientas y estructuras diseñadas para facilitar el desarrollo de servidores en Go. Proporciona funcionalidades esenciales como la gestión de roles, usuarios, adaptadores SQL y utiliza GORM para la interacción con bases de datos SQL.

## Estructura del Proyecto

.
├── cmd
│ ├── main.go
│ └── providers
│ └── di.go
├── config
│ └── enviroment.go
├── go.mod
├── go.sum
├── internal
│ ├── app
│ │ ├── role.go
│ │ └── user.go
│ ├── domain
│ │ ├── dto
│ │ │ ├── role.go
│ │ │ └── user.go
│ │ ├── entity
│ │ │ ├── paginate.go
│ │ │ ├── response.go
│ │ │ └── user.go.txt
│ │ └── ports
│ │ └── db
│ │ └── interfaces
│ │ ├── role.go
│ │ └── user.go
│ ├── infra
│ │ ├── adapters
│ │ │ └── db
│ │ │ └── implementation
│ │ │ ├── role.go
│ │ │ ├── sqlBuilder.go
│ │ │ └── user.go
│ │ └── api
│ │ ├── handlers
│ │ │ ├── health.go
│ │ │ ├── role.go
│ │ │ └── user.go
│ │ └── router
│ │ ├── groups
│ │ │ ├── role.go
│ │ │ └── user.go
│ │ └── router.go
│ └── src
│ └── db
│ └── dbConnectionPostgres.go
├── README.md
├── sqlmigrations
│ └── smartdb.sql
└── utils
└── http
└── http.go

## Características

- Gestión de roles y usuarios.
- Adaptadores SQL flexibles.
- Integración con GORM para operaciones de base de datos.

## Instalación

Primero, asegúrate de tener [Go](https://golang.org/doc/install) instalado en tu máquina.

### Instalación de Echo

Instala el framework [Echo](https://echo.labstack.com/) utilizando el siguiente comando:

```bash
go get -u github.com/labstack/echo/v4
