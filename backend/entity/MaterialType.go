package entity

// Material model info
type Material struct {
	ID            uint   `gorm:"primaryKey" json:"id" `
	Name          string `json:"name" `
	CodeType      string `json:"codeType" `
	UnitValue     int    `json:"unitValue" `
	Qty           int    `json:"qty" `
	LocationRefer uint   `json:"locationRefer" `
	VehicleRefer  uint   `json:"vehicleRefer" `
}
