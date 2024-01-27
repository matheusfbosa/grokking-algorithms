package chapter8

import "github.com/matheusfbosa/grokking-algorithms/helper"

func FindFinalRadioStations(
	statesToCover map[string]bool,
	radioStations map[string]map[string]bool,
) map[string]bool {
	finalRadioStations := make(map[string]bool)

	if len(statesToCover) == 0 || len(radioStations) == 0 {
		return finalRadioStations
	}

	for len(statesToCover) > 0 {
		bestStation := ""
		statesCovered := make(map[string]bool)

		for station, states := range radioStations {
			covered := helper.SetIntersection(statesToCover, states)

			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
			}
		}

		for state := range statesCovered {
			delete(statesToCover, state)
		}

		finalRadioStations[bestStation] = true
	}

	return finalRadioStations
}
