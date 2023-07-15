package repository

import (
    "gorm.io/gorm"
    "github.com/WillyWilsen/Delos-Task-Assignment.git/model"
)

type FarmRepository interface {
    GetByName(name string) (*model.Farm, error)
    Create(farm *model.Farm) error
    Get() ([]model.Farm, error)
    GetById(id int) (*model.Farm, error)
    Update(id int, farm *model.Farm) error
    Delete(farm *model.Farm) error
}

type FarmRepositoryImpl struct {
    db *gorm.DB
}

func NewFarmRepository(db *gorm.DB) FarmRepository {
    return &FarmRepositoryImpl{
        db: db,
    }
}

func (r *FarmRepositoryImpl) GetByName(name string) (*model.Farm, error) {
    var farm model.Farm
    if err := r.db.Table("farms").Where("name = ?", name).First(&farm).Error; err != nil {
        return nil, err
    }
    return &farm, nil
}

func (r *FarmRepositoryImpl) Create(farm *model.Farm) error {
    return r.db.Create(farm).Error
}

func (r *FarmRepositoryImpl) Get() ([]model.Farm, error) {
    var farm []model.Farm
    if err := r.db.Table("farms").Scan(&farm).Error; err != nil {
        return nil, err
    }
    return farm, nil
}

func (r *FarmRepositoryImpl) GetById(id int) (*model.Farm, error) {
    var farm *model.Farm
    if err := r.db.Table("farms").Where("id = ?", id).First(&farm).Error; err != nil {
        return nil, err
    }
    return farm, nil
}

func (r *FarmRepositoryImpl) Update(id int, farm *model.Farm) error {
    return r.db.Table("farms").Where("id = ?", id).Updates(model.Pond{Name: farm.Name}).Error
}

func (r *FarmRepositoryImpl) Delete(farm *model.Farm) error {
    return r.db.Table("farms").Delete(&farm).Error
}