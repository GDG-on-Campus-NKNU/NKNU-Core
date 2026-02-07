package classroom

import (
	"testing"
)

func TestQueryClassroom(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantBuilding string
		wantFloor    string
		wantID       string
		wantMapURL   string
		wantErr      bool
		wantNil      bool
	}{
		{
			name:    "Short input",
			input:   "A",
			wantNil: true,
		},
		{
			name:         "Two-letter building (MA)",
			input:        "MA816",
			wantBuilding: "致理大樓",
			wantFloor:    "8",
			wantID:       "16",
			wantMapURL:   "https://github.com/GDG-on-Campus-NKNU/NKNU-Assets/raw/main/classroom_map/images/MA8.png",
		},
		{
			name:         "Two-letter building case-insensitive",
			input:        "tc201",
			wantBuilding: "科技大樓",
			wantFloor:    "2",
			wantID:       "01",
			wantMapURL:   "https://github.com/GDG-on-Campus-NKNU/NKNU-Assets/raw/main/classroom_map/images/TC2.png",
		},
		{
			name:         "One-letter building (3)",
			input:        "3101",
			wantBuilding: "文學大樓",
			wantFloor:    "1",
			wantID:       "01",
			wantMapURL:   "https://github.com/GDG-on-Campus-NKNU/NKNU-Assets/raw/main/classroom_map/images/31.png",
		},
		{
			name:         "One-letter building with space",
			input:        " 0202 ",
			wantBuilding: "行政大樓",
			wantFloor:    "2",
			wantID:       "02",
			wantMapURL:   "https://github.com/GDG-on-Campus-NKNU/NKNU-Assets/raw/main/classroom_map/images/02.png",
		},
		{
			name:         "One-letter building with space",
			input:        " 0102 ",
			wantBuilding: "行政大樓",
			wantFloor:    "1",
			wantID:       "02",
			wantMapURL:   "", // 01 is not in _floors
		},
		{
			name:    "Unknown building",
			input:   "XX101",
			wantNil: true,
		},
		{
			name:    "Unknown single digit building (2 is missing in map)",
			input:   "2101",
			wantNil: true,
		},
		{
			name:    "Empty input",
			input:   "",
			wantNil: true,
		},
		{
			name:         "Two-character input (1-letter building)",
			input:        "31",
			wantBuilding: "文學大樓",
			wantFloor:    "1",
			wantID:       "",
			wantMapURL:   "https://github.com/GDG-on-Campus-NKNU/NKNU-Assets/raw/main/classroom_map/images/31.png",
		},
		{
			name:         "Valid building but no map",
			input:        "BT999",
			wantBuilding: "生科大樓",
			wantFloor:    "9",
			wantID:       "99",
			wantMapURL:   "", // BT9 is not in _floors
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := QueryClassroom(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryClassroom() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantNil {
				if got != nil {
					t.Errorf("QueryClassroom() got = %v, want nil", got)
				}
				return
			}
			if got == nil {
				t.Errorf("QueryClassroom() got nil, want %v", tt.wantBuilding)
				return
			}
			if got.BuildingName != tt.wantBuilding {
				t.Errorf("BuildingName = %v, want %v", got.BuildingName, tt.wantBuilding)
			}
			if got.Floor != tt.wantFloor {
				t.Errorf("Floor = %v, want %v", got.Floor, tt.wantFloor)
			}
			if got.ClassroomID != tt.wantID {
				t.Errorf("ClassroomID = %v, want %v", got.ClassroomID, tt.wantID)
			}
			if got.MapURL != tt.wantMapURL {
				t.Errorf("MapURL = %v, want %v", got.MapURL, tt.wantMapURL)
			}
		})
	}
}
