package main

import (
	"log"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	siswa "github.com/nashrull/averin/siswa/entrypoint"
)

func main() {
	// ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	// defer cancel()

	engine := gin.Default()
	corsCfg := cors.DefaultConfig()
	corsCfg.AllowAllOrigins = true
	corsCfg.AllowCredentials = true
	corsCfg.AllowHeaders = []string{"*"}
	engine.Use(cors.New(corsCfg))

	logger := log.Default()
	err := siswa.RegisterSiswaModule(logger, engine)
	if err != nil {
		logger.Println("error ", err.Error())
		return
	}

	engine.Run(":8000")
}
