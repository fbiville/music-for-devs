package main

import (
	"flag"
	. "github.com/fbiville/soundbox/pkg/soundbox"
	. "github.com/fbiville/soundbox/pkg/soundbox/waves"
	. "github.com/fbiville/soundbox/pkg/units"
	. "math"
	. "time"
)

// go run ./cmd/02-phase
func main() {
	var rawDuration int
	var rawFrequency float64
	flag.IntVar(&rawDuration, "duration", 3, "Duration in seconds")
	flag.Float64Var(&rawFrequency, "frequency", 440, "Sound frequency")
	flag.Parse()

	duration := Duration(rawDuration) * Second
	frequency := rawFrequency * Hertz

	player := NewPlayerWithViz()
	sampler := NewDefaultSampler()

	// play with the phase - what happens when it's Pi?
	phase := (Pi / 2) * Rad
	player.PlayF32LE(sampler.GenerateF32LE(
		NewSound(duration,
			SineWave(frequency),
			AnyWave(frequency, phase, 2, Sinusoidal))),
	)
}
