package test_test

import (
	"github.com/GuibuAdrian/go-Futbet/bet"
	"github.com/GuibuAdrian/go-Futbet/dao"
	"github.com/GuibuAdrian/go-Futbet/match"
	"github.com/GuibuAdrian/go-Futbet/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

var _ = Describe("Bet", func() {
	var (
		match11 *match.Match
		gamb *models.Gambler
		betBuilder *bet.Builder
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

		gamb = models.InitGambler("Han Solo", 1)
		betBuilder = bet.NewBetBuilder(1000, *gamb, *match11)
	})

	Describe("Bet", func() {
		Context("First player who score goal in the match", func() {
			It("Should be true", func() {
				playerThatScoreGoal, _ := dao.PlayerDaoGetInstance().Read(1121111)
				bet2 := betBuilder.WithPlayerThatScoreFirstGoal(playerThatScoreGoal).Build()

				bpsfg := bet.PlayerScoreFirstGoal{}
				Expect(bpsfg.Win(*bet2)).To(Equal(true))
			})
			It("Should be false", func() {
				playerThatScoreGoal, _ := dao.PlayerDaoGetInstance().Read(3232323)
				bet2 := betBuilder.WithPlayerThatScoreFirstGoal(playerThatScoreGoal).Build()

				bpsfg := bet.PlayerScoreFirstGoal{}
				Expect(bpsfg.Win(*bet2)).To(Equal(false))
			})
		})

		Context("Team Winner", func() {
			It("Should be false", func() {
				bet2 := betBuilder.Build()
				bhw := bet.HomeWinner{}
				Expect(bhw.Win(*bet2)).To(Equal(false))
			})
			It("Should be true", func() {
				bet2 := betBuilder.Build()
				baw := bet.AwayWinner{}
				Expect(baw.Win(*bet2)).To(Equal(true))
			})

			Context("More than N Goals", func() {
				It("Should be true", func() {
					bet2 := betBuilder.WithNumberOfGoals(3).Build()
					bmtng := bet.MoreThanNGoals{}
					Expect(bmtng.Win(*bet2)).To(Equal(true))
				})
				It("Should be false", func() {
					bet2 := betBuilder.WithNumberOfGoals(5).Build()
					bmtng := bet.MoreThanNGoals{}
					Expect(bmtng.Win(*bet2)).To(Equal(false))
				})
			})

			Context("Exact result bet", func() {
				It("Should be true", func() {
					bet2 := betBuilder.WithNumberOfHomeAndAwayGoals(1, 4).Build()
					ber := bet.ExactResult{}
					Expect(ber.Win(*bet2)).To(Equal(true))
				})
				It("Should be false", func() {
					bet2 := betBuilder.WithNumberOfHomeAndAwayGoals(2, 1).Build()
					ber := bet.ExactResult{}
					Expect(ber.Win(*bet2)).To(Equal(false))
				})
			})
		})
	})
})
