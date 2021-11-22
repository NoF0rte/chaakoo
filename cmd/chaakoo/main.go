package main

import (
	"github.com/NoF0rte/chaakoo/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatal().Err(err).Msg("cannot start chaakoo")
	}
}
