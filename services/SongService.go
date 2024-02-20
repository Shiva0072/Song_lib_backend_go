package services

import (
	"learnGo/Crud_CC/daos"
	"learnGo/Crud_CC/dtos"
	"learnGo/Crud_CC/entities"
)

func SaveSong(songDto *dtos.SongDto) (*entities.Song, error) {

	var song entities.Song = entities.Song{}
	song.SetName(songDto.Name)
	song.SetSinger(songDto.Singer)
	song.SetWriter(songDto.Writer)
	song.SetStatus(songDto.Status)

	isSongSaved, error := daos.SaveTheSong(&song)

	if isSongSaved {
		return &song, nil
	}

	return nil, error
}

func GetAllSongs() (*[]entities.Song, error) {
	var songList []entities.Song

	allSongsRetrieved, error := daos.GetAllTheSongs(&songList)

	if allSongsRetrieved {
		return &songList, nil
	}
	return nil, error
}

func GetSongById(id string) (*entities.Song, error) {
	var song entities.Song
	songRetrieved, error := daos.GetTheSongById(&song, id)

	if songRetrieved {
		return &song, nil
	}
	return nil, error
}

func UpdateSongById(id string, songDto *dtos.SongDto) (*entities.Song, error) {
	songEntity, error := GetSongById(id)

	if error != nil {
		return nil, error
	}

	if songDto.Name != "" {
		songEntity.SetName(songDto.Name)
	}
	if songDto.Singer != "" {
		songEntity.SetSinger(songDto.Singer)
	}
	if songDto.Writer != "" {
		songEntity.SetWriter(songDto.Writer)
	}
	if songDto.Status != "" {
		songEntity.SetStatus(songDto.Status)
	}

	songRetrieved, error := daos.UpdateTheSong(songEntity)

	if songRetrieved {
		return songEntity, nil
	}
	return nil, error
}

func DeleteSongById(id string) (bool, error) {
	isDeleted, error := daos.DeleteTheSongById(id)

	if !isDeleted {
		return false, error
	}

	return true, nil
}
