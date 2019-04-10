package models

import "math"

type MachineSchedule struct {
	ComponentName string
	ProcessName string
	StartTime, EndTime float64
}

type Schedule struct {
	plan *Plan
	machineSchedules map[string][]MachineSchedule
}

func (p *Schedule) GetCycleTime(componentIndex int64) float64 {
	// TODO :
	return 0.0
}

func (p *Schedule) GetMaxCyleTime() float64 {
	maxTime := 0.0
	for _, machineSched := range p.machineSchedules {
		maxTime = math.Max(maxTime, machineSched[len(machineSched)-1].EndTime)
	}
	return maxTime
}

func NewGreedyScheduleFromPlan(plan *Plan) *Schedule {
	// TODO :create greedy schedule
	return &Schedule{plan: plan}
}