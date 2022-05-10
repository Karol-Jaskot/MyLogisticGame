package entity

// Material model info
type Material struct {
	ID        uint   `gorm:"primaryKey" json:"id" `
	Name      string `json:"name" `
	CodeType  string `json:"code_type" `
	UnitValue int    `json:"unit_value" `
	Qty       int    `json:"qty" `
}
