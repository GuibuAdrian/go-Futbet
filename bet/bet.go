package models

type Bet struct {
	amountGamble				int
	gambler                 	Gambler
	PlayerThatScoreFirstGoal	Player
	numberOfGoals				int
}

type Builder struct {
	bet *Bet
}

func InitBet(amount int, g Gambler) *Bet {
	return &Bet{
		amountGamble: amount,
		gambler: g,
	}
}

func (bb *Builder) WithPlayerThatScoreFirstGoal(p Player) *Builder {
	bb.bet.PlayerThatScoreFirstGoal = p
	return bb
}

func (bb *Builder) WithNumberOfGoals(numberGoals int) *Builder {
	bb.bet.numberOfGoals = numberGoals
	return bb
}

func (bb *Builder) Build() *Bet {
	return bb.bet
}