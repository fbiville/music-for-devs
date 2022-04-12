package main

import (
	. "github.com/fbiville/soundbox/pkg/soundbox"
	. "github.com/fbiville/soundbox/pkg/units"
	. "time"
)

const (
	d2      = 73.4 * Hertz
	e2      = 82.41 * Hertz
	g2      = 98 * Hertz
	a2      = 110 * Hertz
	b3      = 246.94 * Hertz
	csharp4 = 277.18 * Hertz
	d4      = 293.66 * Hertz
	e4      = 329.63 * Hertz
	fsharp4 = 369.99 * Hertz
	g4      = 392 * Hertz
	a4      = 440 * Hertz
)

// go run ./cmd/02-naive-chords
func main() {
	sampler := NewDefaultSampler()
	player := NewPlayerWithViz()
	baseMelody(sampler, player)
	Sleep(1 * Second)
	musicMaestro(sampler, player)
	Sleep(1 * Second)
}

func baseMelody(sampler *Sampler, player *Player) {
	player.PlayF32LE(sampler.GenerateF32LE(
		NewSound(1*Second, Sinusoidal, d4),
		NewSound(1*Second, Sinusoidal, g4),
		NewSound(1*Second, Sinusoidal, fsharp4),
		NewSound(1*Second, Sinusoidal, g4),
		NewSound(1*Second, Sinusoidal, a4),
		NewSound(1*Second, Sinusoidal, e4),
		NewSound(2*Second, Sinusoidal, a4),
	))
}

func musicMaestro(sampler *Sampler, player *Player) {
	player.PlayF32LE(sampler.GenerateF32LE(
		NewSound(1*Second, Sinusoidal, d4, b3, g2),
		NewSound(1*Second, Sinusoidal, g4),
		NewSound(1*Second, Sinusoidal, fsharp4, b3, d2),
		NewSound(1*Second, Sinusoidal, g4),
		NewSound(1*Second, Sinusoidal, a4, csharp4, a2),
		NewSound(1*Second, Sinusoidal, e4),
		NewSound(2*Second, Sinusoidal, a4, csharp4, e2),
	))
}
