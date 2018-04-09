package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

func main() {

    gin.SetMode(gin.DebugMode)
    router := gin.Default()

    router.Use(Middleware)

    router.GET("/simple/server/get", GetHandler)
    router.POST("/simple/server/post", PostHandler)
    router.PUT("/simple/server/put", PutHandler)
    router.DELETE("/simple/server/delete", DeleteHandler)

    http.ListenAndServe(":8005", router)
}

func Middleware(c *gin.Context) {
    fmt.Println("this is a middleware!")
}

func GetHandler(c *gin.Context) {
    value, exist := c.GetQuery("key")
    if !exist {
        value = "the key is not exist!"
    }
    c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success! %s\n", value)))
    return
}
func PostHandler(c *gin.Context) {
    type JsonHolder struct {
        Id   int    `json:"id"`
        Name string `json:"name"`
    }
    holder := JsonHolder{Id: 1, Name: "my name"}

    c.JSON(http.StatusOK, holder)
    return
}
func PutHandler(c *gin.Context) {
    c.Data(http.StatusOK, "text/plain", []byte("put success!\n"))
    return
}
func DeleteHandler(c *gin.Context) {
    c.Data(http.StatusOK, "text/plain", []byte("delete success!\n"))
    return
}