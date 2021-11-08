package models

type Score struct {
	HomeGoals, AwayGoals []Goal
}

func (s *Score) AddHomeGoal(goal Goal) *Score {
	s.HomeGoals = append(s.HomeGoals, goal)
	return s
}

func (s *Score) AddAwayGoal(goal Goal) *Score {
	s.AwayGoals = append(s.AwayGoals, goal)
	return s
}

func (s *Score) GetHomeGoals() []Goal { return s.HomeGoals }
func (s *Score) GetAwayGoals() []Goal { return s.AwayGoals }