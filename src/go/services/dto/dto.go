package dto

import (
	"../../models"
)

// ============================= REQUEST ================================================
type GaParams struct {
	PopSize     int64 `json:"popSize"`
	EliteSize   int64 `json:"eliteSize"`
	Generations int64 `json:"generations"`
	QuickScan   bool  `json:"quickScan"`
	UseCon      bool  `json:"useCon"`
}

type ComponentRequest struct {
	Name         string  `json:"name"`
	DesiredUnit  int64   `json:"desiredUnit"`
	Price        float64 `json:"price"`
	MaterialCost float64 `json:"materialCost"`
}

type RequestPayload struct {
	Components []ComponentRequest `json:"components"`
	QuickScan  bool               `json:"quickScan"`
}

// ============================= RESPONSE ================================================

type ComponentMetadata struct {
	ComponentName string  `json:"componentName"`
	UnitProduced  int64   `json:"unitProduced"`
	UnitProfit    float64 `json:"unitProfit"`
	CycleTime     float64 `json:"cycleTime"`
}

type Batches struct {
	MachineSchedule    map[string][]models.MachineSchedule `json:"machineSchedules"`
	StartTime          float64                             `json:"startTime"`
	EndTime            float64                             `json:"endTime"`
	ComponentsMetadata []ComponentMetadata                 `json:"componentsMetadata"`
}

type ResponsePayload struct {
	TotalProfit    float64            `json:"totalProfit"`
	ComponentsLeft []ComponentRequest `json:"componentsLeft"`
	Batches        []Batches          `json:"batches"`
}
