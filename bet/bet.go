package bet

import (
	"github.com/GuibuAdrian/go-Futbet/match"
	"github.com/GuibuAdrian/go-Futbet/models"
)

type Bet struct {
	amountGamble				int
	gambler						models.Gambler
	match						match.Match
	//Optional
	PlayerThatScoreFirstGoal	models.Player
	totalOfGoals				int
	numberOfHomeGoals			int
	numberOfAwayGoals			int
}

type Builder struct {
	bet *Bet
}

func NewBetBuilder(amount int, g models.Gambler, match2 match.Match) *Builder {
	return &Builder{bet: InitBet(amount , g, match2)}
}

func InitBet(amount int, g models.Gambler, match2 match.Match) *Bet {
	return &Bet{
		amountGamble: amount,
		gambler: g,
		match: match2,
	}
}

func (bb *Builder) WithPlayerThatScoreFirstGoal(p models.Player) *Builder {
	bb.bet.PlayerThatScoreFirstGoal = p
	return bb
}

func (bb *Builder) WithNumberOfGoals(numberGoals int) *Builder {
	bb.bet.totalOfGoals = numberGoals
	return bb
}

func (bb *Builder) WithNumberOfHomeAndAwayGoals(homeGoals, awayGoals int) *Builder {
	bb.bet.numberOfHomeGoals = homeGoals
	bb.bet.numberOfAwayGoals = awayGoals
	return bb
}

func (bb *Builder) Build() *Bet {
	return bb.bet
}