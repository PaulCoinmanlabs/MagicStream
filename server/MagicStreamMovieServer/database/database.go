package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/PaulCoinmanlabs/MagicStream/server/MagicStreamMovieServer/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

var Client *mongo.Client
var DB *mongo.Database

func Init() {
	cfg := config.App

	// 拼接 MongoDB URI
	// 格式: mongodb://[username:password@]host:port/?authSource=admin
	var uri string
	if cfg.DBUser != "" && cfg.DBPass != "" {
		uri = fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=%s", cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)
	} else {
		uri = fmt.Sprintf("mongodb://%s:%s", cfg.DBHost, cfg.DBPort)
	}

	clientOptions := options.Client().ApplyURI(uri)

	clientOptions.SetMaxPoolSize(100)
	clientOptions.SetMinPoolSize(10)

	// 连接mongoDB
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatalf("[database] MongoDB connect failed: %v", err)
	}

	// MongoDB 驱动强依赖 context，设置 10 秒超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Ping 验证连接是否真正通畅 (Connect 方法本身并不保证网络连通)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("[database] MongoDB ping failed: %v", err)
	}

	Client = client
	DB = client.Database(cfg.DBName)
	log.Println("[database] MongoDB connected successfully")
}
