package main

import (
	"github.com/ger9000/memes-as-a-service/cmd"
	"github.com/ger9000/memes-as-a-service/internal/shared/logs"
	"github.com/rs/zerolog/log"
)

func main() {
	logs.Init()
	app, err := cmd.InitApp()
	if err != nil {
		log.Fatal().Err(err).Msgf("error in init app: %v", err)
	}

	app.Start()
}
