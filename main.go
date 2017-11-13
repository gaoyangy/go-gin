package main

import (
	"github.com/gin-gonic/gin"
	_"github.com/go-sql-driver/mysql"
	"database/sql"
	"gin-demo/Controller"
	"fmt"
)

func main() {
	db, errs := sql.Open("mysql", "root:123@tcp(127.0.0.1:3306)/im")
	if errs == nil {
		fmt.Printf("连接 success")	
	}
	defer db.Close()

	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	
	r.GET("/",Controller.Add)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}