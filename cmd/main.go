package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
	"github.com/nashrull/averin/shared/config"
	siswa "github.com/nashrull/averin/siswa/entrypoint"
)

func main() {
	engine := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowCredentials = true
	corsCfg.AllowHeaders = []string{"*"}
	engine.Use(cors.New(corsCfg))

	logger := log.Default()

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Host", os.Getenv("HOST"))
	fmt.Println("Port", os.Getenv("PORT"))
	fmt.Println("Username", os.Getenv("USER"))
	fmt.Println("Password", os.Getenv("PASSWORD"))
	fmt.Println("Name", os.Getenv("NAME"))
	fmt.Println("Driver", os.Getenv("DRIVER"))
	cfg := config.Config{
		DB: config.DBConfig{
			Host:     os.Getenv("HOST"),
			Port:     os.Getenv("PORT"),
			Username: os.Getenv("USER"),
			Password: os.Getenv("PASSWORD"),
			Name:     os.Getenv("NAME"),
			Driver:   os.Getenv("DRIVER"),
		},
	}

	err = siswa.RegisterSiswaModule(logger, engine, cfg)
	if err != nil {
		logger.Println("error ", err.Error())
		return
	}

	engine.Run(":8000")
}
