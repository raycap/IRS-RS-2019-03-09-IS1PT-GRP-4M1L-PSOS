package ga

import (
	"math/rand"
)

type GeneticsAlgo interface {
	Solve() Chromosome
}

type Chromosome interface {
	CalculateScore() float64
}

type ChromeNoOp struct {
	value float64
}

func (p *ChromeNoOp) CalculateScore() float64 {
	return p.value
}

func NewPlan(val float64) Chromosome {
	return &ChromeNoOp{
		value: val,
	}
}

type ChromosomeModeller interface {
	CalculateFitness(chromosome Chromosome) float64
	GenerateRandom() Chromosome
	Breed(firstParent, secondParend Chromosome) Chromosome
	Mutate(chromosome Chromosome) Chromosome
}

type ChromosomeModellerNoOp struct {
}

func (m *ChromosomeModellerNoOp) CalculateFitness(chromosome Chromosome) float64 {
	p := chromosome.(*ChromeNoOp)
	return p.CalculateScore()
}

func (m *ChromosomeModellerNoOp) GenerateRandom() Chromosome {
	return NewPlan(rand.Float64() * 100.0)
}

func (m *ChromosomeModellerNoOp) Breed(firstParent, secondParend Chromosome) Chromosome {
	return NewPlan(firstParent.CalculateScore()*0.5 + secondParend.CalculateScore()*0.5)
}

func (m *ChromosomeModellerNoOp) Mutate(chromosome Chromosome) Chromosome {
	return m.GenerateRandom()
}
