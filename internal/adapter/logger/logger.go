package logger

import (
	"github.com/rs/zerolog/log"
)

type ErrorLog struct {
	Message string
	Context string
}

func LogCritical(e ErrorLog) {
	log.Error().Str("error-level", "critical").Str("context", e.Context).Msg(e.Message)
}

func LogError(e ErrorLog) {
	log.Error().Str("context", e.Context).Msg(e.Message)
}

func LogInfo(e ErrorLog) {
	log.Info().Str("context", e.Context).Msg(e.Message)
}

func LogDebug(e ErrorLog) {
	log.Debug().Str("context", e.Context).Msg(e.Message)
}
