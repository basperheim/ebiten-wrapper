package main

import (
	"encoding/json"
	"flag"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	_ "image/png"
	"log"
	"os"
)

const (
	screenWidth  = 800
	screenHeight = 600
	frameRate    = 60
)

type Sprite struct {
	FilePath   string      `json:"file_path"`
	X          int         `json:"x"`
	Y          int         `json:"y"`
	Animations []Animation `json:"animations"`
	image      *ebiten.Image
	speed      float64
	direction  int
}

type Animation struct {
	Type   string  `json:"type"`
	StartX int     `json:"start_x"`
	EndX   int     `json:"end_x"`
	Speed  float64 `json:"speed"`
}

type Game struct {
	sprites []Sprite
}

func (g *Game) Update() error {
	for i := range g.sprites {
		sprite := &g.sprites[i]
		for _, animation := range sprite.Animations {
			if animation.Type == "move" {
				sprite.X += sprite.direction * int(animation.Speed)
				if sprite.X >= animation.EndX {
					sprite.direction = -1
				} else if sprite.X <= animation.StartX {
					sprite.direction = 1
				}
			}
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, sprite := range g.sprites {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(sprite.X), float64(sprite.Y))
		screen.DrawImage(sprite.image, opts)

		// Draw bounding box for demonstration
		vector.DrawFilledRect(screen, float32(sprite.X), float32(sprite.Y), float32(sprite.image.Bounds().Dx()), float32(sprite.image.Bounds().Dy()), color.RGBA{0, 255, 0, 128}, false)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func loadSprites(jsonData string) ([]Sprite, error) {
	var spriteData struct {
		Sprites []Sprite `json:"sprites"`
	}
	err := json.Unmarshal([]byte(jsonData), &spriteData)
	if err != nil {
		return nil, err
	}

	for i, sprite := range spriteData.Sprites {
		file, err := os.Open(sprite.FilePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()

		img, _, err := image.Decode(file)
		if err != nil {
			return nil, err
		}
		spriteData.Sprites[i].image = ebiten.NewImageFromImage(img)
		spriteData.Sprites[i].direction = 1
	}

	return spriteData.Sprites, nil
}

func main() {
	var spriteJSON string
	flag.StringVar(&spriteJSON, "sprite", "", "JSON string for sprite data")
	flag.Parse()

	if spriteJSON == "" {
		log.Fatal("sprite JSON data not provided")
	}

	sprites, err := loadSprites(spriteJSON)
	if err != nil {
		log.Fatalf("failed to load sprite data: %v", err)
	}

	game := &Game{sprites: sprites}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Sprite Renderer")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
