package templates

func MainTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
{{ if .Legal.Header }}{{ .Legal.Header }}{{ end }}
*/
package main

import "fmt"

func main() {
	fmt.Println("hello,world")
}
`)
}

func WebTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
{{ if .Legal.Header }}{{ .Legal.Header }}{{ end }}
*/
package main

import(
     "net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/index", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"msg": "ok"})
		})
	router.Run(":9991")
}
`)
}
