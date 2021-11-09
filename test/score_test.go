package test_test

import (
	"github.com/GuibuAdrian/go-Futbet/dao"
	"github.com/GuibuAdrian/go-Futbet/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Score", func() {

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

	Describe("Number of goals in score", func() {
		Context("Number of Home Goals", func() {
			It("Should be 1", func() {
				Expect(len(score.GetHomeGoals())).To(Equal(1))
			})
		})

		Context("Number of Away Goals", func() {
			It("Should be 4", func() {
				Expect(len(score.GetAwayGoals())).To(Equal(4))
			})
		})
	})

	Describe("Score goals", func() {
		Context("Home goal Minute", func() {
			It("Should be 5", func() {
				Expect(goal.GetMinute()).To(Equal(5))
			})
		})

		Context("Player who score goal", func() {
			It("Should be De la Cruz", func() {
				Expect(goal.GetPlayer().GetName()).To(Equal("De la Cruz"))
			})
		})
	})
})
