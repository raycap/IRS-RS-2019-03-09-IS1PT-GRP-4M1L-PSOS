package ga

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type geneticsAlgoImpl struct {
	populationSize, eliteSize, generations int64
	mutationRate                           float64
	chromModel                             ChromosomeModeller
	useConcurrency                         bool
	scoreProgress                          []float64
}

type scoreWrapper struct {
	fitnessScore float64
	chromosome   Chromosome
}

func New(popSize, eliteSize, generations int64, mutationRate float64, chromModel ChromosomeModeller, useCon bool) GeneticsAlgo {
	return &geneticsAlgoImpl{
		populationSize: popSize,
		eliteSize:      eliteSize,
		generations:    generations,
		mutationRate:   mutationRate,
		chromModel:     chromModel,
		useConcurrency: useCon,
		scoreProgress:  []float64{},
	}
}

func (ga *geneticsAlgoImpl) Solve() Chromosome {
	rand.Seed(time.Now().UnixNano())
	var (
		stuck             bool
		populations       []Chromosome
		rankedPopulations []scoreWrapper
	)
	if ga.useConcurrency {
		populations = ga.initPopulationWithAsync()
	} else {
		populations = ga.initPopulation()
	}
	for i := int64(0); i <= ga.generations; i++ {
		populations, stuck = ga.createNextGen(populations)
		if stuck {
			break
		}
	}
	if ga.useConcurrency {
		rankedPopulations = ga.rankByConcurreny(populations)
	} else {
		rankedPopulations = ga.rank(populations)
	}
	return rankedPopulations[0].chromosome
}

func (ga *geneticsAlgoImpl) initPopulationWithAsync() []Chromosome {
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	populations := make([]Chromosome, ga.populationSize)
	for i := int64(0); i < ga.populationSize; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			mutex.Lock()
			populations[i] = ga.chromModel.GenerateRandom()
			mutex.Unlock()
		}(i)
	}
	wg.Wait()

	return populations
}

func (ga *geneticsAlgoImpl) initPopulation() []Chromosome {
	populations := make([]Chromosome, ga.populationSize)
	for i := int64(0); i < ga.populationSize; i++ {
		populations[i] = ga.chromModel.GenerateRandom()
	}
	return populations
}

func (ga *geneticsAlgoImpl) createNextGen(populations []Chromosome) ([]Chromosome, bool) {
	var rankedPopulation []scoreWrapper
	if ga.useConcurrency {
		rankedPopulation = ga.rankByConcurreny(populations)
	} else {
		rankedPopulation = ga.rank(populations)
	}
	ga.recordProgress(rankedPopulation[0].fitnessScore)
	if ga.progressStuck() {
		return populations, true
	}
	matingPool := ga.generateMatingPool(rankedPopulation)
	children := ga.breedPopulation(matingPool)
	mutatedChildren := ga.mutate(children)
	return mutatedChildren, false
}

func (ga *geneticsAlgoImpl) progressStuck() bool {
	if len(ga.scoreProgress) <= 2 {
		return false
	}

	count := 0
	for i := len(ga.scoreProgress) - 1; i >= 1; i-- {
		p1, p2 := ga.scoreProgress[i], ga.scoreProgress[i-1]
		if (100.0 * math.Abs(p1-p2) / p2) > 2 {
			return false
		}
		if count >= 10 {
			return true
		}
		count++
	}
	return false
}

func (ga *geneticsAlgoImpl) recordProgress(score float64) {
	ga.scoreProgress = append(ga.scoreProgress, score)
}

func (ga *geneticsAlgoImpl) rankByConcurreny(populations []Chromosome) []scoreWrapper {
	scores := make([]scoreWrapper, ga.populationSize)
	mutex := sync.Mutex{}
	wg := sync.WaitGroup{}
	for i := int64(0); i < ga.populationSize; i++ {
		wg.Add(1)
		go func(i int64) {
			defer wg.Done()
			fitness := ga.chromModel.CalculateFitness(populations[i])
			mutex.Lock()
			scores[i] = scoreWrapper{
				fitnessScore: fitness,
				chromosome:   populations[i],
			}
			mutex.Unlock()
		}(i)
	}
	wg.Wait()
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].fitnessScore > scores[j].fitnessScore
	})
	return scores
}

func (ga *geneticsAlgoImpl) rank(populations []Chromosome) []scoreWrapper {
	scores := make([]scoreWrapper, ga.populationSize)
	for i := int64(0); i < ga.populationSize; i++ {
		fitness := ga.chromModel.CalculateFitness(populations[i])
		scores[i] = scoreWrapper{
			fitnessScore: fitness,
			chromosome:   populations[i],
		}
	}
	sort.Slice(scores, func(i, j int) bool {
		return scores[i].fitnessScore > scores[j].fitnessScore
	})
	return scores
}

func (ga *geneticsAlgoImpl) generateMatingPool(rankedPopulation []scoreWrapper) []Chromosome {
	matingPool := make([]Chromosome, ga.populationSize)
	cumSum := make([]float64, ga.populationSize)
	cumSum[0] = rankedPopulation[0].fitnessScore
	for i := int64(1); i < ga.populationSize; i++ {
		cumSum[i] = cumSum[i-1] + rankedPopulation[i].fitnessScore
	}

	totalSum := cumSum[len(cumSum)-1]
	for i := int64(0); i < ga.populationSize; i++ {
		cumSum[i] = cumSum[i] / totalSum
	}

	for i := int64(0); i < ga.populationSize; i++ {
		if i < ga.eliteSize {
			matingPool[i] = rankedPopulation[i].chromosome
		} else {
			pick := rand.Float64()
			for j := 0; j < len(rankedPopulation); j++ {
				if pick <= cumSum[j] {
					matingPool[i] = rankedPopulation[j].chromosome
					break
				}
			}
		}
	}
	for i := int64(1); i < ga.populationSize; i++ {
		if matingPool[i] == nil {
			fmt.Println(rankedPopulation)
		}
	}
	return matingPool
}

func (ga *geneticsAlgoImpl) mutate(populations []Chromosome) []Chromosome {
	for i := int64(0); i < ga.populationSize; i++ {
		if rand.Float64() < ga.mutationRate {
			populations[i] = ga.chromModel.Mutate(populations[i])
		}
	}
	return populations
}

func (ga *geneticsAlgoImpl) breedPopulation(matingPool []Chromosome) []Chromosome {

	children := make([]Chromosome, ga.populationSize)
	for i := int64(0); i < ga.eliteSize; i++ {
		children[i] = matingPool[i]
	}
	rand.Shuffle(len(matingPool), func(i, j int) {
		matingPool[i], matingPool[j] = matingPool[j], matingPool[i]
	})
	for i := int64(0); i < ga.populationSize-ga.eliteSize; i++ {
		children[ga.eliteSize+i] = ga.chromModel.Breed(matingPool[i], matingPool[ga.populationSize-1-i])
	}
	return children
}

func (ga *geneticsAlgoImpl) debugChromosome(pops []Chromosome) {
	for i := int64(0); i < ga.populationSize; i++ {
		fmt.Printf("%v ", pops[i].CalculateScore())
	}
	fmt.Printf("\n")
}
