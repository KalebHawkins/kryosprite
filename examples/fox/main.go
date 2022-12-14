//go:build example
// +build example

package main

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	_ "image/png"

	ks "github.com/KalebHawkins/kryosprite"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

//go:embed fox_sprite_sheet.png
var foxFile []byte

const (
	scrWidth  = 640
	scrHeight = 480
)

type Game struct {
	background *ebiten.Image
	plyr       *ks.Sprite
	notMovedIn float64
	jumpedTick float64
	jumping    bool
}

func (g *Game) Update() error {
	g.plyr.Update()
	return nil
}

func (g *Game) Draw(scr *ebiten.Image) {
	g.background.Fill(color.Black)
	g.plyr.Draw(g.background)

	HUD := fmt.Sprintf("Move Up: W\nMove Left: A\nMove Down: S\nMove Right: D\nJump: Space\nIdle Animation Change in: %.2fs\n", (300/ebiten.ActualTPS())-(g.notMovedIn/ebiten.ActualTPS()))
	ebitenutil.DebugPrint(g.background, HUD)
	ebitenutil.DrawLine(g.background, 0, scrHeight/2, scrWidth, scrHeight/2, color.White)
	ebitenutil.DrawLine(g.background, scrWidth/2, 0, scrWidth/2, scrHeight, color.White)
	scr.DrawImage(g.background, nil)
}

func (g *Game) Layout(outWidth, outHeight int) (int, int) {
	return scrWidth, scrHeight
}

func NewGame() *Game {
	foxImg, _, err := image.Decode(bytes.NewReader(foxFile))
	if err != nil {
		panic(err)
	}

	g := &Game{
		background: ebiten.NewImage(scrWidth, scrHeight),
		plyr: &ks.Sprite{
			Texture:  ebiten.NewImageFromImage(foxImg),
			Position: ks.Vector{X: scrWidth / 2, Y: scrHeight / 2},
			Animator: ks.NewAnimator(),
			Origin:   ks.Center,
		},
	}

	g.plyr.UpdateFunc = plyrUpdate(g)
	g.plyr.SetScale(4, 4)

	g.plyr.Animator.Add("idle", &ks.Animation{
		StartFrame: ks.Frame{X: 0, Y: 0, Width: 32, Height: 32},
		FrameCount: 5,
		Delay:      7,
		Direction:  ks.Horizontal,
	})
	g.plyr.Animator.Add("idleLook", &ks.Animation{
		StartFrame: ks.Frame{X: 0, Y: 32, Width: 32, Height: 32},
		FrameCount: 14,
		Delay:      7,
		Direction:  ks.Horizontal,
	})
	g.plyr.Animator.Add("trot", &ks.Animation{
		StartFrame: ks.Frame{X: 0, Y: 64, Width: 32, Height: 32},
		FrameCount: 8,
		Delay:      7,
		Direction:  ks.Horizontal,
	})
	g.plyr.Animator.Add("jump", &ks.Animation{
		StartFrame: ks.Frame{X: 0, Y: 96, Width: 32, Height: 32},
		FrameCount: 11,
		Delay:      5,
		Direction:  ks.Horizontal,
	})

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
		g.plyr.Animator.Play("idle")

		g.notMovedIn += 1

		// Play a different Idle animation is the player hasn't moved in 300 ticks.
		if g.notMovedIn > 300 {
			g.plyr.Animator.Play("idleLook")

			// After the animation reset the animation's frames.
			if g.notMovedIn > 300+g.plyr.Animator.AnimationTicks() {
				g.plyr.Animator.Reset()
				g.notMovedIn = 0
			}
		}

		playerSpeed := 3.0
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			g.plyr.Animator.Play("trot")
			g.plyr.Position.Y -= playerSpeed
			g.notMovedIn = 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyA) {
			g.plyr.Animator.Play("trot")
			g.plyr.Position.X -= playerSpeed
			g.plyr.FlipHorizontal(true)
			g.notMovedIn = 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) {
			g.plyr.Animator.Play("trot")
			g.plyr.Position.Y += playerSpeed
			g.notMovedIn = 0
		}
		if ebiten.IsKeyPressed(ebiten.KeyD) {
			g.plyr.Animator.Play("trot")
			g.plyr.Position.X += playerSpeed
			g.plyr.FlipHorizontal(false)
			g.notMovedIn = 0
		}
		if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
			g.jumping = true
			g.notMovedIn = 0
			g.jumpedTick = 0
		}

		if g.jumping {
			g.plyr.Animator.Play("jump")
			g.jumpedTick++

			if g.jumpedTick > g.plyr.Animator.AnimationTicks() {
				g.plyr.Animator.Reset()
				g.jumping = false
			}
		}

		return nil
	}
}
