package controllers

import (
	"github.com/dasvh/go-learn-vim/internal/models"
	"testing"
)

func Test_NewLevel(t *testing.T) {
	tests := []struct {
		name       string
		wantLevels []int
	}{
		{
			name:       "Initialize Level controller",
			wantLevels: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := NewLevel()
			for _, levelNum := range tt.wantLevels {
				if _, exists := lc.GetLevels()[levelNum]; !exists {
					t.Errorf("expected level %d to be initialized, but it was not", levelNum)
				}
			}
			if lc.GetCurrentLevel() != nil {
				t.Errorf("expected current level to be nil, got %v", lc.GetCurrentLevel())
			}
		})
	}
}

func Test_LevelGetters(t *testing.T) {
	tests := []struct {
		name         string
		setup        func() *Level
		wantLevels   int
		wantLevelNum int
	}{
		{
			name: "Get levels and current level for level 0",
			setup: func() *Level {
				lc := NewLevel()
				lc.SetLevel(lc.GetLevels()[0])
				return lc
			},
			wantLevels:   2,
			wantLevelNum: 0,
		},
		{
			name: "Get levels and current level for level 1",
			setup: func() *Level {
				lc := NewLevel()
				lc.SetLevel(lc.GetLevels()[1])
				return lc
			},
			wantLevels:   2,
			wantLevelNum: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := tt.setup()

			if lc.GetLevelsCount() != tt.wantLevels {
				t.Errorf("expected %d levels, got %d", tt.wantLevels, lc.GetLevelsCount())
			}

			if lc.GetLevelNumber() != tt.wantLevelNum {
				t.Errorf("expected current level number %d, got %d", tt.wantLevelNum, lc.GetLevelNumber())
			}
		})
	}
}

func Test_RestoreLevel(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *Level
		state   models.SavedLevel
		wantErr bool
	}{
		{
			name: "Restore valid level state",
			setup: func() *Level {
				return NewLevel()
			},
			state: models.SavedLevel{
				Number: 0,
				Width:  20,
				Height: 30,
			},
			wantErr: false,
		},
		{
			name: "Restore invalid dimensions",
			setup: func() *Level {
				return NewLevel()
			},
			state: models.SavedLevel{
				Number: 0,
				Width:  0,
				Height: -1,
			},
			wantErr: true,
		},
		{
			name: "Restore non-existent level",
			setup: func() *Level {
				return NewLevel()
			},
			state: models.SavedLevel{
				Number: 99,
				Width:  20,
				Height: 30,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := tt.setup()

			err := lc.RestoreLevel(tt.state)
			if (err != nil) != tt.wantErr {
				t.Errorf("RestoreLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_InitOrResizeLevel(t *testing.T) {
	tests := []struct {
		name    string
		setup   func() *Level
		width   int
		height  int
		wantErr bool
	}{
		{
			name: "Initialize or resize a valid level",
			setup: func() *Level {
				lc := NewLevel()
				lc.SetLevel(lc.GetLevels()[0])
				return lc
			},
			width:   20,
			height:  20,
			wantErr: false,
		},
		{
			name: "Error on resizing with no current level",
			setup: func() *Level {
				return NewLevel()
			},
			width:   20,
			height:  20,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := tt.setup()

			err := lc.InitOrResizeLevel(tt.width, tt.height)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitOrResizeLevel() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
