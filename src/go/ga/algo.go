package ga

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

type geneticsAlgoImpl struct {
	populationSize, eliteSize, generations int64
	mutationRate                           float64
	chromModel                             ChromosomeModeller
}

type scoreWrapper struct {
	fitnessScore float64
	chromosome   Chromosome
}

func New(popSize, eliteSize, generations int64, mutationRate float64, chromModel ChromosomeModeller) GeneticsAlgo {
	return &geneticsAlgoImpl{
		populationSize: popSize,
		eliteSize:      eliteSize,
		generations:    generations,
		mutationRate:   mutationRate,
		chromModel:     chromModel,
	}
}

func (ga *geneticsAlgoImpl) Solve() Chromosome {
	rand.Seed(time.Now().UnixNano())
	populations := ga.initPopulation()
	for i := int64(0); i <= ga.generations; i++ {
		populations = ga.createNextGen(populations)
	}

	return ga.rank(populations)[0].chromosome
}

func (ga *geneticsAlgoImpl) initPopulation() []Chromosome {
	populations := make([]Chromosome, ga.populationSize)
	for i := int64(0); i < ga.populationSize; i++ {
		populations[i] = ga.chromModel.GenerateRandom()
	}
	return populations
}

func (ga *geneticsAlgoImpl) createNextGen(populations []Chromosome) []Chromosome {
	rankedPopulation := ga.rankByConcurreny(populations)
	matingPool := ga.generateMatingPool(rankedPopulation)
	children := ga.breedPopulation(matingPool)
	mutatedChildren := ga.mutate(children)
	return mutatedChildren
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
	for i:= int64(1) ;i< ga.populationSize; i++{
		cumSum[i] = cumSum[i-1] + rankedPopulation[i].fitnessScore
	}

	totalSum := cumSum[len(cumSum)-1]
	for i:= int64(0) ;i< ga.populationSize; i++{
		cumSum[i] = cumSum[i] / totalSum
	}

	for i := int64(0) ;i< ga.populationSize ; i++ {
		if i < ga.eliteSize {
			matingPool[i] = rankedPopulation[i].chromosome
		} else{
			pick := rand.Float64()
			for j := 0 ; j< len(rankedPopulation) ;i++{
				if pick <= cumSum[j] {
					matingPool[j] = rankedPopulation[j].chromosome
					break
				}
			}
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
