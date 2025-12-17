# Controle de Estoque

Este e um projeto simples de controle de estoque desenvolvido em Go. Ele permite listar, adicionar, editar e excluir produtos de um banco de dados SQLite.

## Tecnologias Utilizadas

- Go (Golang)
- Gin (Web Framework)
- GORM (ORM Library)
- SQLite (Banco de dados)
- HTML (Templates)

## Funcionalidades

- Listagem de produtos
- Cadastro de novos produtos
- Edicao de produtos existentes
- Exclusao de produtos

## Como Rodar o Projeto

Siga os passos abaixo para executar a aplicação em sua máquina.

### 1. Inicializar o modulo (caso ainda nao tenha feito)

```bash
go mod init projeto-final
```

### 2. Instalar as dependencias

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get github.com/glebarez/sqlite
```

### 3. Executar a aplicacao

```bash
go run main.go
```

### 4. Acessar no navegador

Abra o seu navegador e acesse o seguinte endereco:

<http://localhost:8080>
