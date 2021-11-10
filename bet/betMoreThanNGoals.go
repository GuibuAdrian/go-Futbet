package bet

type MoreThanNGoals struct {

}

func (bmtng *MoreThanNGoals) Win(b Bet) bool {
	return b.match.GetHomeTotalGoals() > b.totalOfGoals || b.match.GetAwayTotalGoals() > b.totalOfGoals
}