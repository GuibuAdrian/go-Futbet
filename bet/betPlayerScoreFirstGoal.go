package bet

import (
	log "github.com/sirupsen/logrus"
)

type PlayerScoreFirstGoal struct {

}

func (bAW *PlayerScoreFirstGoal) Win(b Bet) bool {
	firstGoal := b.match.GetAllGoalsPos(0)
	idPlayer := firstGoal.GetPlayer().GetId()
	log.Trace(idPlayer)
	return b.PlayerThatScoreFirstGoal.GetId() == idPlayer
}
