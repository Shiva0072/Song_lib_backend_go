package daos

import (
	"learnGo/Crud_CC/entities"
	"learnGo/Crud_CC/initializers"
)

func GetAllTheSongs(songsList *[]entities.Song) (bool, error) {
	result := initializers.DB.Find(songsList)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func SaveTheSong(song *entities.Song) (bool, error) {
	result := initializers.DB.Create(song)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func GetTheSongById(song *entities.Song, id string) (bool, error) {
	result := initializers.DB.Find(song, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func UpdateTheSong(song *entities.Song) (bool, error) {
	result := initializers.DB.Save(song)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func DeleteTheSongById(id string) (bool, error) {
	result := initializers.DB.Delete(&entities.Song{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
