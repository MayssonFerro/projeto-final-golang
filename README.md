# Para rodar o projeto

## Iniciar o gerenciador de módulos do Go

```bash
go mod init projeto-final
```

## Baixe as dependencias

```bash
go get -u github.com/gin-gonic/gin

go get -u gorm.io/gorm
go get github.com/glebarez/sqlite
```

## Rodar

```bash
go run main.go
```

Acesse no navegador: <http://localhost:8080>

você verá o {"message":"Hello, World!"}

o arquivo estoque.db deve ter aparecido
