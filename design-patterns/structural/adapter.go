package structural

import "fmt"

// Target interface
type MediaPlayer interface {
	Play(filename string)
}

// Adaptee (incompatible interface)
type VLC struct{}

func (v *VLC) PlayVLC(filename string) {
	fmt.Println("Playing VLC file:", filename)
}

// Adapter
type VLCAdapter struct {
	vlc *VLC
}

func (a *VLCAdapter) Play(filename string) {
	a.vlc.PlayVLC(filename)
}

func main() {
	vlcPlayer := &VLC{}
	player := &VLCAdapter{vlc: vlcPlayer}
	player.Play("movie.vlc")
}
