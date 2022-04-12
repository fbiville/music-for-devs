package soundbox

import (
	"fmt"
	"github.com/fbiville/soundbox/pkg/units"
	"math"
	"time"
)

type BasicShape int

const (
	Sinusoidal BasicShape = iota
	Square
	Triangle
)

type Sound interface {
	Value(sampleRate, sampleIndex int) float64
	DurationInSeconds() int
}

type sound struct {
	frequencies []units.Frequency
	duration    time.Duration
	shaper      func(float64) float64
}

func NewSound(duration time.Duration, shape BasicShape, freqs ...units.Frequency) Sound {
	return newSound(duration, shape, freqs)
}

func (c *sound) Value(sampleRate, sampleIndex int) float64 {
	sum := 0.0
	for _, sound := range c.frequencies {
		sum += c.sineWavePoint(sampleRate, sampleIndex, sound)
	}
	return sum
}

func (c *sound) DurationInSeconds() int {
	return int(c.duration.Seconds())
}

func (c *sound) sineWavePoint(samplingCount int, sampleIndex int, sound units.Frequency) float64 {
	sine := math.Sin(sound * 2 * math.Pi * float64(sampleIndex) / float64(samplingCount))
	return c.shaper(sine)
}

func sineTransformer(shape BasicShape) func(float64) float64 {
	switch shape {
	case Sinusoidal:
		return func(f float64) float64 { return f }
	case Square:
		// amplitude is reduced since the produce wave sounds louder for some reason
		return func(f float64) float64 { return math.Copysign(1, f) / 2 }
	case Triangle:
		return func(f float64) float64 { return 2 * math.Asin(f) / math.Pi }
	}
	panic(fmt.Sprintf("Unsupported shape: %v", shape))
}

func newSound(duration time.Duration, shape BasicShape, frequencies []units.Frequency) *sound {
	return &sound{
		frequencies: frequencies,
		duration:    duration,
		shaper:      sineTransformer(shape),
	}
}
