package dao

import (
	"errors"
	"github.com/GuibuAdrian/goFutbet/models"
)
type PlayerDao struct {
	playerSlice []models.Player
}

var playerDaoInstance *PlayerDao

func PlayerDaoGetInstance() *PlayerDao {
	if playerDaoInstance == nil {
		playerDaoInstance = &PlayerDao{}

		initializePlayerDao()
	}

	return playerDaoInstance
}

func (playerDao *PlayerDao) Create(player models.Player)  {
	playerDao.playerSlice = append(playerDao.playerSlice, player)
}

func (playerDao PlayerDao) Read(id int) (models.Player, error) {
	player := models.Player{}
	for _, player := range playerDao.playerSlice {
		if player.GetId() == id {
			return player, nil
		}
	}

	return player, errors.New("player not found")
}

func (playerDao *PlayerDao) Update(player models.Player)  {
	//todo
}

func (playerDao *PlayerDao) Delete(player models.Player)  {
	for posV, playerVal := range playerDao.playerSlice {
		if playerVal == player{
			playerDao.playerSlice = append(playerDao.playerSlice[:posV], playerDao.playerSlice[posV+1:]...)
			break
		}
	}
}

func (playerDao PlayerDao) GetPlayerSlice() []models.Player{
	return playerDao.playerSlice
}

func initializePlayerDao() {
	river, _ := TeamDaoGetInstance().Read(1)
	newells, _ := TeamDaoGetInstance().Read(2)
	PlayerDaoGetInstance().Create(models.InitPlayer("Armani", 1, 1111221, "arquero", river ))
	PlayerDaoGetInstance().Create(models.InitPlayer("Angileri", 3, 3322333, "defensor", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("P. Diaz", 17, 1717171, "defensor", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Rojas", 2, 2122222, "defensor", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Vigo", 16, 1616162, "defensor", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Carrascal", 10, 1010101, "medio C", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("E. Perez", 24, 2424242, "medio C", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Zuculini", 5, 5555555, "medio C", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("De la Cruz", 11, 1121111, "medio C", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Alvarez", 9, 9999999, "delantero", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Romero", 19, 1919323, "delantero", river))
	PlayerDaoGetInstance().Create(models.InitPlayer("Aguerre", 1, 1111111, "arquero", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Compagnucci", 27, 2727272, "defensor", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Lema", 2, 2222222, "defensor", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Z. F. Mansilla", 19, 1919191, "defensor", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Bittolo", 28, 2828282, "defensor", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("P. Perez", 8, 8888888, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("J. Fernandez", 20, 2020202, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Comba", 33, 3333333, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Castro", 36, 3636363, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Sordo", 26, 2626262, "medio C", newells))
	PlayerDaoGetInstance().Create(models.InitPlayer("Scocco", 32, 3232323, "delantero", newells))

}
