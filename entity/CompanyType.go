package entity

type Company struct {
	ID        uint `gorm:"primaryKey"`
	Name      string
	Gln       int
	Locations []Location `gorm:"foreignKey:CompanyRefer;references:ID" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
