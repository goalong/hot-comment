package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/goalong/hot-comment/docs"
	"github.com/goalong/hot-comment/router/api"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"os"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())
	//r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	godotenv.Load()
	host := os.Getenv("host")
	url := ginSwagger.URL("http://" + host + "/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	apiv1 := r.Group("/api")

	{
		apiv1.GET("/hot_comments", api.GetHotComments)
		apiv1.GET("/hot_songs", api.GetHotSongs)
		apiv1.GET("/search/songs", api.SearchSongs)
		apiv1.GET("/search/comments", api.SearchComments)
		apiv1.GET("/search/artists", api.SearchArtists)
		apiv1.GET("/songs", api.GetSongsByArtistId)
		apiv1.GET("/comments", api.GetCommentsBySongId)
	}

	return r
}
