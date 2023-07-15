package repository

import (
	"errors"
	
	"github.com/WillyWilsen/Delos-Task-Assignment.git/model"
)

// MockPondRepository is a mock implementation of the PondRepository interface
type MockPondRepository struct {
	ponds      map[int]*model.Pond
}

func NewMockPondRepository() *MockPondRepository {
	return &MockPondRepository{
		ponds:      make(map[int]*model.Pond),
	}
}

func (m *MockPondRepository) Create(pond *model.Pond) error {
	pond.ID = len(m.ponds) + 1
	m.ponds[pond.ID] = pond
	return nil
}

func (m *MockPondRepository) GetByName(name string) (*model.Pond, error) {
	for _, pond := range m.ponds {
		if pond.Name == name {
			return pond, nil
		}
	}
	return nil, nil
}

func (m *MockPondRepository) Get() ([]model.Pond, error) {
	ponds := make([]model.Pond, 0, len(m.ponds))
	for _, pond := range m.ponds {
		ponds = append(ponds, *pond)
	}
	return ponds, nil
}

func (m *MockPondRepository) GetById(id int) (*model.Pond, error) {
	pond, ok := m.ponds[id]
	if !ok {
		return nil, nil
	}
	return pond, nil
}

func (m *MockPondRepository) Update(id int, pond *model.Pond) error {
	existingPond, ok := m.ponds[id]
	if !ok {
		return errors.New("Pond not found")
	}

	pond.ID = existingPond.ID
	m.ponds[id] = pond
	return nil
}

func (m *MockPondRepository) Delete(pond *model.Pond) error {
	delete(m.ponds, pond.ID)
	return nil
}