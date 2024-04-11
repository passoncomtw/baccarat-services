package models

import (
	"gorm.io/gorm"
)

type Lobby struct {
	gorm.Model
	Id   uint   `gorm:"primaryKey;comment: '主鍵'"`
	Name string `gorm:"type:varchar(40);comment: '大廳名稱'"`
}

type CreateLobbyRequestDto struct {
	name string
}

func createLobby(values *CreateLobbyRequestDto) Lobby {
	var lobby Lobby
	lobby.Name = values.name
	DBConnection.Create(&lobby)
	return lobby
}
