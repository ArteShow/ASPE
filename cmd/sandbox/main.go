package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	gravity      = 0.2
	acceleration = 0.1
	velocity     = 0.5
)

type Game struct {
	Object *ebiten.Image

	x  float64
	y  float64
	vx float64
	vy float64
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(g.x, g.y)

	screen.DrawImage(g.Object, op)
}

func (g *Game) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		g.vx = g.vx + velocity
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		g.vx = g.vx - velocity
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		g.vy = g.vy + velocity
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		g.vy = g.vy - velocity
	}

	g.vy = g.vy + gravity

	g.y = g.y + g.vy
	g.x = g.x + g.vx

	if g.vx < 0 {
		g.vx = g.vx + acceleration
	}
	if g.vx > 0 {
		g.vx = g.vx - acceleration
	}

	if g.vy != 0 {
		g.vy = g.vy - acceleration
	}

	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

func main() {
	ebiten.SetWindowTitle("ASPE")
	ebiten.SetWindowSize(800, 600)

	game := &Game{}
	game.Object = ebiten.NewImage(10, 10)
	game.Object.Fill(color.RGBA{255, 0, 0, 255})

	game.vx = 0
	game.vx = 0
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
