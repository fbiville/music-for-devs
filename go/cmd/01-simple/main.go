package main

import (
	"flag"
	. "github.com/fbiville/soundbox/pkg/soundbox"
	. "github.com/fbiville/soundbox/pkg/units"
	. "time"
)

// go run ./cmd/01-simple -duration 5 -frequency 550
func main() {
	var rawDuration int
	var rawFrequency float64
	flag.IntVar(&rawDuration, "duration", 3, "Duration in seconds")
	flag.Float64Var(&rawFrequency, "frequency", 440, "Sound frequency")
	flag.Parse()

	duration := Duration(rawDuration) * Second
	frequency := rawFrequency * Hertz

	player := NewDefaultPlayer()
	sampler := NewDefaultSampler()
	player.PlayF32LE(sampler.GenerateF32LE(NewSound(duration, Sinusoidal, frequency)))
}
