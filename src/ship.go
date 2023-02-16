package main

import (
    _ "golang.org/x/image/bmp" // 这里是找不到这种解压图片的格式了
)

import (
	// "fmt"
	"log"
	// "image/color"						
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)
type Ship struct {
	image  *ebiten.Image
	width  int
	height int
}

func NewShip() *Ship {
	img, _, err := ebitenutil.NewImageFromFile("../images/ship.bmp")
	if err != nil {
		log.Fatal(err)
	}
	width, height := img.Size()
	ship := &Ship{
		image:  img,
		width:  width,
		height: height,
	}
	return ship
}

// 画飞船
func (ship *Ship) Draw(screen *ebiten.Image, cfg *Config) {
    op := &ebiten.DrawImageOptions{}
    op.GeoM.Translate(float64(cfg.ScreenWidth-ship.width)/2, float64(cfg.ScreenHeight-ship.height))
    screen.DrawImage(ship.image, op)
}
