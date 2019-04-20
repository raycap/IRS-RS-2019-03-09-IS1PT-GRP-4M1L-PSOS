package services

import (
	"math"

	"fmt"

	"../ga"
	keytranslation "../key_translation"
	"../models"
	"./dto"
)

func ResultWithGaParam(reqParams *dto.GaParams) (dto.ResponsePayload, error) {
	modeller := generateModels(reqParams.QuickScan)

	gaSolver := ga.New(reqParams.PopSize, reqParams.EliteSize, reqParams.Generations, 0.07, modeller, reqParams.UseCon)
	bestPlan := gaSolver.Solve().(*models.Plan)
	fmt.Println("score")
	fmt.Println(bestPlan.CalculateScore())
	return dto.ResponsePayload{
		Batches: []dto.Batches{
			{MachineSchedule: bestPlan.GetBestSchedule().GetMachineSchedule()},
		},
	}, nil
}

func SolveWithFixtures(reqParams *dto.RequestPayload) (dto.ResponsePayload, error) {
	allMachines, allComponents, err := generateFromFixtures()
	if err != nil {
		return dto.ResponsePayload{}, err
	}

	return solveRequest(models.Constraint{Machines: allMachines, Components: allComponents}, reqParams.Components, reqParams.QuickScan)

}

func solveRequest(constraint models.Constraint, componentsRequest []dto.ComponentRequest, quickScan bool) (dto.ResponsePayload, error) {
	batches := []dto.Batches{}
	componentsMap := map[string]int{}
	componentReqMap := map[string]int{}
	for i, com := range constraint.Components {
		componentsMap[com.Name] = i
	}
	for i, com := range componentsRequest {
		componentReqMap[com.Name] = i
	}

	currentTime := 0.0
	maxAllowedTime := float64(30 * 8 * 60) // 8 hours everyday
	totalProfit := 0.0
	for {
		if currentTime >= maxAllowedTime {
			break
		}
		// create all components that need to be solved
		componentsQueried := []models.Component{}
		for _, com := range componentsRequest {
			if com.DesiredUnit > 0 {
				constraint.Components[componentsMap[com.Name]].Price = com.Price
				constraint.Components[componentsMap[com.Name]].MaterialCost = com.MaterialCost
				c := constraint.Components[componentsMap[com.Name]]
				componentsQueried = append(componentsQueried, c)
			}
		}
		if len(componentsQueried) == 0 {
			break
		}

		// create and solve GA
		modeller := models.NewPlanGASolverModel(constraint.Machines, componentsQueried, quickScan)
		gaSolver := ga.New(100, 20, 100, 0.05, modeller, true)
		bestPlan := gaSolver.Solve().(*models.Plan)
		jobSched := bestPlan.GetBestSchedule()

		// process the result into batch
		minDuration := maxAllowedTime - currentTime
		for i, qCom := range componentsQueried {
			cycleTime := jobSched.GetCycleTime(int64(i))
			if cycleTime == 0.0 {
				continue
			}
			totalDurationNeeded := cycleTime * float64(componentsRequest[componentReqMap[qCom.Name]].DesiredUnit)
			minDuration = math.Min(minDuration, totalDurationNeeded)
		}

		compMetadata := []dto.ComponentMetadata{}
		for i, qCom := range componentsQueried {
			cycleTime := jobSched.GetCycleTime(int64(i))
			if cycleTime == 0.0 {
				continue
			}
			unitProduced := int64(math.Floor(minDuration / cycleTime))
			if unitProduced == 0 {
				continue
			}
			componentsRequest[componentReqMap[qCom.Name]].DesiredUnit -= unitProduced
			if componentsRequest[componentReqMap[qCom.Name]].DesiredUnit <= 0 {
				componentsRequest[componentReqMap[qCom.Name]].DesiredUnit = 0
			}
			unitProfit := float64(int64(constraint.Components[componentsMap[qCom.Name]].GetProfit(bestPlan.GetMachineAssignment()[i])*100)) / 100.0

			compMetadata = append(compMetadata, dto.ComponentMetadata{
				ComponentName: keytranslation.Get(qCom.Name), UnitProduced: unitProduced, CycleTime: cycleTime,
				UnitProfit: unitProfit,
			})
			totalProfit += unitProfit * float64(unitProduced)
		}
		if len(compMetadata) == 0 {
			break
		}

		minDuration += jobSched.GetMaxCyleTime()
		batch := dto.Batches{
			MachineSchedule:    jobSched.GetSimulatedMachineSchedule(),
			StartTime:          currentTime,
			EndTime:            currentTime + minDuration,
			ComponentsMetadata: compMetadata,
		}
		batches = append(batches, batch)
		currentTime += minDuration
	}
	for i, _ := range componentsRequest {
		componentsRequest[i].Name = keytranslation.Get(componentsRequest[i].Name)
	}
	return dto.ResponsePayload{
		ComponentsLeft: componentsRequest,
		TotalProfit:    totalProfit,
		Batches:        batches,
	}, nil
}
