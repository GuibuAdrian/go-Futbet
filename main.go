package main

import (
	"fmt"
	"github.com/GuibuAdrian/go-Futbet/dao"
	db "github.com/GuibuAdrian/go-Futbet/db"
	"github.com/GuibuAdrian/go-Futbet/models"
	"github.com/kamva/mgm/v3"
)

type Team struct {
	Name	string		`json:"name" bson:"name"`
	IdT		int			`json:"id_t" bson:"id_t"`
	mgm.DefaultModel	`bson:",inline"`
}


func main() {
	db.InitMongoDB()
	teamDao := dao.TeamDaoGetInstance()

	team, _ := teamDao.ReadByName("River")
	fmt.Println(team.GetTeamObjId())

	playerDao := dao.PlayerDaoGetInstance()
	p, _ := playerDao.ReadByPlayerName("Armani", team.GetTeamObjId())
	fmt.Println(p)
	fmt.Println(p.GetName())

	teamPrueba, _ := teamDao.ReadByName("Newells")
	playerPrueba := models.InitPlayer("Prueba2", 000000002, 000000002, "prueba2", teamPrueba)
	//playerDao.Create(playerPrueba)
	playerDao.Update("619ce00c194bc222fcf65c77", playerPrueba)

	playerDao.Delete("619ce00c194bc222fcf65c77")
	//
	//playerDao := dao.PlayerDaoGetInstance()
	//p := playerDao.GetPlayerSlice()[0]
	//goal := models.Goal{Minute: 5, Player: p}
	//p = playerDao.GetPlayerSlice()[0]
	//goal1 := models.Goal{Minute: 35, Player: p}
	//p = playerDao.GetPlayerSlice()[0]
	//goal2 := models.Goal{Minute: 61, Player: p}
	//goal3 := models.Goal{Minute: 77, Player: p}
	//p = playerDao.GetPlayerSlice()[0]
	//goal4 := models.Goal{Minute: 72, Player: p}
	//
	//score := models.Score{ HomeGoals: []models.Goal{goal1}, AwayGoals: []models.Goal{goal} }
	//
	//score.AddAwayGoal(goal2).AddAwayGoal(goal3).AddAwayGoal(goal4)
	//
	//fmt.Println("Away Goals")
	//for _, goal := range score.GetAwayGoals() {
	//	fmt.Printf("Minute: %d Player: %s\n", goal.GetMinute(), goal.GetPlayer().GetName())
	//}
	//fmt.Println("Home Goals")
	//for _, goal := range score.GetHomeGoals() {
	//	fmt.Printf("Minute: %d Player: %s\n", goal.GetMinute(), goal.GetPlayer().GetName())
	//}
	//
	//riverRepository := dao.GetTeamRepository(models.InitTeam("River", 1))
	//newellsRepository := dao.GetTeamRepository(models.InitTeam("Newells", 2))
	//location, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
	//date := time.Date(2021, 9, 15, 20, 30, 0, 0, location)
	//match11 := match.InitMatch(newellsRepository, riverRepository, 11, date, score)
	//
	//fmt.Println(match11.GetHomeTotalGoals())
	//fmt.Println(match11.GetAwayTotalGoals())
	//
	//g := models.InitGambler("Han Solo", 1)
	//
	//playerThatScoreGoal := playerDao.GetPlayerSlice()[0]
	//bet2 := bet.NewBetBuilder(1000, *g, *match11).WithPlayerThatScoreFirstGoal(playerThatScoreGoal).Build()
	//bpsfg := bet.PlayerScoreFirstGoal{}
	//
	//fmt.Println(bpsfg.Win(*bet2))
}
