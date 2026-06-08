package seed

import (
	"math/rand"
	"time"
)

type SeedString string

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789+-"
const prefix = "0"

type SeedGenerator struct {
	rng *rand.Rand
}

func NewSeedGenerator() *SeedGenerator {
	return &SeedGenerator{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (sg *SeedGenerator) Generate() SeedString {
	b := make([]byte, 7)
	for i := range b {
		b[i] = charset[sg.rng.Intn(len(charset))]
	}
	return SeedString(prefix + string(b))
}
