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

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

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
	router.Use(CORSMiddleware())

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

	router.POST("/shipments/allocate", func(c *gin.Context) {
		var req AllocationRequest
		c.BindJSON(&req)
		resp, err := AllocateShipment(req)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}

		c.JSON(200, resp)
	})

	router.POST("/shipments/update_status", func(c *gin.Context) {
		var req UpdateStatusShipmentRequest
		c.BindJSON(&req)
		resp, err := UpdateStatusShipment(req)
		if err != nil {
			c.JSON(404, err.Error())
			return
		}

		c.JSON(200, resp)
	})

	router.Run(":" + port)
}
