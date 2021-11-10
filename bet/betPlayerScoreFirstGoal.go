package bet

import "fmt"

type PlayerScoreFirstGoal struct {

}

func (bAW *PlayerScoreFirstGoal) Win(b Bet) bool {
	firstGoal := b.match.GetAllGoalsPos(0)
	idPlayer := firstGoal.GetPlayer().GetId()
	fmt.Println(idPlayer)
	return b.PlayerThatScoreFirstGoal.GetId() == idPlayer
}
