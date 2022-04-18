package waves

import (
	"fmt"
	"github.com/fbiville/soundbox/pkg/units"
	. "math"
)

type BasicShape int

const (
	Sinusoidal BasicShape = iota
	Square
	Triangle
	Sawtooth
)

type Wave interface {
	Value(rate, sample int) float64
}

type Signal struct {
	Frequency units.Frequency
	Phase     units.Radians
	Amplitude float64
	Shape     BasicShape
}

func SineWaves(frequencies ...units.Frequency) []Wave {
	result := make([]Wave, len(frequencies))
	for i, frequency := range frequencies {
		result[i] = SineWave(frequency)
	}
	return result
}

func SineWave(frequency units.Frequency) Signal {
	return Signal{
		Frequency: frequency,
		Phase:     0,
		Amplitude: 1,
		Shape:     Sinusoidal,
	}
}

func SquareWave(frequency units.Frequency) Signal {
	return Signal{
		Frequency: frequency,
		Phase:     0,
		Amplitude: 1,
		Shape:     Square,
	}
}

func TriangleWave(frequency units.Frequency) Signal {
	return Signal{
		Frequency: frequency,
		Phase:     0,
		Amplitude: 1,
		Shape:     Triangle,
	}
}

func SawtoothWave(frequency units.Frequency) Signal {
	return Signal{
		Frequency: frequency,
		Phase:     0,
		Amplitude: 0.25,
		Shape:     Sawtooth,
	}
}

func AnyWave(
	frequency units.Frequency,
	phase units.Radians,
	amplitude float64,
	shape BasicShape,
) Signal {
	return Signal{
		Frequency: frequency,
		Phase:     phase,
		Amplitude: amplitude,
		Shape:     shape,
	}
}

func (s Signal) Value(rate, sample int) float64 {
	input := float64(sample) / float64(rate)
	sineValue := s.Amplitude * Sin(2*Pi*s.Frequency*input+s.Phase)
	switch s.Shape {
	case Sinusoidal:
		return sineValue
	case Square:
		return Copysign(1, sineValue)
	case Triangle:
		return 2 * float64(s.Amplitude) * Asin(sineValue) / Pi
	case Sawtooth:
		return -2 * float64(s.Amplitude) * Atan(1/(Tan(Pi*s.Frequency*input)))
	}
	panic(fmt.Errorf("unknown shape: %d", s.Shape))
}
