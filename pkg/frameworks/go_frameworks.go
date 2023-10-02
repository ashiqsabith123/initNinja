package frameworks

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Gin(project_name, architecture string) {

	var filename string

	switch architecture {

	case "Clean Code Architecture":
		filename = "/cmd/api/main.go"
	case "MVC":
		filename = "/main.go"

	}
	cmd := exec.Command("go", "get", "-u", "github.com/gin-gonic/gin")
	cmd.Dir = project_name + "/"

	_, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Println("Error executing command:", err)
		return
	}

	file, err := os.OpenFile(project_name+filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	code := `package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}`

	file.WriteString(code)

}

func Gorilla(project_name string) {

}
