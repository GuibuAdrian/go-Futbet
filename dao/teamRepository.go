package dao

import (
	"errors"
	"github.com/GuibuAdrian/go-Futbet/models"
)

type TeamRepository struct {
	Team models.Team
}

func GetTeamRepository(team models.Team) TeamRepository {
	return TeamRepository{ Team: team}
}
func (teamRep TeamRepository) GetTeam() models.Team { return teamRep.Team }

func (teamRep TeamRepository) GetTeamPlayers() []models.Player {
	playerSlice := PlayerDaoGetInstance().GetPlayerSlice()
	var playerTeamSlice []models.Player

	for _, playerVal := range playerSlice {
		if playerVal.GetTeam() == teamRep.Team {
			playerTeamSlice = append(playerTeamSlice, *playerVal)
		}
	}

	return playerTeamSlice
}

func (teamRep TeamRepository) FindByName(name string) (models.Player, error ) {
	return models.Player{}, errors.New("player not found")
}

func (teamRep TeamRepository) FindById(id int) (models.Player, error ) {
	return models.Player{}, errors.New("player not found")
}
