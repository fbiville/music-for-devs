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
	var rawPhase float64
	flag.Float64Var(&rawPhase, "phase", Pi/2, "Phase (in radians)")
	flag.Parse()

	frequency := 440 * Hertz

	player := NewPlayerWithViz()
	sampler := NewDefaultSampler()

	// play with the phase - what happens when it's Pi?
	phase := rawPhase * Rad
	player.PlayF32LE(sampler.GenerateF32LE(
		NewSound(3*Second,
			SineWave(frequency),
			AnyWave(frequency, phase, 2, Sinusoidal))),
	)
}
