package main


import (
	// "fmt"
	"log"
	"os"
	"encoding/json"
	"image/color"						
	// "github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Config struct {
    ScreenWidth  int        `json:"screenWidth"`
    ScreenHeight int        `json:"screenHeight"`
    Title        string     `json:"title"`
    BgColor      color.RGBA `json:"bgColor"`
}

func loadConfig() *Config {
    f, err := os.Open("./config.json")
    if err != nil {
        log.Fatalf("os.Open failed: %v\n", err)
    }

    var cfg Config
    err = json.NewDecoder(f).Decode(&cfg)
    if err != nil {
		log.Fatalf("json.Decode failed: %v\n", err)
	}

    return &cfg
}