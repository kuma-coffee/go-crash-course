package main

import "github.com/kuma-coffee/go-crash-course/generating-text-to-speech-mp3-files-with-amazon-polly/service"

var (
	kimberly service.PollyService = service.NewKimberlyPollyService()
	joey     service.PollyService = service.NewJoeyPollyService()
)

func main() {
	err := kimberly.Synthesize("Hi, I am Kimberly. Nice to meet you", "kimberly.mp3")
	if err != nil {
		panic(err)
	}

	err = joey.Synthesize("Hi, I am Joey, how are you?", "joey.mp3")
	if err != nil {
		panic(err)
	}
}
