package repository

import (
	"errors"
	
	"github.com/WillyWilsen/Delos-Task-Assignment.git/model"
)

// MockFarmRepository is a mock implementation of the FarmRepository interface
type MockFarmRepository struct {
	farms      map[int]*model.Farm
}

func NewMockFarmRepository() *MockFarmRepository {
	return &MockFarmRepository{
		farms:      make(map[int]*model.Farm),
	}
}

func (m *MockFarmRepository) Create(farm *model.Farm) error {
	farm.ID = len(m.farms) + 1
	m.farms[farm.ID] = farm
	return nil
}

func (m *MockFarmRepository) GetByName(name string) (*model.Farm, error) {
	for _, farm := range m.farms {
		if farm.Name == name {
			return farm, nil
		}
	}
	return nil, nil
}

func (m *MockFarmRepository) Get() ([]model.Farm, error) {
	farms := make([]model.Farm, 0, len(m.farms))
	for _, farm := range m.farms {
		farms = append(farms, *farm)
	}
	return farms, nil
}

func (m *MockFarmRepository) GetById(id int) (*model.Farm, error) {
	farm, ok := m.farms[id]
	if !ok {
		return nil, nil
	}
	return farm, nil
}

func (m *MockFarmRepository) Update(id int, farm *model.Farm) error {
	existingFarm, ok := m.farms[id]
	if !ok {
		return errors.New("Farm not found")
	}

	farm.ID = existingFarm.ID
	m.farms[id] = farm
	return nil
}

func (m *MockFarmRepository) Delete(farm *model.Farm) error {
	delete(m.farms, farm.ID)
	return nil
}