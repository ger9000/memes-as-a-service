package logs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	logger := zerolog.New(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.TimeOnly,
		FormatFieldValue: func(i interface{}) string {
			return strings.ReplaceAll(fmt.Sprintf("\n%s", i), "\\n", "\n")
		}},
	).Level(zerolog.DebugLevel).With().Timestamp().Caller()
	log.Logger = logger.Logger()
}
