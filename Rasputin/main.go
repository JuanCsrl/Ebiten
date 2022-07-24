package main

import (
	"bytes"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/examples/resources/images"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240

	frameOX     = 0
	frameOY     = 16
	frameWidth  = 16
	frameHeight = 16
	frameCount  = 4
)

var img *ebiten.Image

func init() {
	var err error
	img, _, err = ebitenutil.NewImageFromFile("cat_idle.png")
	if err != nil {
		log.Fatal(err)
	}
}

type Game struct {
	count int
}

func (g *Game) Update() error {
	g.count++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// os configurações setadas aqui, se sobrepoem
	//screen.DrawImage(img, nil) // renderizando imagem
	//ebitenutil.DebugPrint(screen, "Hello, World!") // print na tela
	screen.Fill(color.RGBA{0xff, 0, 0, 0xff}) // defino a cor da tela
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(-float64(frameWidth)/2, -float64(frameHeight)/2)
	op.GeoM.Translate(screenWidth/2, screenHeight/2)
	i := (g.count / 5) % frameCount
	sx, sy := frameOX+i*frameWidth, frameOY
	screen.DrawImage(img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image), op)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
	//return screenWidth, screenHeight
}

func main() {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		log.Fatal(err)
	}
	img = ebiten.NewImageFromImage(img)
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Rasputin, the cat")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
