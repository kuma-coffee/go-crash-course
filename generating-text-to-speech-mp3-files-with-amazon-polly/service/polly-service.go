package service

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/polly"
)

type PollyService interface {
	Synthesize(text string, fileName string) error
}

type pollyConfig struct {
	voice string
}

const (
	AUDIO_FORMAT   = "mp3"
	KIMBERLY_VOICE = "Kimberly"
	JOEY_VOICE     = "Joey"
)

func NewKimberlyPollyService() PollyService {
	return &pollyConfig{
		voice: KIMBERLY_VOICE,
	}
}

func NewJoeyPollyService() PollyService {
	return &pollyConfig{
		voice: JOEY_VOICE,
	}
}

func createPollyClient() *polly.Polly {
	session := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return polly.New(session)
}

func (config *pollyConfig) Synthesize(text string, fileName string) error {
	pollyClient := createPollyClient()

	input := &polly.SynthesizeSpeechInput{
		OutputFormat: aws.String(AUDIO_FORMAT),
		Text:         aws.String(text),
		VoiceId:      aws.String(config.voice),
	}

	output, err := pollyClient.SynthesizeSpeech(input)
	if err != nil {
		return err
	}

	outFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer outFile.Close()

	_, err = io.Copy(outFile, output.AudioStream)
	if err != nil {
		return err
	}

	return nil
}
