package models

import (
	"../ga"
	"math/rand"
)

type ScheduleSolver struct {
	constraint            Constraint
	plan Plan
}

func NewScheduleSolverGA(constraint Constraint, plainPlan Plan) ga.ChromosomeModeller {
	return &ScheduleSolver{
		constraint:            constraint,
		plan:	plainPlan,
	}
}

func (ss *ScheduleSolver) CalculateFitness(chromosome ga.Chromosome) float64 {
	shuffledPlan := chromosome.(*Plan)
	sched := NewGreedyScheduleFromPlan(shuffledPlan)
	return 1.0 / sched.GetMaxCyleTime()
}

func (ss *ScheduleSolver) GenerateRandom() ga.Chromosome {
	components := ss.plan.orderedComponents
	machineAssignments := ss.plan.machineAssignment

	rand.Shuffle(len(originalComponents), func(i, j int) {
		components[i], components[j] = components[j], components[i]
		machineAssignments[i], machineAssignments[j] = machineAssignments[j], machineAssignments[i]
	})
	return NewPlan(Constraint{Components:components, Machines:ss.constraint.Machines}, machineAssignments)
}

func (ss *ScheduleSolver) Breed(firstParent, secondParend ga.Chromosome) ga.Chromosome {
	firstShuffledPlan := firstParent.(*Plan)
	secondShuffledPlan := secondParend.(*Plan)

	l := len(firstShuffledPlan.orderedComponents)
	i := rand.Int63n(int64(l))
	j := rand.Int63n(int64(l))
	var left , right int64
	if i< j {
		left, right = i, j
	}else{
		right, left = i, j
	}

	newMachineAssignments := make([][]Machine, l)
	newOrderedComponents := make([]Component, l)
	componentsName := map[string]bool{}
	count := 0
	for i := left ; i< right ; i++ {
		newOrderedComponents[count] = firstShuffledPlan.orderedComponents[i]
		newMachineAssignments[count] = firstShuffledPlan.machineAssignment[i]
		componentsName[firstShuffledPlan.orderedComponents[i].Name] = true
		count++
	}

	for i := 0 ; i<l ; i++{
		compName := secondShuffledPlan.orderedComponents[i].Name
		if _,ok := componentsName[compName]; !ok {
			newOrderedComponents[count] = secondShuffledPlan.orderedComponents[i]
			newMachineAssignments[count] = secondShuffledPlan.machineAssignment[i]
			componentsName[secondShuffledPlan.orderedComponents[i].Name] = true
			count++
		}
	}
	return NewPlan(Constraint{Components:newOrderedComponents, Machines: ss.constraint.Machines}, newMachineAssignments)
}

func (ss *ScheduleSolver) Mutate(chromosome ga.Chromosome) ga.Chromosome {
	return chromosome
}
