package soundbox

import (
	"encoding/binary"
	"github.com/fbiville/soundbox/pkg/soundbox/waves"
	"math"
	"os"
	"time"
)

const defaultSampleRate = 44100

type Sampler struct {
	sampleRate int
}

func NewDefaultSampler() *Sampler {
	return &Sampler{sampleRate: defaultSampleRate}
}

func (s *Sampler) GenerateF32LE(sounds ...waves.Sound) (string, time.Duration) {
	file, _ := os.CreateTemp("", "sound.raw")
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	totalDuration := time.Duration(0)
	for _, sound := range sounds {
		totalDuration += time.Duration(sound.DurationInSeconds()) * time.Second
		for t := 0; t < sound.DurationInSeconds(); t++ {
			for i := 0; i < s.sampleRate; i++ {
				sample := sound.Value(s.sampleRate, i)
				var buf [8]byte
				binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
				if _, err := file.Write(buf[:]); err != nil {
					panic(err)
				}
			}
		}
	}
	return file.Name(), totalDuration
}
