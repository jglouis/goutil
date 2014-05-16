package myrand

import (
	"math/rand"
)

//type MyRand embeds type rand form stdlib
type MyRand struct {
	rand.Rand
}

func New(src rand.Source) *MyRand {
	return &MyRand{*rand.New(src)}
}

//return a random int in range [min,max[
func (r *MyRand) RangeInt(min, max int) int {
	return r.Intn(max-min) + min
}
