package domain

type MaxMindRow struct {
	Location struct {
		Lat float64 `maxminddb:"latitude"`
		Lng float64 `maxminddb:"longitude"`
	} `maxminddb:"location"`
}