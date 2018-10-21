package ops

import (
	"github.com/rs/zerolog/log"
)

// Add adds 2 numbers
func Add(a, b int) int {
	log.Debug().Int("a", a).Int("b", b).Msg("Mul")
	return a + b
}
