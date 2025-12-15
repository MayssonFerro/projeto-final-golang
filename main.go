package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// definindo a struct do Produto (e criando a tabela no banco de dados)
type Produto struct {
	gorm.Model
	Nome       string
	Quantidade int
	Preco      float64
}

func main() {
	// conexão com o banco de dados (se nao existir, cria o arquivo estoque.db)
	db, err := gorm.Open(sqlite.Open("estoque.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro detalhado:", err)
		panic("falha ao conectar ao banco de dados")
	}

	// auto migration - cria a tabela Produtos se nao existir
	db.AutoMigrate(&Produto{})

	// configuração do gin
	r := gin.Default()

	// rota Hello World
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	// rodar servidor na porta 8080
	r.Run(":8080")
}
