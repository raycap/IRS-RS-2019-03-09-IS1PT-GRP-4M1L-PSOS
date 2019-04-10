package models

type Machine struct {
	Name         string
	Cost         float64
	ProcessNames map[string]bool
}

func (m *Machine) HasProcess(processName string) bool {
	return m.ProcessNames[processName]
}

func NewMachine(name string, cost float64, processNames []string) *Machine {
	m := &Machine{Name: name, Cost: cost, ProcessNames: map[string]bool{}}
	for _, processName := range processNames {
		m.ProcessNames[processName] = true
	}
	return m
}
