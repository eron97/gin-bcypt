package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Simula um banco de dados
var database []byte

func main() {
	r := gin.Default()

	r.POST("/register", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Gera o hash da senha com um custo de trabalho (cost) de 14
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar o hash da senha"})
			return
		}

		// Aqui você armazenaria o nome de usuário e o hash da senha no banco de dados
		database = hashedPassword

		c.JSON(http.StatusOK, gin.H{"message": "Usuário registrado com sucesso", "Hash gerado": string(hashedPassword)})
	})

	r.POST("/login", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Compara a senha fornecida com a senha armazenada no banco de dados
		err := bcrypt.CompareHashAndPassword([]byte(database), []byte(user.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Senha incorreta"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Login bem-sucedido"})
	})

	r.Run(":5040")
}

/*

saída


{
    "Hash gerado": "$2a$14$V1wMPWhOVZKJexSJ11577.VfGU6uUxM4waw/nDcqwuXtPfjjH9Pka",
    "message": "Usuário registrado com sucesso"
}


*/
