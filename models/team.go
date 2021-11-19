package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Team struct {
	Name string
	objId primitive.ObjectID
}

func InitTeam( name string ) Team {
	return Team {
		Name: name,
	}
}

func (team Team) GetTeamName() string { return team.Name }
func (team Team) GetTeamObjId() primitive.ObjectID { return team.objId }

func (team *Team) SetTeamObjId(objId primitive.ObjectID) {
	team.objId = objId
}
