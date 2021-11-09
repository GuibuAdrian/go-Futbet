package models

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestModels(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Models Suite")
}

var _ = Describe("Player", func() {
	var (
		riverPlayer Player
		newellsPlayer Player
	)

	BeforeEach(func() {
		riverPlayer = Player{
			name: "Alvarez",
			id: 9999999,
			number: 9,
			position: "delantero",
			team: Team{	Name: "River", Id: 1 },
		}

		newellsPlayer = Player{
			name: "Scocco",
			id: 3232323,
			number: 32,
			position: "delantero",
			team: Team{ Name: "Newells", Id: 2 },
		}
	})

	Describe("Player Team", func() {
		Context("In River", func() {
			It("Should be in River", func() {
				Expect(riverPlayer.GetTeam().GetTeamName()).To(Equal("River"))
			})
		})

		Context("In Newells", func() {
			It("Should be in Newells", func() {
				Expect(newellsPlayer.GetTeam().GetTeamName()).To(Equal("Newells"))
			})
		})
	})
})