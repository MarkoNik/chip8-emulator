package main

import (
	"chip8/model"
	"chip8/utils"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

const pixelSize int = 20

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	if model.DelayTimer > 0 {
		model.DelayTimer--
	}
	if model.SoundTimer > 0 {
		model.SoundTimer--
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// create a pixel
	pixel, _ := ebiten.NewImage(pixelSize, pixelSize, ebiten.FilterDefault)
	_ = pixel.Fill(color.White)

	// reset the screen
	_ = screen.Fill(color.Black)

	for i := 0; i < 32; i++ {
		for j := 0; j < 64; j++ {
			if model.Display[i][j] == true {
				// needed to set pixel coordinates
				options := ebiten.DrawImageOptions{
					GeoM:          ebiten.GeoM{},
					ColorM:        ebiten.ColorM{},
					CompositeMode: 0,
					Filter:        0,
				}

				// set pixel coordinates
				options.GeoM.Translate(float64(j*pixelSize), float64(i*pixelSize))
				// draw pixel on screen
				err := screen.DrawImage(pixel, &options)
				utils.Assert(err)
			}
		}
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64 * pixelSize, 32 * pixelSize
}

func initEbiten() {

	game := &Game{}
	ebiten.SetWindowSize(64*pixelSize, 32*pixelSize)
	ebiten.SetWindowTitle("Chip 8")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
