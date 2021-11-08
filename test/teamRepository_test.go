package test_test

import (
	"github.com/GuibuAdrian/go-Futbet/dao"
	"github.com/GuibuAdrian/go-Futbet/models"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TeamRepository", func() {
	var (
		riverRepository dao.TeamRepository
		ferroRepository	dao.TeamRepository
	)

	BeforeEach(func() {
		riverRepository = dao.GetTeamRepository(models.InitTeam("River", 1))
		ferroRepository = dao.GetTeamRepository(models.InitTeam("Ferro", 4))
	})

	Describe("Number of players in Team", func() {
		Context("Team with 11 players", func() {
			It("Should be 11", func() {
					Expect(len(riverRepository.GetTeamPlayers())).To(Equal(11))
				})
		})

		Context("Team with 0 players", func() {
			It("Should be 0", func() {
				Expect(len(ferroRepository.GetTeamPlayers())).To(Equal(0))
			})
		})
	})
})
