package physics

import (
	"github.com/ArteShow/ASPE/internal/math"
)

type World struct {
	Bodies  []*Body
	Gravity math.Vector2
	Input   math.Vector2

	Height float64
	Width  float64
}

func NewWorld(width, height float64, gravity math.Vector2) World {
	return World{
		Gravity: gravity,
		Bodies:  []*Body{},
		Width:   width,
		Height:  height,
	}
}

func (w *World) Update() {
	for index := range w.Bodies {
		body := w.Bodies[index]

		Force := math.Vector2{}

		Force.Y += w.Gravity.Y * body.Mass
		Force.X += w.Gravity.X * body.Mass

		Force.X += w.Input.X * body.Speed
		Force.Y += w.Input.Y * body.Speed

		body.Acceleration.X = Force.X / body.Mass
		body.Velocity.X += body.Acceleration.X

		body.Acceleration.Y = Force.Y / body.Mass
		body.Velocity.Y += body.Acceleration.Y

		body.Position.X += body.Velocity.X
		body.Position.Y += body.Velocity.Y

		if body.Position.Y > w.Height-body.Height {
			body.Position.Y = w.Height - body.Height
			body.Velocity.Y = 0
		}
		if body.Position.Y < 0 {
			body.Position.Y = 0
			body.Velocity.Y = 0
		}
		if body.Position.X > w.Width-body.Width {
			body.Position.X = w.Width - body.Width
			body.Velocity.X = 0
		}
		if body.Position.X < 0 {
			body.Position.X = 0
			body.Velocity.X = 0
		}

		for _, collisionBody := range w.Bodies {
			if collisionBody.Position == body.Position {
				return
			}

			separatedByX := body.Position.X+body.Width > collisionBody.Position.X || body.Position.X > collisionBody.Position.X+collisionBody.Width
			separatedByY := body.Position.Y+body.Height > collisionBody.Position.Y || body.Position.Y > collisionBody.Position.Y+collisionBody.Height
			if !separatedByX && !separatedByY {
				body.Velocity.X = 0
			}
		}
	}

	w.Input = math.Vector2{}
}
