package domain

type MaxMindRow struct {
	Location struct {
		Lat float64 `maxminddb:"latitude" json:"lat"`
		Lng float64 `maxminddb:"longitude" json:"lng"`
	} `maxminddb:"location" json:"location"`
}
