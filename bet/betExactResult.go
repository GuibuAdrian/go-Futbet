package bet

type ExactResult struct {

}

func (bb *ExactResult) Win(b Bet) bool {
	return b.numberOfHomeGoals == b.match.GetHomeTotalGoals() && b.numberOfAwayGoals == b.match.GetAwayTotalGoals()
}
