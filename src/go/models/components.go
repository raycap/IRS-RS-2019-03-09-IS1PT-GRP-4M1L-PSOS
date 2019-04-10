package models

import "math/rand"

type ComponentProcess struct {
	Name     string
	Duration float64
}

type Component struct {
	Name                string
	Price, MarginProfit float64
	Processes           []ComponentProcess
}

func (c *Component) GetProfit(sortedMachineAssignment []Machine) float64 {
	return c.getFromPrice(sortedMachineAssignment)
}

func (c *Component) GenerateRandomAssignment(machineList []Machine) []Machine {
	machines := make([]Machine, len(c.Processes))
	for i:= 0 ; i <len(c.Processes);i++{
		processName := c.Processes[i].Name
		candidates := []Machine{}
		for j := 0 ; j< len(machineList);j++{
			if processName == machineList[j].Name {
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
	for i := 0 ;i<len(sortedMachineAssignment);i++{
		totalCost += sortedMachineAssignment[i].Cost * c.Processes[i].Duration
	}

	return c.Price - totalCost
}
