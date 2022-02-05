package database

import (
	"github.com/jinzhu/gorm"
	"github.com/pallab-gogoi/covid_vaccine/model"
)

func GetData(db *gorm.DB) ([]model.Details, error) {
	data := []model.Details{}
	quary := db.Select("details.*")
	err := quary.Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}
func SaveData(db *gorm.DB, data *model.Details) error {
	err := db.Save(data).Error
	if err != nil {
		return err
	}
	return nil
}
