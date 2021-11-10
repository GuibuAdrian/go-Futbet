package bet

type BetAwayWinner struct {

}

func (bAW *BetAwayWinner) Win(b Bet) bool {
	return b.playerThatScoreFirstGoal.GetNumber() > 1
}
