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

	err := router.Run(":8090")
	if err != nil {
		log.Fatal(err)
	}

}
func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/repositories", getRepositories)
	router.GET("/commits/:id", getCommits)
	return router
}

func getRepositories(c *gin.Context) {
	client :=
		clients.GithubClient{
			Client: goHttp.Client{},
			Url:    "https://api.github.com"}
	gitRepositories, err := client.GetRepositories()
	c.Header("Access-Control-Allow-Origin", "*")
	if err != nil {
		c.String(500, err.Error())
	}
	c.JSON(200, gitRepositories)
}

func getCommits(c *gin.Context) {
	client :=
		clients.GithubClient{
			Client: goHttp.Client{},
			Url:    "https://api.github.com"}
	c.Header("Access-Control-Allow-Origin", "*")
	gitCommits, err := client.GetCommits(c.Param("id"))
	if err != nil {
		c.String(500, err.Error())
	}
	c.JSON(200, gitCommits)
}
