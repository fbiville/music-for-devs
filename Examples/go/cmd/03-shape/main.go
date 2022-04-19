package main

import (
	. "github.com/fbiville/soundbox/pkg/soundbox"
	. "github.com/fbiville/soundbox/pkg/soundbox/waves"
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

// go run ./cmd/03-shape
func main() {
	sampler := NewDefaultSampler()
	player := NewPlayerWithViz()
	musicMaestro(sampler, player)
}

type shaper func(frequency Frequency) Signal

func musicMaestro(sampler *Sampler, player *Player) {
	// try different shapes - does it sound different?
	shapers := []shaper{
		SineWave,
		SquareWave,
		TriangleWave,
		SawtoothWave,
	}
	for _, shaper := range shapers {
		player.PlayF32LE(sampler.GenerateF32LE(
			NewSound(1*Second, shaper(d4)),
			NewSound(1*Second, shaper(g4)),
			NewSound(1*Second, shaper(fsharp4)),
			NewSound(1*Second, shaper(g4)),
			NewSound(1*Second, shaper(a4)),
			NewSound(1*Second, shaper(e4)),
			NewSound(2*Second, shaper(a4)),
		))
		Sleep(1 * Second)
	}
}
