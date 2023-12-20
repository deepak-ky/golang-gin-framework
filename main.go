package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	
	r := gin.Default()
	
	//r := gin.New()

	r.Use(func (c *gin.Context)  {
		c.Writer.Header().Set("Content-Type", "application/json")
		fmt.Println("Faaltu ka in global middleware")
		//c.Next()
	})


	r.GET("/hello", func (c *gin.Context)  {
		panic("panic happened")
		c.String(http.StatusOK, "Hello world!")
	})
	r.GET("/hello1", func (c *gin.Context)  {
		c.String(http.StatusOK, "Hello world 1!")
	})


	gin.SetMode("release") // gin debug logs would stop printing after this has been set, wherever in the code


	v1 := r.Group("v1")
	v1.GET("/v1-hello", v1HelloHandler)

	r.GET("/redirect", redirectHandler)

	//r.Static("/asdf.png", "/asdf.png")  // needs more understanding

	r.GET("/middleware", AiseHeeMiddleware, middlewareHandler)

	r.Run(":8007")
}

func v1HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, "this is from v1HelloHandler")
}

func redirectHandler(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "http://www.google.com")
}

func AiseHeeMiddleware(c *gin.Context) {
	fmt.Println("mein idhar bhi aaya tha ha ha ha")
	c.String(http.StatusFailedDependency, "yeh middleware se return hua hain")
}

func middlewareHandler(c *gin.Context){
	c.JSON(http.StatusOK, "this is from middlewareHandler")
}