package main

import (
	// "fmt"
	// "log"
	// "image/color"						
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	input *Input
	ship *Ship
	cfg *Config
}

func NewGame() *Game {
	cfg := loadConfig()
	ebiten.SetWindowSize(cfg.ScreenWidth, cfg.ScreenHeight)
	ebiten.SetWindowTitle(cfg.Title)
	// return &Game{
	// 	input: &Input{},
	// 	cfg:   cfg,
	// }
	// 相同的代码省略...
	return &Game {
		input:   &Input{},
		ship:  NewShip(),
		cfg:  cfg,
	}
}
