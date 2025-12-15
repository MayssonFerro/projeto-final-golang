package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Nome       string
	Quantidade int
	Preco      float64
}

func main() {
	db, err := gorm.Open(sqlite.Open("estoque.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro detalhado:", err)
		panic("falha ao conectar ao banco de dados")
	}

	db.AutoMigrate(&Produto{})

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, World!",
		})
	})

	r.Run(":8080")
}
