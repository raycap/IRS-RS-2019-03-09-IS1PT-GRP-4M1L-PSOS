package models

import (
	"fmt"
	"math"

	keytranslation "../key_translation"
)

type MachineSchedule struct {
	ComponentName string  `json:"componentName"`
	ProcessName   string  `json:"processName"`
	StartTime     float64 `json:"startTime"`
	EndTime       float64 `json:"endTime"`
}

type scheduleUnit struct {
	machineName, componentName, processName string
	duration                                float64
}

type Schedule struct {
	plan             *Plan
	machineSchedules map[string][]MachineSchedule
}

func (s *Schedule) GetCycleTime(componentIndex int64) float64 {
	return s.buildCycleTime(componentIndex)
}

func (s *Schedule) buildCycleTime(componentIndex int64) float64 {
	componentNameToIndex := s.plan.ComponentNameToIndex()
	componentName := s.plan.orderedComponents[componentIndex].Name
	componentMap := map[string][]string{}
	machineMap := map[string]bool{}
	machineList := []string{}
	// get init machines that are directly affected
	if len(s.plan.machineAssignment[componentIndex]) == 0 {
		return 0.0
	}

	for _, machine := range s.plan.machineAssignment[componentIndex] {
		machineName := machine.Name
		machineList = append(machineList, machineName)
		machineMap[machineName] = true
		if _, ok := componentMap[componentName]; !ok {
			componentMap[componentName] = []string{machineName}
		} else {
			componentMap[componentName] = append(componentMap[componentName], machineName)
		}
	}
	counter := 0
	for {
		if counter >= len(machineList) {
			break
		}
		machineName := machineList[counter]
		machineScheds := s.machineSchedules[machineName]
		newComponentNames := []string{}
		for _, machineSched := range machineScheds {
			componentName := machineSched.ComponentName
			if _, ok := componentMap[componentName]; !ok {
				newComponentNames = append(newComponentNames, componentName)
			}
		}

		for _, componentName := range newComponentNames {
			componentIndex := componentNameToIndex[componentName]
			candidateMachineList := s.plan.machineAssignment[componentIndex]
			for _, candidateMachine := range candidateMachineList {
				machineName := candidateMachine.Name
				if _, ok := componentMap[componentName]; !ok {
					componentMap[componentName] = []string{machineName}
				} else {
					componentMap[componentName] = append(componentMap[componentName], machineName)
				}

				if _, ok := machineMap[machineName]; !ok {
					machineList = append(machineList, machineName)
					machineMap[machineName] = true
				}
			}
		}

		counter++
	}

	maxCycleTime := 0.0
	affectedMachines := componentMap[componentName]
	for _, machineName := range affectedMachines {
		machineSched := s.machineSchedules[machineName]
		maxCycleTime = math.Max(maxCycleTime, machineSched[len(machineSched)-1].EndTime-machineSched[0].StartTime)
	}
	return maxCycleTime
}

func (s *Schedule) GetMaxCyleTime() float64 {
	maxTime := 0.0
	for _, machineSched := range s.machineSchedules {
		maxTime = math.Max(maxTime, machineSched[len(machineSched)-1].EndTime)
	}
	return maxTime
}

func (s *Schedule) GetMachineSchedule() map[string][]MachineSchedule {
	return s.machineSchedules
}

func (s *Schedule) GetSimulatedMachineSchedule() map[string][]MachineSchedule {
	maxTime := s.GetMaxCyleTime()
	simulatedMachineSchedules := map[string][]MachineSchedule{}
	for machineName, machineSchedules := range s.machineSchedules {
		if machineSchedules[len(machineSchedules)-1].EndTime >= maxTime {
			simulatedMachineSchedules[machineName] = machineSchedulesWithTranslation(machineSchedules)
			continue
		}
		componentName := machineSchedules[0].ComponentName
		cycleTime := 0.0
		for index, c := range s.plan.orderedComponents {
			if c.Name == componentName {
				cycleTime = s.buildCycleTime(int64(index))
				break
			}
		}
		if cycleTime == 0 {
			simulatedMachineSchedules[machineName] = machineSchedulesWithTranslation(machineSchedules)
			continue
		}
		i := 0
		for {
			ms := machineSchedules[i]
			startTime := ms.StartTime + cycleTime
			endTime := ms.EndTime + cycleTime
			if endTime >= maxTime {
				break
			}
			machineSchedules = append(machineSchedules, MachineSchedule{
				ComponentName: ms.ComponentName, ProcessName: ms.ProcessName, StartTime: startTime, EndTime: endTime})
			i++
		}

		simulatedMachineSchedules[machineName] = machineSchedulesWithTranslation(machineSchedules)
	}

	return simulatedMachineSchedules
}

func NewGreedyScheduleFromPlan(plan *Plan) *Schedule {
	queue := []scheduleUnit{}
	i := 0

	for {
		count := 0
		for j, componentPlan := range plan.machineAssignment {
			if i >= len(componentPlan) {
				count++
				continue
			}
			if i >= len(plan.orderedComponents[j].Processes) {
				fmt.Println("panic !!!!")
				fmt.Println(plan.orderedComponents)
				fmt.Println(componentPlan)
				fmt.Println(i, j)
				panic("panic!!")
			}
			machineChosen := componentPlan[i]
			queue = append(queue, scheduleUnit{machineName: machineChosen.Name,
				componentName: plan.orderedComponents[j].Name,
				duration:      plan.orderedComponents[j].Processes[i].Duration,
				processName:   plan.orderedComponents[j].Processes[i].Name})
		}
		i++
		if count >= len(plan.machineAssignment) {
			break
		}
	}
	componentTaskTime := map[string]float64{}
	greedyArrangement := map[string][]MachineSchedule{}
	for _, queueTask := range queue {
		queueTaskMachineName := queueTask.machineName
		queueTaskComponentName := queueTask.componentName
		queueTaskProcessName := queueTask.processName
		duration := queueTask.duration
		var startTime, endTime float64
		if _, ok := greedyArrangement[queueTaskMachineName]; !ok {
			startTime = 0.0
			if currentTime, ok := componentTaskTime[queueTaskComponentName]; ok {
				startTime = currentTime
			}
			endTime = startTime + duration
			greedyArrangement[queueTaskMachineName] = []MachineSchedule{{
				ComponentName: queueTaskComponentName, ProcessName: queueTaskProcessName, StartTime: startTime, EndTime: endTime}}
		} else {
			startTime = greedyArrangement[queueTaskMachineName][len(greedyArrangement[queueTaskMachineName])-1].EndTime
			if currentTime, ok := componentTaskTime[queueTaskComponentName]; ok {
				startTime = math.Max(startTime, currentTime)
			}
			endTime = startTime + duration
			greedyArrangement[queueTaskMachineName] = append(greedyArrangement[queueTaskMachineName], MachineSchedule{
				ComponentName: queueTaskComponentName, ProcessName: queueTaskProcessName, StartTime: startTime, EndTime: endTime})
		}

		componentTaskTime[queueTaskComponentName] = endTime
	}
	return &Schedule{plan: plan, machineSchedules: greedyArrangement}
}

func machineSchedulesWithTranslation(ms []MachineSchedule) []MachineSchedule {
	for index, _ := range ms {
		ms[index].ProcessName = keytranslation.Get(ms[index].ProcessName)
		ms[index].ComponentName = keytranslation.Get(ms[index].ComponentName)
	}
	return ms
}
