package repository

import (
    "gorm.io/gorm"
    "github.com/WillyWilsen/Delos-Task-Assignment.git/model"
)

type LogRepository interface {
    Create(log *model.Log) error
    GetDistinctEndpoints() ([]string, error)
    GetEndpointStatistics(endpoint string) (*model.EndpointStatistics, error)
}

type LogRepositoryImpl struct {
    db *gorm.DB
}

func NewLogRepository(db *gorm.DB) LogRepository {
    return &LogRepositoryImpl{
        db: db,
    }
}

func (r *LogRepositoryImpl) Create(log *model.Log) error {
    return r.db.Create(log).Error
}

func (r *LogRepositoryImpl) GetDistinctEndpoints() ([]string, error) {
    var endpoints []string
    if err := r.db.Table("logs").Distinct("endpoint").Pluck("endpoint", &endpoints).Error; err != nil {
        return nil, err
    }
    return endpoints, nil
}

func (r *LogRepositoryImpl) GetEndpointStatistics(endpoint string) (*model.EndpointStatistics, error) {
    var endpointStatistics *model.EndpointStatistics
    query := `
		SELECT 
			COUNT(*) AS count,
			COUNT(DISTINCT user_agent) AS unique_user_agent
		FROM 
			logs
		WHERE 
			endpoint = ?`
    if err := r.db.Raw(query, endpoint).Scan(&endpointStatistics).Error; err != nil {
        return nil, err
    }
    return endpointStatistics, nil
}