package interfaces

import "github.com/GuibuAdrian/goFutbet/models"

type TeamRepositoryI interface {
	GetTeamPlayers() []models.Player
	FindByName(string) (models.Player, error)
	FindById(int) (models.Player, error)
}