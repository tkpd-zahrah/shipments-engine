package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/zahrahfebriani/shipments-engine/database"
)

var rsc *database.ShipmentResource

func main() {
	err := godotenv.Load(filepath.Join(".", ".env"))
	if err != nil {
		log.Fatal("failed to load env")
	}

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("Error opening database: %q", err)
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	rsc = database.InitShipment(db)

	router.POST("/shipments/get", func(c *gin.Context) {
		var req GetShipmentsDataRequest
		c.BindJSON(&req)
		resp, err := GetShipmentsData(req)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}

		c.JSON(200, resp)
	})

	router.POST("/shipments/add", func(c *gin.Context) {
		var req AddShipmentRequest
		c.BindJSON(&req)
		resp, err := AddShipmentData(req)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}

		c.JSON(200, resp)
	})

	router.Run(":" + port)
}
