package models

type Machine struct {
	Name         string          `json:"name"`
	Cost         float64         `json:"cost"`
	ProcessNames map[string]bool `json:"handledProcess"`
}

func (m *Machine) HasProcess(processName string) bool {
	return m.ProcessNames[processName]
}

func NewMachine(name string, cost float64, processNames []string) Machine {
	m := Machine{Name: name, Cost: cost, ProcessNames: map[string]bool{}}
	for _, processName := range processNames {
		m.ProcessNames[processName] = true
	}
	return m
}

func NewMachineWithProcessMap(name string, cost float64, processMap map[string]bool) Machine {
	m := Machine{Name: name, Cost: cost, ProcessNames: processMap}
	return m
}
