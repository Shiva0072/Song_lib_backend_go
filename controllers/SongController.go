package controllers

import (
	"learnGo/Crud_CC/dtos"
	"learnGo/Crud_CC/services"
	"log"

	"github.com/gin-gonic/gin"
)

// get all the songs from DB
func GetSongs(c *gin.Context) {
	log.Print("Starting to get all songs")
	songsList, error := services.GetAllSongs()
	log.Print("Ending...")

	if error != nil {
		c.JSON(500, gin.H{
			"error":   error,
			"message": "Please try again later...",
		})
		return
	}

	c.JSON(200, gin.H{
		"song": songsList,
	})
}

// saves the given song in DB
func PostSong(c *gin.Context) {
	var songDto dtos.SongDto

	if err := c.BindJSON(&songDto); err != nil {
		c.JSON(400, gin.H{
			"error": "error in response body",
		})
		return
	}

	songEntity, error := services.SaveSong(&songDto)

	if error != nil {
		c.JSON(400, gin.H{
			"error": error,
		})
		return
	}

	c.JSON(200, gin.H{
		"song saved ": *songEntity,
	})
}

// grabs the songId from pathParam and then finds the corresponding song from DB
func GetSongByGivenId(c *gin.Context) {
	id := c.Param("id")

	songEntity, error := services.GetSongById(id)

	if error != nil {
		c.JSON(400, gin.H{
			"error": "No such song found in DB.",
		})
		return
	}

	c.JSON(200, gin.H{
		"song": *songEntity,
	})
}

// grabs the songId and requestBody. find the songByGivenId, updates the songObject and then saves back in DB
func UpdateSongByGivenId(c *gin.Context) {
	id := c.Param("id")

	var songDto dtos.SongDto

	if err := c.BindJSON(&songDto); err != nil {
		c.JSON(400, gin.H{
			"error": "error in response body",
		})
		return
	}

	songEntity, error := services.UpdateSongById(id, &songDto)

	if error != nil {
		c.JSON(400, gin.H{
			"error": "No such song found in DB. please insert this song.",
		})
		return
	}

	c.JSON(200, gin.H{
		"updated song": *songEntity,
	})
}

// delete song by given id
func DeleteTheSongByGivenId(c *gin.Context) {
	id := c.Param("id")
	isDeleted, error := services.DeleteSongById(id)

	if error != nil || !isDeleted {
		c.JSON(500, gin.H{
			"error":   error,
			"message": "Please try again later...",
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "song deleted successfully!",
	})
}
