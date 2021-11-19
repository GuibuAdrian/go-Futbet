package models

type Goal struct {
	Minute int
	Player *Player
}

func (goal *Goal) GetMinute() int { return goal.Minute }
func (goal *Goal) GetPlayer() *Player { return goal.Player }