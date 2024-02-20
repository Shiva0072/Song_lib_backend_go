package main

import (
	"learnGo/Crud_CC/controllers"
	"learnGo/Crud_CC/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
}

func main() {
	r := gin.Default()
	r.GET("/get-all-songs", controllers.GetSongs)
	r.POST("/save-song", controllers.PostSong)
	r.GET("/get-song-by-id/:id", controllers.GetSongByGivenId)
	r.PUT("/update-song-by-id/:id", controllers.UpdateSongByGivenId)
	r.DELETE("/delete-song-by-id/:id", controllers.DeleteTheSongByGivenId)
	r.Run() // listen and serve on 0.0.0.0:8080
}
