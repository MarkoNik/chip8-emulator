package IO

import (
	"chip8/model"
	"chip8/utils"
	"github.com/hajimehoshi/ebiten"
	"image/color"
	"log"
)

const pixelSize int = 20

// create a pixel
var pixel, _ = ebiten.NewImage(pixelSize, pixelSize, ebiten.FilterDefault)

// create main panel
var panel, _ = ebiten.NewImage(64*pixelSize, 32*pixelSize, ebiten.FilterDefault)

type Game struct{}

func (g *Game) Update(screen *ebiten.Image) error {
	model.Run()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// decrement timers
	if model.DelayTimer > 0 {
		model.DelayTimer--
	}
	if model.SoundTimer > 0 {
		model.SoundTimer--
	}

	// performance optimization
	if !model.DisplayChanged {
		err := screen.DrawImage(panel, nil)
		utils.Assert(err)
		return
	}
	model.DisplayChanged = false

	// reset the screen
	_ = panel.Clear()
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
				err := panel.DrawImage(pixel, &options)
				utils.Assert(err)
			}
		}
	}
	err := screen.DrawImage(panel, nil)
	utils.Assert(err)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 64 * pixelSize, 32 * pixelSize
}

func InitEbiten() {

	game := &Game{}
	ebiten.SetWindowSize(64*pixelSize, 32*pixelSize)
	ebiten.SetWindowTitle("Chip 8")
	ebiten.SetMaxTPS(1000)
	_ = pixel.Fill(color.White)

	initAudio()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}

func initAudio() {

}
