package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// {
// 	"id": "1",
// 	"brand": "Honda",
// 	"type": "City"
// }

type Car struct {
	ID    string `json:"id"`
	Brand string `json:"brand"`
	Type  string `json:"type"`
}

var cars = []Car{
	{ID: "1", Brand: "Honda", Type: "City"},
	{ID: "2", Brand: "Hyundai", Type: "Ionic 5"},
}

func main() {
	r := gin.New()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})

	// GET /cars - list cars
	// POST /cars - create cars
	// DELETE /cars/:id - delete cars

	r.Run()
}
