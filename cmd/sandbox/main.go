package main

import (
	"image/color"

	"github.com/ArteShow/ASPE/internal/math"
	"github.com/ArteShow/ASPE/internal/physics"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	World = physics.NewWorld(800, 600, math.Vector2{
		Y: 0.1,
	})
)

type Game struct {
	Object    *ebiten.Image
	NewObject *ebiten.Image
	x         float64
	y         float64
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(g.x, g.y)

	screen.DrawImage(g.Object, op)
}

func (g *Game) Update() error {
	input := math.Vector2{}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		input.X += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		input.X -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		input.Y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		input.Y -= 1
	}

	World.Input = input
	World.Update()

	g.x = World.Bodies[0].Position.X
	g.y = World.Bodies[0].Position.Y

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

	game.NewObject = ebiten.NewImage(100, 100)
	game.NewObject.Fill(color.White)

	World.Bodies = append(World.Bodies, physics.NewBody(
		50,
		50,
		1,
		3,
	))
	World.Bodies = append(World.Bodies, physics.NewBody(
		100,
		100,
		1,
		3,
	))

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
