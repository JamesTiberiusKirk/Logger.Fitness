package types

// Workout is for a collection of exercises
type Workout struct {
	ID        string       `json:"id"`
	Execises  [][]Exercise `json:"execises"`
	StartTime int64        `json:"start_time"`
	EndTime   int64        `json:"end_time"`
	Notes     string       `json:"notes"`
	Title     string       `json:"title"`
}
