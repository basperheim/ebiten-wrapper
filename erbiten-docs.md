https://ebitengine.org/en/documents/install.html

Installing Go
Install Go on your machine. Ebitengine requires Go 1.18 or later.

Installing a C compiler
A C compiler is required as Ebitengine uses not only Go but also C.

On the latest macOS, just type clang on your terminal and a dialog would appear if you don't have clang compiler. Follow the instruction to install it.

You might find the following error when executing clang.

xcrun: error: invalid active developer path (/Library/Developer/CommandLineTools), missing xcrun at: /Library/Developer/CommandLineTools/usr/bin/xcrun
In this case, run xcode-select --install and install commandline tools.

## Install Erbiten in directory

```bash
go get github.com/hajimehoshi/ebiten/v2
```

## Confirming your environment
You can check whether you have a correct environment by executing an example:

```bash
go run github.com/hajimehoshi/ebiten/v2/examples/rotate@latest
```

If you see this window with a rotating Gophers image, congratulations! You have a correct environment to use Ebitengine!

Running a program with Ebitengine
Ebitengine can be used as a usual Go library. Go command automatically installs Ebitengine when your program uses Ebitengine.

First, create your local module.

# Create a directory for your game.

```bash
mkdir yourgame
cd yourgame
```

# Initialize go.mod by `go mod init`.
go mod init github.com/yourname/yourgame
Use URL for your module name like github.com/yourname/yourgame. Actually, any module name like example.com/m is fine as long as you don't plan to share this publicly. You can change the module name anytime later.

Add main.go with this content:

```go
package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
```

Run go mod tidy to add dependencies to your go.mod:

```bash
go mod tidy
```

Finally, run go run to execute your program.

```bash
go run .
```

You will be able to see a window with a message: