package entity

// Location model info
// @Description Simple object
type Location struct {
	ID           uint   `gorm:"primaryKey" json:"id" `
	Name         string `json:"name" `
	Address      string `json:"address" `
	HouseNr      int    `json:"houseNr" `
	CompanyRefer uint   `json:"companyRefer" `
}
