package interfaces

import "github.com/GuibuAdrian/goFutbet/models"

type PlayerDaoI interface {
	Create(models.Player)
	Read(int) models.Player
	Update(models.Player)
	Delete(models.Player)
}