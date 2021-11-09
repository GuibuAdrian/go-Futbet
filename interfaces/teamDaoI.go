package interfaces

import "github.com/GuibuAdrian/go-Futbet/models"

type TeamDaoI interface {
	Create(models.Team)
	Read(int) models.Team
	Update(models.Team)
	Delete(models.Team)
}