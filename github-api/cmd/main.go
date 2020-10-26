package main

import (
	"../cache"
	"../clients"
	"github.com/gin-gonic/gin"
	"log"
	goHttp "net/http"
	"time"
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

	ttlCache := cache.NewTTLCache(time.Minute, 1000)
	client := clients.GithubClient{Client: goHttp.Client{}, TtlCache: ttlCache, Url: "https://api.github.com"}

	router.GET("/repositories", func(c *gin.Context) {
		repositories, err := client.TtlCache.Get("repositories")

		if err != nil {
			c.String(500, err.Error())
		}
		if repositories != nil {
			c.JSON(200, repositories)
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
			gitRepositories, err := client.GetRepositories()
			if err != nil {
				c.String(500, err.Error())
			}
			err = client.TtlCache.Put("repositories", gitRepositories)
			if err != nil {
				c.String(500, err.Error())
			}
			c.JSON(200, gitRepositories)
		}
	})

	router.GET("/commits/:id", func(c *gin.Context) {
		commits, err := client.TtlCache.Get("commits:"+c.Param("id"))

		if err != nil {
			c.String(500, err.Error())
		}
		if commits != nil {
			c.JSON(200, commits)
		} else {

		c.Header("Access-Control-Allow-Origin", "*")
			gitCommits, err := client.GetCommits(c.Param("id"))
			if err != nil {
				c.String(500, err.Error())
			}
			err = client.TtlCache.Put("commits:"+c.Param("id"), gitCommits)
			if err != nil {
				c.String(500, err.Error())
			}
			c.JSON(200, gitCommits)
		}
	})
	return router
}
