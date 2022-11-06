package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type response struct {
	ID       int64  `json:"id"`
	Response string `json:"response"`
}

func main() {
	router := gin.Default()
	router.POST("/auth/signup", signup)
	router.Run("localhost:8080")
}

func usernameexsisting(star string) bool {
	return true
}
func signup(c *gin.Context) {
	var newuser User
	err := c.BindJSON(&newuser)

	if err != nil {
		return
	}
	if usernameexsisting(newuser.Username) {
		var res response
		res.ID = 1809
		res.Response = "username already exists"
		c.IndentedJSON(http.StatusBadRequest, response{
			ID:       3189,
			Response: "Username already exists",
		})
	}
	c.IndentedJSON(http.StatusCreated, newuser)
}
