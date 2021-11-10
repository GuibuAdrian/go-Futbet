package bet

type HomeWinner struct {

}

func (bHW *HomeWinner) Win(b Bet) bool {
	return b.match.GetHomeTotalGoals() > b.match.GetAwayTotalGoals()
}