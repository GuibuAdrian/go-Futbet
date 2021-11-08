package models

type Team struct {
	Name string
	Id   int
}

func InitTeam( name string, id int ) Team {
	return Team {
		Name: name,
		Id:   id,
	}
}

func (team Team) GetTeamName() string { return team.Name }
func (team Team) GetTeamId() int { return team.Id }
