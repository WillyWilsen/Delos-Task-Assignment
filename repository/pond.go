package repository

import (
    "gorm.io/gorm"
    "github.com/WillyWilsen/Delos-Task-Assignment.git/model"
)

type PondRepository interface {
    GetByName(name string) (*model.Pond, error)
    Create(pond *model.Pond) error
    Get() ([]model.Pond, error)
    GetById(id int) (*model.Pond, error)
    Update(id int, pond *model.Pond) error
    Delete(pond *model.Pond) error
}

type PondRepositoryImpl struct {
    db *gorm.DB
}

func NewPondRepository(db *gorm.DB) PondRepository {
    return &PondRepositoryImpl{
        db: db,
    }
}

func (r *PondRepositoryImpl) GetByName(name string) (*model.Pond, error) {
    var pond model.Pond
    if err := r.db.Table("ponds").Where("name = ?", name).First(&pond).Error; err != nil {
        return nil, err
    }
    return &pond, nil
}

func (r *PondRepositoryImpl) Create(pond *model.Pond) error {
    return r.db.Create(pond).Error
}

func (r *PondRepositoryImpl) Get() ([]model.Pond, error) {
    var pond []model.Pond
    if err := r.db.Table("ponds").Scan(&pond).Error; err != nil {
        return nil, err
    }
    return pond, nil
}

func (r *PondRepositoryImpl) GetById(id int) (*model.Pond, error) {
    var pond *model.Pond
    if err := r.db.Table("ponds").Where("id = ?", id).First(&pond).Error; err != nil {
        return nil, err
    }
    return pond, nil
}

func (r *PondRepositoryImpl) Update(id int, pond *model.Pond) error {
    return r.db.Table("ponds").Where("id = ?", id).Updates(model.Pond{Name: pond.Name, FarmID: pond.FarmID}).Error
}

func (r *PondRepositoryImpl) Delete(pond *model.Pond) error {
    return r.db.Table("ponds").Delete(&pond).Error
}