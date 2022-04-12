package main

import (
	. "github.com/fbiville/soundbox/pkg/soundbox"
	. "github.com/fbiville/soundbox/pkg/units"
	. "time"
)

const (
	d4      = 293.66 * Hertz
	e4      = 329.63 * Hertz
	fsharp4 = 369.99 * Hertz
	g4      = 392 * Hertz
	a4      = 440 * Hertz
)

// go run ./cmd/03-square
func main() {
	sampler := NewDefaultSampler()
	player := NewPlayerWithViz()
	musicMaestro(sampler, player)
}

func musicMaestro(sampler *Sampler, player *Player) {
	player.PlayF32LE(sampler.GenerateF32LE(
		NewSound(1*Second, Square, d4),
		NewSound(1*Second, Square, g4),
		NewSound(1*Second, Square, fsharp4),
		NewSound(1*Second, Square, g4),
		NewSound(1*Second, Square, a4),
		NewSound(1*Second, Square, e4),
		NewSound(2*Second, Square, a4),
	))
}
