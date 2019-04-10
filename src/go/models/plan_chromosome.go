package models

import(
	"../ga"
)

const(
	profitCoef = 0.4
	productVarietyCoef = 0.6
)

type Plan struct {
	orderedComponents []Component
	orderedMachines []Machine
	machineAssignment [][]Machine
}

func (p *Plan) CalculateScore() float64 {
	bestSchedulerSolver := NewScheduleSolverGA(Constraint{}, *p)
	bestOrderedPlan := ga.New(100,20,100, 0.0,bestSchedulerSolver).Solve().(*Plan)
	//bestOrderedPlan := p
	bestSchedule := NewGreedyScheduleFromPlan(bestOrderedPlan)

	productTypeCount := 0
	totalProfit := 0.0
	components := bestOrderedPlan.orderedComponents
	for i := int64(0) ;i< int64(len(components));i++{
		if len(bestOrderedPlan.machineAssignment[i]) > 0 {
			productTypeCount++
			totalProfit += components[i].GetProfit(bestOrderedPlan.machineAssignment[i]) / bestSchedule.GetCycleTime(i)
		}
	}

	return profitCoef * totalProfit + productVarietyCoef * float64(productTypeCount)
}

func NewPlan(constraint Constraint, MachineAssignment [][]Machine) *Plan {
	return &Plan{orderedComponents: constraint.Components, orderedMachines:constraint.Machines, machineAssignment: MachineAssignment}
}

func NewRandomPlan(constraint Constraint) *Plan {
	machineAssignments := make([][]Machine, len(constraint.Components))
	for i := 0; i < len(constraint.Components); i++ {
		machineAssignments[i] = constraint.Components[i].GenerateRandomAssignment(constraint.Machines)
	}
	return NewPlan(constraint, machineAssignments)
}
