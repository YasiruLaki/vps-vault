package main

import (
	"encoding/json"
	"fmt"
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

func saveVPS(newVPS []VPSData, jsonPath string) {
	data, err := json.MarshalIndent(newVPS, "", "  ")
	if err != nil {
		fmt.Println("An error occurred while saving VPS data.")
		return
	}

	err = os.WriteFile(jsonPath, data, 0644)

	if err != nil {
		fmt.Println("An error occurred while writing VPS data to file.")
	} else {
		fmt.Println("VPS data saved successfully. Run 'vps list' to see your updated VPS list.")
	}
}
