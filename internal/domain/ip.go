package domain

import "gorm.io/gorm"

type IP struct {
	gorm.Model
	Network string
	Geoname_id string
	Registered_country_geoname_id string
	Represented_country_geoname_id string
	Is_anonymous_proxy string
	Is_satellite_provider string
	Postal_code string
	Latitude string
	Longitude string
	Accuracy_radius string
}
