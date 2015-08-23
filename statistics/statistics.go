package main

import "sort"

type statistics struct {
	numbers []float64
	sum     float64
	mean    float64
	median  float64
}

func NewStatistics() *statistics {
	return &statistics{}
}

func (s *statistics) Compute(inputs []float64) {
	s.numbers = inputs
	sort.Float64s(s.numbers)
	s.mean = s.computeMean()
	s.median = s.computeMedian()
}

func (s *statistics) computeMean() (total float64) {
	return s.computeSum() / float64(len(s.numbers))
}

func (s *statistics) computeSum() (total float64) {
	for _, x := range s.numbers {
		total += x
	}
	return total
}

func (s *statistics) computeMedian() float64 {
	middle := len(s.numbers) / 2
	result := s.numbers[middle]
	if len(s.numbers)%2 == 0 {
		result = (result + s.numbers[middle-1]) / 2
	}
	return result
}
