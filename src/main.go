package main

// a small project implemented for some fun .......
// reference: https://darjun.github.io/2022/11/15/godailylib/ebiten1/

import (
	// "fmt"
	"log"
	// "image/color"						
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Input struct {
	msg string
}
func (i *Input) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		// fmt.Println("←←←←←←←←←←←←←←←←←←←←←←←")
		i.msg = "left pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		// fmt.Println("→→→→→→→→→→→→→→→→→→→→→→→")
		i.msg = "right pressed"
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// fmt.Println("-----------------------")
		i.msg = "space pressed"
	}
}
func (g *Game) Update() error {
	g.input.Update()
	return nil
}

// 改版后的船画不出来,我觉得是我下载的图片来源不对,先放一放.....
func (g *Game) Draw(screen *ebiten.Image) {
    screen.Fill(g.cfg.BgColor)
    g.ship.Draw(screen, g.cfg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.cfg.ScreenWidth / 2, g.cfg.ScreenHeight / 2
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("外星人入侵")
	game := NewGame()
	// if err := ebiten.RunGame(&Game{}); err != nil {
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}