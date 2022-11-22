//go:build example
// +build example

package main

import (
	"image/color"

	ks "github.com/KalebHawkins/kryosprite"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	scrWidth  = 640
	scrHeight = 480
)

type Game struct {
	background *ebiten.Image
	plyr       *ks.Sprite
}

func (g *Game) Update() error {
	g.plyr.Update()

	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {
	g.background.Fill(color.Black)
	g.plyr.Draw(g.background)

	ebitenutil.DrawLine(g.background, 0, scrHeight/2, scrWidth, scrHeight/2, color.White)
	ebitenutil.DrawLine(g.background, scrWidth/2, 0, scrWidth/2, scrHeight, color.White)
	scr.DrawImage(g.background, nil)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return scrWidth, scrHeight
}

func NewGame() *Game {
	g := &Game{
		background: ebiten.NewImage(scrWidth, scrHeight),
		plyr: &ks.Sprite{
			Texture:  ebiten.NewImage(32, 32),
			Position: &ks.Vector{X: scrWidth / 2, Y: scrHeight / 2},
			Animator: nil,
			Origin:   ks.Center,
			Color:    color.White,
		},
	}
	g.plyr.UpdateFunc = plyrUpdate(g)

	return g
}

func main() {
	ebiten.SetWindowSize(scrWidth, scrHeight)
	ebiten.SetWindowTitle("Basic Demo")

	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}

func plyrUpdate(g *Game) func() error {
	return func() error {
		playerSpeed := 5.0

		if ebiten.IsKeyPressed(ebiten.KeyW) {
			g.plyr.Position.Y -= playerSpeed
		}
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			g.plyr.Position.X -= playerSpeed
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			g.plyr.Position.Y += playerSpeed
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			g.plyr.Position.X += playerSpeed

		}

		return nil
	}
}
