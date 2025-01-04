package models

// Stats represent the statistics of a game
type Stats struct {
	KeyPresses      map[string]int `json:"key_presses"`
	TotalKeystrokes int            `json:"total_keystrokes"`
	TimeElapsed     int            `json:"time_elapsed"`
}

// NewStats creates a new Stats instance
func NewStats() *Stats {
	return &Stats{
		KeyPresses:      make(map[string]int),
		TotalKeystrokes: 0,
		TimeElapsed:     0,
	}
}

// RegisterKey increments the count for a given key if allowed
func (s *Stats) RegisterKey(key string, allowed bool) {
	if allowed {
		s.TotalKeystrokes++
		s.KeyPresses[key]++
	}
}

// IncrementTime increments the time counter
func (s *Stats) IncrementTime() {
	s.TimeElapsed++
}

// Reset clears all statistics
func (s *Stats) Reset() {
	s.KeyPresses = make(map[string]int)
	s.TotalKeystrokes = 0
}

// LifetimeStats represents aggregated statistics for all games
type LifetimeStats struct {
	TotalKeystrokes int            `json:"total_keystrokes"`
	TotalPlaytime   int            `json:"total_playtime"`
	TotalGames      int            `json:"total_games"`
	KeyPresses      map[string]int `json:"key_presses"`
}

// NewLifetimeStats initializes an empty LifetimeStats
func NewLifetimeStats() *LifetimeStats {
	return &LifetimeStats{
		TotalKeystrokes: 0,
		TotalPlaytime:   0,
		TotalGames:      0,
		KeyPresses:      make(map[string]int),
	}
}

// Merge aggregates stats from another Stats instance
func (ls *LifetimeStats) Merge(stats Stats) {
	ls.TotalKeystrokes += stats.TotalKeystrokes
	ls.TotalPlaytime += stats.TimeElapsed

	for key, count := range stats.KeyPresses {
		ls.KeyPresses[key] += count
	}
}
