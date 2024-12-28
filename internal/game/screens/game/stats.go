package game

// Stats represent the statistics of a game session
type Stats struct {
	TotalKeystrokes int
	KeyPresses      map[string]int
	Time            int
}

// NewStats creates a new Stats instance
func NewStats() *Stats {
	return &Stats{
		TotalKeystrokes: 0,
		KeyPresses:      make(map[string]int),
		Time:            0,
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
	s.Time++
}

// Reset clears all statistics
func (s *Stats) Reset() {
	s.TotalKeystrokes = 0
	s.KeyPresses = make(map[string]int)
}
