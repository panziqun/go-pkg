package address

type Country struct {
	ID          int    `gorm:"id" json:"id"`
	CountryCode string `gorm:"country_code" json:"country_code"`
	CountryName string `gorm:"country_name" json:"country_name"`
}

type Province struct {
	ID           int    `gorm:"id" json:"id"`
	ProvinceCode string `gorm:"province_code" json:"province_code"`
	ProvinceName string `gorm:"province_name" json:"province_name"`
}

type City struct {
	ID           int    `gorm:"id" json:"id"`
	CityCode     string `gorm:"city_code" json:"city_code"`
	CityName     string `gorm:"city_name" json:"city_name"`
	ProvinceCode string `gorm:"province_code" json:"province_code"`
}

type District struct {
	ID           int    `gorm:"id" json:"id"`
	DistrictCode string `gorm:"district_code" json:"district_code"`
	DistrictName string `gorm:"district_name" json:"district_name"`
	CityCode     string `gorm:"city_code" json:"city_code"`
}
