package models

import (
	"math/rand"
)

type ComponentProcess struct {
	Name     string  `json:"processName"`
	Duration float64 `json:"duration"`
}

type Component struct {
	Name         string             `json:"name"`
	Price        float64            `json:"price"`
	MarginProfit float64            `json:"marginProfit"`
	Processes    []ComponentProcess `json:"process"`
}

func NewComponent(name string, price, marginProfit float64, process []ComponentProcess) Component {
	return Component{
		Name: name, Price: price, MarginProfit: marginProfit, Processes: process,
	}
}

func (c *Component) GetProfit(sortedMachineAssignment []Machine) float64 {
	return c.getFromPrice(sortedMachineAssignment)
}

func (c *Component) GenerateRandomAssignment(machineList []Machine) []Machine {
	machines := make([]Machine, len(c.Processes))
	for i := 0; i < len(c.Processes); i++ {
		processName := c.Processes[i].Name
		candidates := []Machine{}
		for j := 0; j < len(machineList); j++ {
			if machineList[j].HasProcess(processName) {
				candidates = append(candidates, machineList[j])
			}
		}
		machines[i] = candidates[rand.Int63n(int64(len(candidates)))]
	}
	return machines
}

func (c *Component) getFromMarginalProfit() float64 {
	return c.MarginProfit
}

func (c *Component) getFromPrice(sortedMachineAssignment []Machine) float64 {
	var totalCost float64
	for i := 0; i < len(sortedMachineAssignment); i++ {
		totalCost += sortedMachineAssignment[i].Cost * c.Processes[i].Duration
	}

	return c.Price - totalCost
}
