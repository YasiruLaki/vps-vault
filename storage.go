package main

import (
	"encoding/json"
	"os"
)

func loadVPS(jsonPath string) []VPSData {

	data, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil
	}

	var vpsList []VPSData
	err = json.Unmarshal(data, &vpsList)

	return vpsList
}

func nameExists(name string, vpsList []VPSData) bool {
	for _, vps := range vpsList {
		if vps.Name == name {
			return true
		}
	}
	return false
}

func saveVPS(newVPS []VPSData, jsonPath string) error {
	data, err := json.MarshalIndent(newVPS, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(jsonPath, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
