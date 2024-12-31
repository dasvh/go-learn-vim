package game

// Stats represent the statistics of a game session
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
