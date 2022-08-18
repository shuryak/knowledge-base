package main

import "fmt"

func main() {
	statesNeeded := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}

	stations := make(map[string][]string)
	stations["kone"] = []string{"id", "nv", "ut"}
	stations["ktwo"] = []string{"wa", "id", "mt"}
	stations["kthree"] = []string{"or", "nv", "ca"}
	stations["kfour"] = []string{"nv", "ut"}
	stations["kfive"] = []string{"ca", "az"}

	stationKey := []string{"kone", "ktwo", "kthree", "kfour", "kfive"}

	var finalStations []string

	for len(statesNeeded) > 0 {
		var bestStation string
		var statesCovered []string

		for _, station := range stationKey {
			states := stations[station]
			covered := getEqual(statesNeeded, states)
			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}
		statesNeeded = removeManyElements(statesNeeded, statesCovered)
		finalStations = append(finalStations, bestStation)
	}

	fmt.Println(finalStations)
}

// Implementing of set:

func getEqual(arr1, arr2 []string) []string {
	var covered []string

	for _, stateNeeded := range arr1 {
		for _, state := range arr2 {
			if stateNeeded == state {
				covered = append(covered, stateNeeded)
			}
		}
	}

	return covered
}

func removeManyElements(arr, elementsToRemove []string) []string {
	for _, stateCovered := range elementsToRemove {
		arr = removeElement(arr, stateCovered)
	}
	return arr
}

func removeElement(arr []string, elementToRemove string) []string {
	for i, stateNeeded := range arr {
		if elementToRemove == stateNeeded {
			return append(arr[:i], arr[i+1:]...)
		}
	}
	return arr
}
