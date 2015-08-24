package main

import (
	"math"
	"sort"
)

type statistics struct {
	numbers           []float64
	sum               float64
	mean              float64
	median            float64
	standardDeviation float64
	precision         int
}

func NewStatistics(precision int) *statistics {
	return &statistics{
		precision: precision,
	}
}

func (s *statistics) Compute(inputs []float64) {
	s.numbers = inputs
	sort.Float64s(s.numbers)
	s.mean = s.computeMean()
	s.median = s.computeMedian()
	s.standardDeviation = s.computeStandardDeviation()
}

func (s *statistics) computeMean() (total float64) {
	mean := s.computeSum() / float64(len(s.numbers))
	return s.roundToPrecision(mean)
}

func (s *statistics) computeSum() (total float64) {
	for _, x := range s.numbers {
		total += x
	}
	return s.roundToPrecision(total)
}

func (s *statistics) computeMedian() float64 {
	middle := len(s.numbers) / 2
	median := s.numbers[middle]
	if len(s.numbers)%2 == 0 {
		median = (median + s.numbers[middle-1]) / 2
	}
	return s.roundToPrecision(median)
}

func (s *statistics) computeStandardDeviation() float64 {
	mean := s.computeMean()
	var sum float64
	for _, number := range s.numbers {
		sum += math.Pow(number-mean, 2)
	}

	result := math.Sqrt(sum / float64((len(s.numbers) - 1)))
	return s.roundToPrecision(result)
}

func (s *statistics) roundToPrecision(input float64) float64 {
	multiplier := math.Pow(10, float64(s.precision))
	return (float64(int(input * multiplier))) / multiplier
}
