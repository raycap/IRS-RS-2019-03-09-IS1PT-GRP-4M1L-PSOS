package models

import (
	"../ga"
)

const (
	profitCoef         = 0.5
	productVarietyCoef = 0.5
)

type Plan struct {
	orderedComponents []Component
	orderedMachines   []Machine
	machineAssignment [][]Machine
	score             float64
	schedule          *Schedule
	quickScan         bool
}

func (p *Plan) CalculateScore() float64 {
	bestOrderedPlan, bestSchedule := p.getBestSched()

	productTypeCount := 0
	totalProfit := 0.0
	components := bestOrderedPlan.orderedComponents
	for i := int64(0); i < int64(len(components)); i++ {
		if len(bestOrderedPlan.machineAssignment[i]) > 0 {
			productTypeCount++
			totalProfit += components[i].GetProfit(bestOrderedPlan.machineAssignment[i]) / bestSchedule.GetCycleTime(i)
		}
	}

	return profitCoef*totalProfit + productVarietyCoef*float64(productTypeCount)
}

func (p *Plan) ComponentNameToIndex() map[string]int64 {
	reverseIndex := map[string]int64{}
	for i := int64(0); i < int64(len(p.orderedComponents)); i++ {
		reverseIndex[p.orderedComponents[i].Name] = i
	}
	return reverseIndex
}

func (p *Plan) GetMachineAssignment() [][]Machine {
	return p.machineAssignment
}

func (p *Plan) GetScore() float64 {
	if p.score == 0 {
		p.score = p.CalculateScore()
	}
	return p.score
}

func (p *Plan) GetBestSchedule() *Schedule {
	if p.schedule == nil {
		_, p.schedule = p.getBestSched()
	}
	return p.schedule
}

func (p *Plan) getBestSched() (*Plan, *Schedule) {
	var bestOrderedPlan *Plan
	if p.quickScan {
		bestOrderedPlan = p
	} else {
		bestSchedulerSolver := NewScheduleSolverGA(Constraint{Machines: p.orderedMachines, Components: p.orderedComponents}, *p)
		bestOrderedPlan = ga.New(100, 15, 50, 0.02, bestSchedulerSolver, true).Solve().(*Plan)
	}
	bestSchedule := NewGreedyScheduleFromPlan(bestOrderedPlan)
	return p, bestSchedule
}

func NewPlan(constraint Constraint, MachineAssignment [][]Machine, quickScan bool) *Plan {
	return &Plan{orderedComponents: constraint.Components, orderedMachines: constraint.Machines,
		machineAssignment: MachineAssignment, quickScan: quickScan}
}

func NewRandomPlan(constraint Constraint, quickScan bool) *Plan {
	machineAssignments := make([][]Machine, len(constraint.Components))
	for i := 0; i < len(constraint.Components); i++ {
		machineAssignments[i] = constraint.Components[i].GenerateRandomAssignment(constraint.Machines)
	}
	return NewPlan(constraint, machineAssignments, quickScan)
}
