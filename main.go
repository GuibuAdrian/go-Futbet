package main

import (
	"fmt"
	"github.com/GuibuAdrian/goFutbet/dao"
	"github.com/GuibuAdrian/goFutbet/match"
	"github.com/GuibuAdrian/goFutbet/models"
	"time"
)

func main() {
	playerDao := dao.PlayerDaoGetInstance()
	p, _ :=playerDao.Read(1121111)
	goal := models.Goal{Minute: 5, Player: p}
	p, _ = playerDao.Read(3232323)
	goal1 := models.Goal{Minute: 35, Player: p}
	p, _ = playerDao.Read(9999999)
	goal2 := models.Goal{Minute: 61, Player: p}
	goal3 := models.Goal{Minute: 77, Player: p}
	p, _ = playerDao.Read(1010101)
	goal4 := models.Goal{Minute: 72, Player: p}

	score := models.Score{ HomeGoals: []models.Goal{goal1}, AwayGoals: []models.Goal{goal} }

	score.AddAwayGoal(goal2).AddAwayGoal(goal3).AddAwayGoal(goal4)

	fmt.Println("Away Goals")
	for _, goal := range score.GetAwayGoals() {
		fmt.Printf("Minute: %d Player: %s\n", goal.GetMinute(), goal.GetPlayer().GetName())
	}
	fmt.Println("Home Goals")
	for _, goal := range score.GetHomeGoals() {
		fmt.Printf("Minute: %d Player: %s\n", goal.GetMinute(), goal.GetPlayer().GetName())
	}

	riverRepository := dao.GetTeamRepository(models.InitTeam("River", 1))
	newellsRepository := dao.GetTeamRepository(models.InitTeam("Newells", 2))
	location, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	date := time.Date(2021, 9, 15, 20, 30, 0, 0, location)
	match11 := match.InitMatch(newellsRepository, riverRepository, 11, date, score)

	fmt.Println(match11.GetHomeTotalGoals())
	fmt.Println(match11.GetAwayTotalGoals())
}
