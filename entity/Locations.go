package entity

type Location struct {
	ID      uint `gorm:"primaryKey"`
	Name    string
	Address string
}

func GetAllLocations() []Location {
	var locations []Location
	GetConnection().Find(&locations)
	return locations
}

func GetLocationById(id int) Location {
	var loc Location
	GetConnection().First(&loc, id)
	return loc
}

func SaveLocation(loc Location) {
	GetConnection().Create(&loc)
}

func DeleteLocation(id int) {
	GetConnection().Delete(&Location{}, id)
}
