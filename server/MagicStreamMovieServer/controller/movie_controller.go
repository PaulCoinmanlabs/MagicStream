package controller

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/PaulCoinmanlabs/MagicStream/server/MagicStreamMovieServer/database"
	"github.com/PaulCoinmanlabs/MagicStream/server/MagicStreamMovieServer/models"
	"github.com/PaulCoinmanlabs/MagicStream/server/MagicStreamMovieServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// /movies/allMovies
func GetAllMovies(ctx *gin.Context) {
	// 注意：不要直接使用 Gin 的 ctx 作为 MongoDB 的 Context，因为 HTTP 请求断开不代表数据库查询应该立刻被强杀
	dbCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// 这里默认是0 避免了null
	movies := make([]models.Movie, 0)

	cursor, err := database.DB.Collection("movies").Find(dbCtx, bson.D{})

	if err != nil {
		log.Printf("[GetAllMovies] Find execution failed: %v", err)
		utils.Fail(ctx, http.StatusBadRequest, "查找失败")
		return

	}

	defer cursor.Close(dbCtx)

	if err := cursor.All(ctx, &movies); err != nil {
		log.Printf("[GetAllMovies] Find execution failed: %v", err)
		utils.Fail(ctx, http.StatusBadRequest, "数据解析失败")
		return
	}

	utils.OK(ctx, movies)
}

// /movies/movies/:imdbId
func GetMovies(ctx *gin.Context) {
	id := ctx.Param("imdbId")

	dbCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var movie models.Movie

	err := database.DB.Collection("movies").FindOne(dbCtx, bson.M{"imdb_id": id}).Decode(&movie)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			utils.Fail(ctx, http.StatusBadRequest, "找不到该电影")
			return
		}
		// 其他错误（如网络断开、数据库宕机等）
		log.Printf("[GetMovieByID] FindOne failed: %v", err)
		utils.Fail(ctx, http.StatusBadRequest, "找不到该电影")
		return
	}

	utils.OK(ctx, movie)
}
