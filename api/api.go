package main

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"chain-traverser/api/handlers"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func pingHandler(c *fasthttp.RequestCtx) {
	c.SetBodyString("pong")
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	r := router.New()

	r.GET("/ping/", pingHandler)
	r.GET("/orb/eth/{address}", handlers.CollectGraphHandler)
	r.GET("/orb/eth/paths/{addressFrom}/to/{addressTo}", handlers.CollectPathHandler)

	log.Info().Msg("Fasthttp server is starting...")

	fasthttp.ListenAndServe(":9000", r.Handler)
}
