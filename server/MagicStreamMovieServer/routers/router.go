package routers

import (
	"github.com/PaulCoinmanlabs/MagicStream/server/MagicStreamMovieServer/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	movie := v1.Group("/movies")
	{
		movie.GET("/allMovies", controller.GetAllMovies)
		movie.GET("/movies/:imdbId", controller.GetMovies)
	}

	return r
}
