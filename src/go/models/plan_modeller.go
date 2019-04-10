package models

import (
	"math/rand"

	"../ga"
)

type Constraint struct {
	Machines   []Machine
	Components []Component
}

type PlanModeller struct {
	componentMutationRate float64
	constraint            Constraint
}

func NewPlanGASolverModel(Machines []Machine, Components []Component) ga.ChromosomeModeller {
	return &PlanModeller{
		constraint:            Constraint{Machines: Machines, Components: Components},
		componentMutationRate: 0.05,
	}
}

func (pm *PlanModeller) CalculateFitness(chromosome ga.Chromosome) float64 {
	plan := chromosome.(*Plan)
	return plan.CalculateScore()
}

func (pm *PlanModeller) GenerateRandom() ga.Chromosome {
	return NewRandomPlan(pm.constraint)
}

func (pm *PlanModeller) Breed(firstParent, secondParend ga.Chromosome) ga.Chromosome {
	firstPlan := firstParent.(*Plan)
	secondPlan := secondParend.(*Plan)
	chromeLen := int64(len(pm.constraint.Components))
	r := rand.Int63n(chromeLen)

	// TODO : Some permutation between inter gene between 2 chromes
	newMachineAssignment := make([][]Machine, chromeLen)
	for i := int64(0); i < chromeLen; i++ {
		if i < r {
			newMachineAssignment[i] = firstPlan.machineAssignment[i]
		} else {
			newMachineAssignment[i] = secondPlan.machineAssignment[i]
		}
	}

	return NewPlan(pm.constraint, newMachineAssignment)
}

func (pm *PlanModeller) Mutate(chromosome ga.Chromosome) ga.Chromosome {
	chromeLen := int64(len(pm.constraint.Components))
	plan := chromosome.(*Plan)
	for i := int64(0); i < chromeLen; i++ {
		if rand.Float64() < pm.componentMutationRate {
			plan.machineAssignment[i] = pm.constraint.Components[i].GenerateRandomAssignment(pm.constraint.Machines)
		}
	}
	return plan
}
