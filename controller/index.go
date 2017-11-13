package Controller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gin-demo/Module"
	"fmt"
)

func Add (c *gin.Context) {
	msgs := Module.Getuser()
	fmt.Println(msgs)
  c.JSON(http.StatusOK, gin.H{
		"msg": msgs,
	 })
}