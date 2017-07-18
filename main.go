package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

var DATABASE_USER = os.Getenv("DATABASE_USER")
var DATABASE_PASSWORD = os.Getenv("DATABASE_PASSWORD")
var DATABASE_NAME = os.Getenv("DATABASE")
var DATABASE_HOST = fmt.Sprintf("%s:5432", os.Getenv("DATABASE_HOST"))
var POOL_SIZE = os.Getenv("POOL_SIZE")

func main() {
	pool_size, err := strconv.Atoi(POOL_SIZE)
	if err != nil {
		pool_size = 1
	}
	db := pg.Connect(&pg.Options{
		User:     DATABASE_USER,
		Password: DATABASE_PASSWORD,
		Database: DATABASE_NAME,
		Addr:     DATABASE_HOST,
		PoolSize: pool_size,
	})
	defer db.Close()

	router := gin.Default()

	// Ping test
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
