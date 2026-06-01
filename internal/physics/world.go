package physics

import "github.com/ArteShow/ASPE/internal/math"

type World struct {
	Bodies  []*Body
	Gravity math.Vector2
	Input   math.Vector2
}

func NewWorld(gravity math.Vector2) World {
	return World{
		Gravity: gravity,
		Bodies:  []*Body{},
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

	}

	w.Input = math.Vector2{}
}
