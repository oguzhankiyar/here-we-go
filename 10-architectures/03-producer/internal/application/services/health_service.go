package services

import (
	"time"

	"github.com/mackerelio/go-osstat/cpu"
	"github.com/mackerelio/go-osstat/memory"

	healthModels "producer-sample/internal/core/models"
	configModels "producer-sample/internal/infrastructure/config/models"
)

type HealthService struct {
	config configModels.AppConfig
}

func NewHealthService(config configModels.AppConfig) *HealthService {
	return &HealthService{
		config: config,
	}
}

func (s *HealthService) Get() *healthModels.HealthModel {
	return &healthModels.HealthModel{
		App:     s.config.Name,
		State:   "healthy",
		Message: "i'm healthy",
		CPU:     GetCPUInfo(),
		Memory:  GetMemoryInfo(),
		Time:    time.Now().Unix(),
	}
}

func GetCPUInfo() *healthModels.CPUModel {
	before, err := cpu.Get()
	if err != nil {
		return nil
	}

	time.Sleep(time.Duration(1) * time.Second)

	after, err := cpu.Get()
	if err != nil {
		return nil
	}

	total := float64(after.Total - before.Total)

	return &healthModels.CPUModel{
		User:   float64(after.User-before.User) / total * 100,
		System: float64(after.System-before.System) / total * 100,
		Idle:   float64(after.Idle-before.Idle) / total * 100,
	}
}

func GetMemoryInfo() *healthModels.MemoryModel {
	memory, err := memory.Get()
	if err != nil {
		return nil
	}

	return &healthModels.MemoryModel{
		Total:  memory.Total,
		Used:   memory.Used,
		Cached: memory.Cached,
		Free:   memory.Free,
	}
}
