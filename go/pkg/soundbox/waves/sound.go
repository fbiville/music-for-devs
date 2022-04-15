package waves

import (
	"time"
)

type Sound interface {
	Value(sampleRate, sampleIndex int) float64
	DurationInSeconds() int
}

type sound struct {
	waves    []Wave
	duration time.Duration
}

func NewSound(duration time.Duration, waves ...Wave) Sound {
	return newSound(duration, waves)
}

func (c *sound) Value(rate, sample int) float64 {
	sum := 0.0
	for _, wave := range c.waves {
		sum += wave.Value(rate, sample)
	}
	return sum
}

func (c *sound) DurationInSeconds() int {
	return int(c.duration.Seconds())
}

func newSound(duration time.Duration, waves []Wave) *sound {
	return &sound{
		waves:    waves,
		duration: duration,
	}
}
