package dice

import (
	"math/rand"
	"time"
)

func Roll(sides int) int {
	return rand.Intn(sides) + 1
}

// Library initialization, runs before than main.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// That function should be better, and be the client who calls it.
func Seed(n int64) {
	if n == 0 {
		rand.Seed(time.Now().UnixNano())
	} else {
		rand.Seed(n)
	}
}
