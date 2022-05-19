package entity

import "MyLogisticGame/backend/data/materials"

// Location model info
// @Description Simple object
type Location struct {
	ID           uint                 `gorm:"primaryKey" json:"id" `
	Name         string               `json:"name" `
	Address      string               `json:"address" `
	HouseNr      int                  `json:"houseNr" `
	CompanyRefer uint                 `json:"companyRefer" `
	Materials    []materials.Material `gorm:"foreignKey:LocationRefer;references:ID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
