package main

import (
	"github.com/gordonklaus/portaudio"
	"log"
)

const sampleRate = 44100
const seconds = 1

func main() {
	err := portaudio.Initialize()
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		err := portaudio.Terminate()
		if err != nil {
			log.Print(err)
		}
	}()
	buffer := make([]float32, sampleRate*seconds)
	stream, err := portaudio.OpenDefaultStream(1, 0, sampleRate, len(buffer), func(in []float32) {
		for i := range buffer {
			buffer[i] = in[i]
		}
	})
	if err != nil {
		panic(err)
	}
	stream.Start()
	defer stream.Close()
}
