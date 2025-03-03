package models

import (
	"reflect"
	"testing"
)

func TestStats_RegisterKey(t *testing.T) {
	tests := []struct {
		name           string
		initialStats   Stats
		key            string
		allowed        bool
		wantKeystrokes int
		wantKeyPresses map[string]int
	}{
		{
			name: "register allowed key",
			initialStats: Stats{
				KeyPresses:      map[string]int{},
				TotalKeystrokes: 0,
			},
			key:            "j",
			allowed:        true,
			wantKeystrokes: 1,
			wantKeyPresses: map[string]int{"j": 1},
		},
		{
			name: "register same key multiple times",
			initialStats: Stats{
				KeyPresses:      map[string]int{"j": 1},
				TotalKeystrokes: 1,
			},
			key:            "j",
			allowed:        true,
			wantKeystrokes: 2,
			wantKeyPresses: map[string]int{"j": 2},
		},
		{
			name: "register disallowed key",
			initialStats: Stats{
				KeyPresses:      map[string]int{"j": 2},
				TotalKeystrokes: 2,
			},
			key:            "x",
			allowed:        false,
			wantKeystrokes: 2,
			wantKeyPresses: map[string]int{"j": 2},
		},
		{
			name: "register different allowed key",
			initialStats: Stats{
				KeyPresses:      map[string]int{"j": 2},
				TotalKeystrokes: 2,
			},
			key:            "k",
			allowed:        true,
			wantKeystrokes: 3,
			wantKeyPresses: map[string]int{"j": 2, "k": 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats := tt.initialStats
			stats.RegisterKey(tt.key, tt.allowed)

			if stats.TotalKeystrokes != tt.wantKeystrokes {
				t.Errorf("expected TotalKeystrokes to be %d, got %d", tt.wantKeystrokes, stats.TotalKeystrokes)
			}

			if !reflect.DeepEqual(stats.KeyPresses, tt.wantKeyPresses) {
				t.Errorf("expected KeyPresses to be %v, got %v", tt.wantKeyPresses, stats.KeyPresses)
			}
		})
	}
}

func TestStats_IncrementTime(t *testing.T) {
	tests := []struct {
		name            string
		initialTime     int
		increments      int
		wantTimeElapsed int
	}{
		{
			name:            "from zero",
			initialTime:     0,
			increments:      1,
			wantTimeElapsed: 1,
		},
		{
			name:            "multiple increments",
			initialTime:     2,
			increments:      3,
			wantTimeElapsed: 5,
		},
		{
			name:            "no increments",
			initialTime:     10,
			increments:      0,
			wantTimeElapsed: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats := Stats{TimeElapsed: tt.initialTime}

			for i := 0; i < tt.increments; i++ {
				stats.IncrementTime()
			}

			if stats.TimeElapsed != tt.wantTimeElapsed {
				t.Errorf("expected TimeElapsed to be %d, got %d", tt.wantTimeElapsed, stats.TimeElapsed)
			}
		})
	}
}

func TestStats_Reset(t *testing.T) {
	tests := []struct {
		name         string
		initialStats Stats
		wantStats    Stats
	}{
		{
			name: "reset with data",
			initialStats: Stats{
				KeyPresses:      map[string]int{"j": 2, "k": 1},
				TotalKeystrokes: 3,
				TimeElapsed:     10,
			},
			wantStats: Stats{
				KeyPresses:      map[string]int{},
				TotalKeystrokes: 0,
				TimeElapsed:     10,
			},
		},
		{
			name: "reset empty stats",
			initialStats: Stats{
				KeyPresses:      map[string]int{},
				TotalKeystrokes: 0,
				TimeElapsed:     5,
			},
			wantStats: Stats{
				KeyPresses:      map[string]int{},
				TotalKeystrokes: 0,
				TimeElapsed:     5,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stats := tt.initialStats
			stats.Reset()

			if stats.TotalKeystrokes != tt.wantStats.TotalKeystrokes {
				t.Errorf("expected TotalKeystrokes to be %d, got %d", tt.wantStats.TotalKeystrokes, stats.TotalKeystrokes)
			}

			if stats.TimeElapsed != tt.wantStats.TimeElapsed {
				t.Errorf("expected TimeElapsed to be %d, got %d", tt.wantStats.TimeElapsed, stats.TimeElapsed)
			}

			if !reflect.DeepEqual(stats.KeyPresses, tt.wantStats.KeyPresses) {
				t.Errorf("expected KeyPresses to be %v, got %v", tt.wantStats.KeyPresses, stats.KeyPresses)
			}
		})
	}
}

func TestLifetimeStats_Merge(t *testing.T) {
	lifetime := NewLifetimeStats()

	stats1 := NewStats()
	stats1.RegisterKey("k", true)
	stats1.RegisterKey("j", true)
	stats1.RegisterKey("k", true)
	stats1.IncrementTime()
	stats1.IncrementTime()

	lifetime.Merge(*stats1)
	if lifetime.TotalKeystrokes != 3 {
		t.Errorf("expected TotalKeystrokes to be 3, got %d", lifetime.TotalKeystrokes)
	}
	if lifetime.TotalPlaytime != 2 {
		t.Errorf("expected TotalPlaytime to be 2, got %d", lifetime.TotalPlaytime)
	}
	if lifetime.KeyPresses["k"] != 2 {
		t.Errorf("expected KeyPresses['k'] to be 2, got %d", lifetime.KeyPresses["k"])
	}
	if lifetime.KeyPresses["j"] != 1 {
		t.Errorf("expected KeyPresses['j'] to be 1, got %d", lifetime.KeyPresses["j"])
	}

	stats2 := NewStats()
	stats2.RegisterKey("h", true)
	stats2.RegisterKey("k", true)
	stats2.IncrementTime()
	lifetime.Merge(*stats2)

	if lifetime.TotalKeystrokes != 5 {
		t.Errorf("expected TotalKeystrokes to be 5, got %d", lifetime.TotalKeystrokes)
	}
	if lifetime.TotalPlaytime != 3 {
		t.Errorf("expected TotalPlaytime to be 3, got %d", lifetime.TotalPlaytime)
	}

	expectedKeyPresses := map[string]int{
		"k": 3,
		"j": 1,
		"h": 1,
	}
	if !reflect.DeepEqual(lifetime.KeyPresses, expectedKeyPresses) {
		t.Errorf("expected KeyPresses to be %v, got %v", expectedKeyPresses, lifetime.KeyPresses)
	}
}
