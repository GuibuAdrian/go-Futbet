package match

import (
	"github.com/GuibuAdrian/goFutbet/interfaces"
	"github.com/GuibuAdrian/goFutbet/models"
	"sort"
	"time"
)

type Match struct {
	home		interfaces.TeamRepositoryI
	away		interfaces.TeamRepositoryI
	matchNumber	int
	dateMatch 	time.Time
	score     	models.Score
}

func InitMatch(home, away interfaces.TeamRepositoryI, matchN int, date time.Time, s models.Score) *Match {
	return &Match{
		home: home,
		away: away,
		matchNumber: matchN,
		dateMatch: date,
		score: s,
	}
}

func (m *Match) GetHomeTotalGoals() int { return len(m.score.GetHomeGoals()) }
func (m *Match) GetAwayTotalGoals() int { return len(m.score.GetAwayGoals()) }

func (m *Match) GetHomeGoals() []models.Goal {
	sort.Slice(m.score.GetHomeGoals(), func(i, j int) bool {
		return m.score.HomeGoals[i].GetMinute() < m.score.HomeGoals[j].GetMinute()
	})
	return m.score.GetHomeGoals()
}
func (m *Match) GetAwayGoals() []models.Goal {
	sort.Slice(m.score.GetAwayGoals(), func(i, j int) bool {
		return m.score.AwayGoals[i].GetMinute() < m.score.AwayGoals[j].GetMinute()
	})
	return m.score.GetAwayGoals()
}

func (m *Match) GetAwayGoalPos(pos int) models.Goal {
	if m.GetAwayTotalGoals() > pos {
		return m.GetAwayGoals()[pos]
	}
	return models.Goal{}
}
func (m *Match) GetHomeGoalPos(pos int) models.Goal {
	if m.GetHomeTotalGoals() > pos {
		return m.GetHomeGoals()[pos]
	}
	return models.Goal{}
}