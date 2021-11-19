package interfaces

import (
	"github.com/GuibuAdrian/go-Futbet/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PlayerDaoI interface {
	Create(models.Player)
	Read(int) models.Player
	ReadByPlayerName(playerName string, teamId primitive.ObjectID) models.Player
	Update(models.Player)
	Delete(models.Player)
}