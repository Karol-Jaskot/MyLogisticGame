package entity

// Vehicle model info
type Vehicle struct {
	ID            uint       `gorm:"primaryKey" json:"id" `
	Name          string     `json:"name" `
	Type          string     `json:"type" `
	PurchasePrice int        `json:"purchasePrice" `
	Materials     []Material `gorm:"foreignKey:VehicleRefer;references:ID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
