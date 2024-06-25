package main

import (
	"encoding/json"
	"flag"
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	// "github.com/hajimehoshi/ebiten/v2/inpututil"
	// "github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	// "image/color"
	_ "image/png"
	"log"
	"os"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

var spriteDataPath string

type Sprite struct {
	FilePath string `json:"file_path"`
	X        int    `json:"x"`
	Y        int    `json:"y"`
}

type Game struct {
	spriteImage *ebiten.Image
	spriteX     float64
	spriteY     float64
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(g.spriteX, g.spriteY)
	screen.DrawImage(g.spriteImage, opts)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func loadSpriteData(path string) (Sprite, error) {
	file, err := os.Open(path)
	if err != nil {
		return Sprite{}, err
	}
	defer file.Close()

	var sprite Sprite
	err = json.NewDecoder(file).Decode(&sprite)
	if err != nil {
		return Sprite{}, err
	}

	return sprite, nil
}

func main() {
	flag.StringVar(&spriteDataPath, "sprite", "sprite.json", "path to sprite data JSON file")
	flag.Parse()

	sprite, err := loadSpriteData(spriteDataPath)
	if err != nil {
		log.Fatalf("failed to load sprite data: %v", err)
	}

	file, err := os.Open(sprite.FilePath)
	if err != nil {
		log.Fatalf("failed to open sprite image file: %v", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("failed to decode sprite image: %v", err)
	}

	game := &Game{
		spriteImage: ebiten.NewImageFromImage(img),
		spriteX:     float64(sprite.X),
		spriteY:     float64(sprite.Y),
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sprite Renderer")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
