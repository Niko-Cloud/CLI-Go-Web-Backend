package main

import (
	"CliPorto/internal/config"
	"CliPorto/internal/database"
	"CliPorto/internal/router"
	"github.com/gin-contrib/cors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}

	database.RunMigrations(db)

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	_ = r.SetTrustedProxies(nil)

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:8191",
		},
		AllowMethods: []string{"GET"},
		AllowHeaders: []string{"Content-Type", "Authorization"},
		MaxAge:       12 * time.Hour,
	}))

	router.RegisterRoutes(r, db)

	log.Printf("Server running on :%s", cfg.Server.Port)
	r.Run(":" + cfg.Server.Port)
}
