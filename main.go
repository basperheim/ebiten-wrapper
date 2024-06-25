package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	_ "image/png" // Import to decode PNG images
	"log"
	"os"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	ebitenImage *ebiten.Image
	clicked     bool
)

func init() {
	// Open the image file
	file, err := os.Open("assets/forest-tile-128px.png")
	if err != nil {
		log.Fatalf("failed to open image file: %v", err)
	}
	defer file.Close()

	// Decode the image
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
	}
	ebitenImage = ebiten.NewImageFromImage(img)
}

type Game struct{}

// Update proceeds the game state. Update is called every tick (1/60 [s] by default).
func (g *Game) Update() error {
	// Check if the left mouse button was pressed
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// Get the mouse position
		x, y := ebiten.CursorPosition()

		// Check if the mouse click was within the image bounds
		if x >= 0 && x < screenWidth && y >= 0 && y < screenHeight {
			clicked = true
		}
	}

	// Reset clicked state if needed
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		clicked = false
	}

	return nil
}

// Draw draws the game screen. Draw is called every frame (1/60 [s] by default).
func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the image at the center of the screen
	opts := &ebiten.DrawImageOptions{}
	w, h := ebitenImage.Bounds().Dx(), ebitenImage.Bounds().Dy()
	x := (screenWidth - w) / 2
	y := (screenHeight - h) / 2
	opts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(ebitenImage, opts)

	// Draw a rectangle around the image if it was clicked
	if clicked {
		var green color.Color = color.RGBA{0, 255, 0, 255}
		vector.DrawFilledRect(screen, float32(x), float32(y), float32(w), float32(h), green, true)
		// vector.DrawFilledRect(screen, float32(b.x), float32(b.y), float32(b.width), float32(b.height), green, true)
		// ebitenutil.DrawFilledRect(screen, float64(x), float64(y), float64(w), float64(h), color.RGBA{255, 0, 0, 255})
	}
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Simple Ebiten Example")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
