package model

import (
	"fmt"
	"math"
)

type Stats struct {
	min float64
	max float64
	acc float64
}

func NewStats() Stats {
	return Stats{
		min: math.MaxFloat64,
		max: -math.MaxFloat64,
	}
}

func (s *Stats) Update(value float64) {
	if s.min > value {
		s.min = value
	}
	if s.max < value {
		s.max = value
	}
	s.acc += value
}

func (s *Stats) Print(name string, count uint) {
	fmt.Printf("Max %s: %.2f\n", name, s.max)
	fmt.Printf("Min %s: %.2f\n", name, s.min)
	fmt.Printf("Average %s: %.2f\n\n", name, s.acc/float64(count))
}
