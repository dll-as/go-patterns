package main

import "fmt"

// State interface
type State interface {
	Play()
	Pause()
	Stop()
}

// Context
type MediaPlayer struct {
	state State
}

func (m *MediaPlayer) SetState(state State) {
	m.state = state
}

func (m *MediaPlayer) Play()  { m.state.Play() }
func (m *MediaPlayer) Pause() { m.state.Pause() }
func (m *MediaPlayer) Stop()  { m.state.Stop() }

// Concrete States
type PlayingState struct {
	player *MediaPlayer
}

func (p *PlayingState) Play() {
	fmt.Println("Already playing")
}

func (p *PlayingState) Pause() {
	fmt.Println("Pausing playback")
	p.player.SetState(&PausedState{player: p.player})
}

func (p *PlayingState) Stop() {
	fmt.Println("Stopping playback")
	p.player.SetState(&StoppedState{player: p.player})
}

type PausedState struct {
	player *MediaPlayer
}

func (p *PausedState) Play() {
	fmt.Println("Resuming playback")
	p.player.SetState(&PlayingState{player: p.player})
}

func (p *PausedState) Pause() {
	fmt.Println("Already paused")
}

func (p *PausedState) Stop() {
	fmt.Println("Stopping playback")
	p.player.SetState(&StoppedState{player: p.player})
}

type StoppedState struct {
	player *MediaPlayer
}

func (s *StoppedState) Play() {
	fmt.Println("Starting playback")
	s.player.SetState(&PlayingState{player: s.player})
}

func (s *StoppedState) Pause() {
	fmt.Println("Cannot pause, already stopped")
}

func (s *StoppedState) Stop() {
	fmt.Println("Already stopped")
}

func main() {
	player := &MediaPlayer{}
	player.SetState(&StoppedState{player: player})

	player.Play()  // Starting playback
	player.Pause() // Pausing playback
	player.Play()  // Resuming playback
	player.Stop()  // Stopping playback
	player.Stop()  // Already stopped
}
