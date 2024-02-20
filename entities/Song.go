package entities

import "time"

type Song struct {
	Id        uint `gorm:"primaryKey"`
	Name      string
	Writer    string
	Singer    string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (song *Song) GetId() uint {
	return song.Id
}

func (song *Song) GetCreatedAt() time.Time {
	return song.CreatedAt
}

func (song *Song) GetUpdatedAt() time.Time {
	return song.UpdatedAt
}

func (song *Song) GetStatus() string {
	return song.Status
}

func (song *Song) GetName() string {
	return song.Name
}

func (song *Song) GetSinger() string {
	return song.Singer
}

func (song *Song) GetWriter() string {
	return song.Writer
}

func (song *Song) SetName(name string) {
	song.Name = name
}
func (song *Song) SetWriter(writer string) {
	song.Writer = writer
}
func (song *Song) SetSinger(singer string) {
	song.Singer = singer
}
func (song *Song) SetStatus(status string) {
	song.Status = status
}
