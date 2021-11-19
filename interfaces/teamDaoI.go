package interfaces

import (
	"github.com/GuibuAdrian/go-Futbet/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TeamDaoI interface {
	Create(models.Team)
	Read(primitive.ObjectID) models.Team
	ReadByName(string) models.Team
	Update(models.Team)
	Delete(models.Team)
}