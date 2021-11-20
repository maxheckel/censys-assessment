package domain

type IP struct {
	Network                        string `gorm:"index"`
	Geoname_id                     string
	Registered_country_geoname_id  string
	Represented_country_geoname_id string
	Is_anonymous_proxy             string
	Is_satellite_provider          string
	Postal_code                    string
	Latitude                       string
	Longitude                      string
	Accuracy_radius                string
}
