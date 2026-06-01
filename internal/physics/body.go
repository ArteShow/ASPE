package physics

import "github.com/ArteShow/ASPE/internal/math"

type Body struct {
	Position math.Vector2
	Velocity math.Vector2

	Height float64
	Width  float64

	Mass float64

	Acceleration math.Vector2
	Speed        float64
}

func NewBody(height, width, speed, mass float64) *Body {
	return &Body{
		Position: math.Vector2{},
		Velocity: math.Vector2{},
		Height:   height,
		Width:    width,
		Mass:     mass,
		Speed:    speed,
	}
}
