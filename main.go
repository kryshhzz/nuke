package main

import (
	"fmt"
	"net/http"
	"nuke/attack"
	"io/ioutil"
	"strings"
	"github.com/gin-gonic/gin"
)


func main(){ 
	fmt.Println("Attack Started")

	router := gin.Default()

	router.LoadHTMLGlob("templates/index.html")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Hello from Gin!",
		})
	})
	router.POST("/", func(c *gin.Context) {
		
		body, err := ioutil.ReadAll(c.Request.Body)
        if err != nil {
                fmt.Println("Error reading request body:", err)
                c.JSON(http.StatusBadRequest, gin.H{"error": "Error reading request body"})
                return
        }
        defer c.Request.Body.Close()

        mob_num := strings.SplitN(string(body), "=", 2)
		
		fmt.Println(mob_num[1])

		attack.Attack(mob_num[1])

		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Nuked "+mob_num[1],
		})

	})

	// Run the server on port 8080
	router.Run(":8000")
}
