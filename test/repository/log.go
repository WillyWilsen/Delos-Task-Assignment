package repository

import (
	"github.com/WillyWilsen/Delos-Task-Assignment.git/model"
)

// MockLogRepository is a mock implementation of the LogRepository interface
type MockLogRepository struct {
	logs []*model.Log
}

func NewMockLogRepository() *MockLogRepository {
	return &MockLogRepository{
		logs: make([]*model.Log, 0),
	}
}

func (m *MockLogRepository) Create(log *model.Log) error {
	m.logs = append(m.logs, log)
	return nil
}

func (m *MockLogRepository) GetDistinctEndpoints() ([]string, error) {
	endpoints := make([]string, 0)
	visitedEndpoints := make(map[string]bool)
	for _, log := range m.logs {
		if !visitedEndpoints[log.Endpoint] {
			endpoints = append(endpoints, log.Endpoint)
			visitedEndpoints[log.Endpoint] = true
		}
	}
	return endpoints, nil
}

func (m *MockLogRepository) GetEndpointStatistics(endpoint string) (*model.EndpointStatistics, error) {
	count := 0
	uniqueUserAgents := make(map[string]bool)
	for _, log := range m.logs {
		if log.Endpoint == endpoint {
			count++
			uniqueUserAgents[log.UserAgent] = true
		}
	}
	statistics := &model.EndpointStatistics{
		Count:             int64(count),
		UniqueUserAgent:   int64(len(uniqueUserAgents)),
	}
	return statistics, nil
}