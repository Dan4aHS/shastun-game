package random

import (
	"math/rand"
)

func ThrowCubes() (int, int) {
	return rand.Intn(6) + 1, rand.Intn(6) + 1
}
