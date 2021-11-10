package bet

type AwayWinner struct {

}

func (baw *AwayWinner) Win(b Bet) bool {
	return b.match.GetAwayTotalGoals() > b.match.GetHomeTotalGoals()
}