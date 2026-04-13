package main

import (
	"log"
	"net/http"

	"github.com/PaulCoinmanlabs/MagicStream/server/MagicStreamMovieServer/config"
	"github.com/PaulCoinmanlabs/MagicStream/server/MagicStreamMovieServer/database"
	"github.com/PaulCoinmanlabs/MagicStream/server/MagicStreamMovieServer/routers"
)

func main() {
	// 读取配置信息
	config.Load()
	database.Init()

	router := routers.NewRouter()

	addr := ":" + config.App.AppPort
	log.Printf("[server] starting on http://localhost%s", addr)

	src := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	if err := src.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("[server] failed: %v", err)
	}
}
