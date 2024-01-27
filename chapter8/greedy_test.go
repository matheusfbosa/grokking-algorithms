package chapter8_test

import (
	"reflect"
	"testing"

	"github.com/matheusfbosa/grokking-algorithms/chapter8"
)

func TestFindFinalRadioStations(t *testing.T) {
	tests := []struct {
		name          string
		statesToCover map[string]bool
		radioStations map[string]map[string]bool
		want          map[string]bool
	}{
		{
			name:          "RadioStationsFound",
			statesToCover: map[string]bool{"mt": true, "wa": true, "or": true, "id": true, "nv": true, "ut": true, "ca": true, "az": true},
			radioStations: map[string]map[string]bool{
				"kone":   {"id": true, "nv": true, "ut": true},
				"ktwo":   {"wa": true, "id": true, "mt": true},
				"kthree": {"or": true, "nv": true, "ca": true},
				"kfour":  {"nv": true, "ut": true},
				"kfive":  {"ca": true, "az": true},
			},
			want: map[string]bool{"kone": true, "ktwo": true, "kthree": true, "kfive": true},
		},
		{
			name:          "NoStationsToCover",
			statesToCover: map[string]bool{},
			radioStations: map[string]map[string]bool{
				"kone":   {"id": true, "nv": true, "ut": true},
				"ktwo":   {"wa": true, "id": true, "mt": true},
				"kthree": {"or": true, "nv": true, "ca": true},
			},
			want: map[string]bool{},
		},
		{
			name:          "NoRadioStations",
			statesToCover: map[string]bool{"mt": true, "wa": true, "or": true},
			radioStations: map[string]map[string]bool{},
			want:          map[string]bool{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chapter8.FindFinalRadioStations(tt.statesToCover, tt.radioStations)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Test %s failed. Want: %v, Got: %v", tt.name, tt.want, got)
			}
		})
	}
}
