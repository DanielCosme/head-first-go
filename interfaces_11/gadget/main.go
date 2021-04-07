package main

import (
	"./tape"
)

type Player interface {
	Play(string)
	Stop()
}

func playList(device Player, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()

	// what a value IS
	// what a value can DO

	// Type assertion.
	// Revert back to the original type instead of the interface.
	if recorder, ok := device.(tape.TapeRecorder); ok {
		recorder.Record()
	}
}

func main() {
	var player Player = tape.TapePlayer{Batteries: "AAA"}
	songs := []string{"camisa negra", "bohemian rapsody", "waka waka"}

	playList(player, songs)
	player = tape.TapeRecorder{}
	playList(player, songs)
}
