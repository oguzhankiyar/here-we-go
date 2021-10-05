package models

type HealthModel struct {
	App     string       `json:"app"`
	State   string       `json:"state"`
	Message string       `json:"message"`
	CPU     *CPUModel    `json:"cpu"`
	Memory  *MemoryModel `json:"memory"`
	Time    int64        `json:"time"`
}

type CPUModel struct {
	User   float64 `json:"user"`
	System float64 `json:"system"`
	Idle   float64 `json:"idle"`
}

type MemoryModel struct {
	Total  uint64 `json:"total"`
	Used   uint64 `json:"used"`
	Cached uint64 `json:"cached"`
	Free   uint64 `json:"free"`
}
