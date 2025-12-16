package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

type Produto struct {
	gorm.Model
	Nome       string  `json:"nome" form:"Nome"`
	Quantidade int     `json:"quantidade" form:"Quantidade"`
	Preco      float64 `json:"preco" form:"Preco"`
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("estoque.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("Erro detalhado:", err)
		panic("falha ao conectar ao banco de dados")
	}

	db.AutoMigrate(&Produto{})

	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		var produtos []Produto
		db.Find(&produtos)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"produtos": produtos,
		})
	})

	r.GET("/novo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "form.html", gin.H{
			"produto": Produto{},
		})
	})

	r.GET("/editar/:id", func(c *gin.Context) {
		id := c.Param("id")
		var produto Produto
		if err := db.First(&produto, id).Error; err != nil {
			c.String(http.StatusNotFound, "Produto não encontrado")
			return
		}

		c.HTML(http.StatusOK, "form.html", gin.H{
			"produto": produto,
		})
	})

	r.GET("/sucesso", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sucesso.html", nil)
	})

	r.POST("/salvar", func(c *gin.Context) {
		var produto Produto
		if err := c.ShouldBind(&produto); err != nil {
			c.String(http.StatusBadRequest, "Dados inválidos")
			return
		}

		idStr := c.PostForm("ID")
		id, _ := strconv.Atoi(idStr)

		if id > 0 {
			produto.ID = uint(id)
			db.Save(&produto)
		} else {
			db.Create(&produto)
		}

		c.Redirect(http.StatusFound, "/sucesso")
	})

	r.GET("/deletar/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&Produto{}, id)
		c.Redirect(http.StatusFound, "/")
	})

	r.Run(":8080")
}
