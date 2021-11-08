package dao

import (
	"errors"
	"github.com/GuibuAdrian/goFutbet/models"
)

type TeamDao struct {
	teamSlice []models.Team
}

var teamDaoInstance *TeamDao

func TeamDaoGetInstance() *TeamDao {
	if teamDaoInstance == nil {
		teamDaoInstance = &TeamDao{}

		initializeTeamDao()
	}
	return teamDaoInstance
}

func (teamDao *TeamDao) Create(team models.Team) {
	teamDao.teamSlice = append(teamDao.teamSlice, team)
}

func (teamDao TeamDao) Read(id int) (models.Team, error) {
	team := models.Team{}
	for _, team := range teamDao.teamSlice {
		if team.GetTeamId() == id {
			return team, nil
		}
	}

	return team, errors.New("team not found")
}

func (teamDao TeamDao) Update(team models.Team)  {
	//Todo
}

func (teamDao *TeamDao) Delete(team models.Team)  {
	for posV, teamVal := range teamDao.teamSlice {
		if teamVal == team{
			teamDao.teamSlice = append(teamDao.teamSlice[:posV], teamDao.teamSlice[posV+1:]...)
			break
		}
	}
}

func (teamDao TeamDao) GetTeamSlice() []models.Team {
	return teamDao.teamSlice
}

func initializeTeamDao() {
	TeamDaoGetInstance().Create(models.InitTeam("River", 1))
	TeamDaoGetInstance().Create(models.InitTeam("Newells", 2))
	TeamDaoGetInstance().Create(models.InitTeam("Banfield", 3))
	TeamDaoGetInstance().Create(models.InitTeam("Velez", 4))
}