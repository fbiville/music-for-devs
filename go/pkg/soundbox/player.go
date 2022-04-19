package soundbox

import (
	"fmt"
	"os/exec"
	"time"
)

type Player struct {
	sampleRate    int
	visualization bool
}

func NewDefaultPlayer() *Player {
	return &Player{sampleRate: defaultSampleRate}
}

func NewPlayerWithViz() *Player {
	return &Player{sampleRate: defaultSampleRate, visualization: true}
}

func (p *Player) PlayF32LE(file string, duration time.Duration) {
	visualizationFlags := p.visualizationFlags()
	flags := append([]string{
		"-f", "f32le", // f32le: PCM 32-bit floating-point little-endian
		"-t", fmt.Sprintf("%d", int(duration.Seconds())), // f32le: PCM 32-bit floating-point little-endian
		"-ar", fmt.Sprintf("%d", p.sampleRate*2), // FIXME: not sure why factor fixes the weird pitch division
		"-autoexit",
	}, visualizationFlags...)
	flags = append(flags, file)
	command := exec.Command("ffplay", flags...)
	if err := command.Run(); err != nil {
		panic(err)
	}
}

func (p *Player) visualizationFlags() []string {
	if !p.visualization {
		return []string{"-nodisp"}
	}
	return []string{"-showmode", "1"}
}
