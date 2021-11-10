package interfaces

import "github.com/GuibuAdrian/go-Futbet/models"

type TeamRepositoryI interface {
	GetTeamPlayers() []models.Player
	FindByName(string) (models.Player, error)
	FindById(int) (models.Player, error)
}