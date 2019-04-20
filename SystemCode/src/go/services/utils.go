package services

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"

	"fmt"

	"../ga"
	"../models"
)

type MachineFixtures struct {
	models.Machine
	Quantity int `json:"quantity"`
}

func generateModels(quickScan bool) ga.ChromosomeModeller {
	machines := []models.Machine{
		models.NewMachine("M1", 1.5, []string{"P1", "P2"}),
		models.NewMachine("M2", 0.75, []string{"P1"}),
		models.NewMachine("M3", 3.0, []string{"P1", "P2", "P3"}),
		models.NewMachine("M4", 2.0, []string{"P1", "P3"}),
		models.NewMachine("M5", 3.0, []string{"P5"}),
		models.NewMachine("M6", 1.5, []string{"P2", "P3"}),
	}

	components := []models.Component{
		models.NewComponent("C1", 10.0, 0.0,
			[]models.ComponentProcess{{Name: "P1", Duration: 3.0}}),
		models.NewComponent("C2", 12.0, 0.0,
			[]models.ComponentProcess{{Name: "P1", Duration: 2.0}, {Name: "P2", Duration: 3.0}}),
		models.NewComponent("C3", 8.0, 0.0,
			[]models.ComponentProcess{{Name: "P2", Duration: 3.0}}),
		models.NewComponent("C4", 14.0, 0.0,
			[]models.ComponentProcess{{Name: "P2", Duration: 3.5}, {Name: "P3", Duration: 2.0}}),
		models.NewComponent("C5", 25.0, 0.0,
			[]models.ComponentProcess{{Name: "P1", Duration: 4.0}, {Name: "P2", Duration: 2.5}, {Name: "P3", Duration: 1.0}}),
		models.NewComponent("C6", 20.0, 0.0,
			[]models.ComponentProcess{{Name: "P1", Duration: 2.0}, {Name: "P3", Duration: 2.0}, {Name: "P2", Duration: 2.0}}),
		models.NewComponent("C7", 15.0, 0.0,
			[]models.ComponentProcess{{Name: "P5", Duration: 2.0}}),
		models.NewComponent("C8", 14.0, 0.0,
			[]models.ComponentProcess{{Name: "P3", Duration: 3.5}, {Name: "P2", Duration: 2.0}}),
	}
	return models.NewPlanGASolverModel(machines, components, quickScan)
}

func generateFromFixtures() ([]models.Machine, []models.Component, error) {
	m, err := parseJSON("./fixtures/machines.json", &[]MachineFixtures{})
	if err != nil {
		return nil, nil, err
	}

	c, err := parseJSON("./fixtures/components.json", &[]models.Component{})
	if err != nil {
		return nil, nil, err
	}
	components := c.(*[]models.Component)

	machines := []models.Machine{}
	for _, machine := range *m.(*[]MachineFixtures) {
		for i := 1; i <= machine.Quantity; i++ {
			machines = append(machines, models.NewMachineWithProcessMap(fmt.Sprintf("%s-%d", machine.Name, i), machine.Cost, machine.ProcessNames))
		}
	}
	return machines, *components, nil
}

func parseJSON(fixturesPath string, defaultValue interface{}) (interface{}, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(fixturesPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		return defaultValue, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var copied interface{}
	if reflect.ValueOf(defaultValue).Kind() == reflect.Ptr {
		copied = reflect.New(reflect.ValueOf(defaultValue).Elem().Type()).Interface()
	} else {
		copied = reflect.New(reflect.ValueOf(defaultValue).Type()).Interface()
	}
	if err := json.Unmarshal([]byte(byteValue), copied); err != nil {
		return defaultValue, err
	}

	return copied, nil
}
