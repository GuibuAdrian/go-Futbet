package test_test

import (
	"time"

	"github.com/GuibuAdrian/go-Futbet/dao"
	"github.com/GuibuAdrian/go-Futbet/match"
	"github.com/GuibuAdrian/go-Futbet/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Match", func() {
	var (
		match11 *match.Match
		match12 *match.Match
	)
	BeforeEach(func() {
		riverRepository := dao.GetTeamRepository(models.InitTeam("River", 1))
		newellsRepository := dao.GetTeamRepository(models.InitTeam("Newells", 2))
		location, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
		date := time.Date(2021, 9, 15, 20, 30, 0, 0, location)
		playerDao := dao.PlayerDaoGetInstance()
		p, _ := playerDao.Read(1010101)
		goal4 := models.Goal{Minute: 72, Player: p}
		p, _ = playerDao.Read(1121111)
		goal := models.Goal{Minute: 5, Player: p}
		p, _ = playerDao.Read(3232323)
		goal1 := models.Goal{Minute: 35, Player: p}
		p, _ = playerDao.Read(9999999)
		goal2 := models.Goal{Minute: 61, Player: p}
		goal3 := models.Goal{Minute: 77, Player: p}

		score := models.Score{HomeGoals: []models.Goal{goal1}, AwayGoals: []models.Goal{goal}}

		score.AddAwayGoal(goal2).AddAwayGoal(goal3).AddAwayGoal(goal4)
		match11 = match.InitMatch(newellsRepository, riverRepository, 11, date, score)

		score = models.Score{HomeGoals: []models.Goal{}, AwayGoals: []models.Goal{}}
		match12 = match.InitMatch(newellsRepository, riverRepository, 12, date, score)
	})

	Describe("Scores in match 11", func() {
		Context("Number of goals in away team", func() {
			It("Should be 4", func() {
				Expect(match11.GetAwayTotalGoals()).To(Equal(4))
			})
		})

		Context("Number of goals in home team", func() {
			It("Should be 1", func() {
				Expect(match11.GetHomeTotalGoals()).To(Equal(1))
			})
		})

		Context("Player who score first Goal", func() {
			It("Away should be De la Cruz", func() {
				Expect(match11.GetAwayGoals()[0].GetPlayer().GetName()).To(Equal("De la Cruz"))
			})
			It("Home should be Scocco", func() {
				Expect(match11.GetHomeGoals()[0].GetPlayer().GetName()).To(Equal("Scocco"))
			})
		})

		Context("Player who score third Goal", func() {
			It("Away should be Carrascal", func() {
				goalTest := match11.GetAwayGoalPos(2)
				Expect(goalTest.GetPlayer().GetName()).To(Equal("Carrascal"))
			})
			It("Home should be empty (not found)", func() {
				goalTest := match11.GetHomeGoalPos(2)
				Expect(goalTest).To(Equal(models.Goal{}))
			})
		})

		Context("No Goals", func() {
			It("Away should be empty (not found)", func() {
				goalTest := match12.GetAwayGoalPos(2)
				Expect(goalTest).To(Equal(models.Goal{}))
			})
			It("Home should be empty (not found)", func() {
				goalTest := match12.GetHomeGoalPos(2)
				Expect(goalTest).To(Equal(models.Goal{}))
			})
		})
	})
})
