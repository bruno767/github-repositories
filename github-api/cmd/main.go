package main

import (
	"../clients"
	"github.com/gin-gonic/gin"
	"log"
	goHttp "net/http"
)

func main() {

	gin.ForceConsoleColor()

	router := setupRouter()

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

}
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/repositories", getRepositories)
	return router
}

func getRepositories(c *gin.Context) {
	client :=
		clients.GithubClient{
			Client: goHttp.Client{},
			Url:    "https://api.github.com"}
	gitRepositories, err := client.GetRepositories()
	if err != nil {
		c.String(500, err.Error())
	}
	c.JSON(200, gitRepositories)
}
