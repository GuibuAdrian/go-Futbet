package models

type Gambler struct {
	name string
	id int
}

func InitGambler(name string, id int) *Gambler {
	return &Gambler{
		name: name,
		id: id,
	}
}

func (g *Gambler) GetName() string { return g.name }
func (g *Gambler) GetId() int { return g.id }
