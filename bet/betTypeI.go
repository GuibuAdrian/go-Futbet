package interfaces

import (
	"github.com/GuibuAdrian/go-Futbet/bet"
)

type BetTypeI interface {
	Win(bet bet.Bet) bool
}