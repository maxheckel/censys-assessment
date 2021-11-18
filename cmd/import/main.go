package main

import (
	"fmt"
	"github.com/maxheckel/censys-assessment/internal/config"
	"github.com/maxheckel/censys-assessment/internal/domain"
	"github.com/oschwald/maxminddb-golang"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"

	"log"
	"path/filepath"
)

func main() {
	cfg, err := config.Load("CENSYS")
	if err != nil {
		log.Fatal(err)
	}
	db, err := getDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.IP{})
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(fmt.Sprintf(`COPY ips(network,geoname_id,registered_country_geoname_id,represented_country_geoname_id,is_anonymous_proxy,is_satellite_provider,postal_code,latitude,longitude,accuracy_radius)
FROM '%s'
DELIMITER ','
CSV HEADER;`,  "/var/lib/postgresql/data/GeoLite2-City-Blocks-IPv4.csv"))
	//importFile(db)

}

func getDB(config *config.Config) (*gorm.DB, error){
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.DBHost, config.DBUser, config.DBPassword, config.DBName, config.DBPort, config.DBSSLMode)
	count := 0
	var db *gorm.DB
	var err error
	// Very basic sleep / retry
	for count < 5{
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			return db, err
		}
		count++
		time.Sleep(2 * time.Second)
	}
	return db, err
}

func importFile(db *gorm.DB) {
	fileName, err := filepath.Abs("./cmd/import/data/GeoLite2-City.mmdb")
	fmt.Println(fileName)
	file, err := maxminddb.Open(fileName)

	if err != nil {
		panic(err)
	}

	record := domain.MaxMindRow{}
	if countsMatch(db, file) {
		return
	}
	err = db.Exec("TRUNCATE TABLE ips").Error
	if err != nil {
		panic(err)
	}
	networks := file.Networks(maxminddb.SkipAliasedNetworks)
	count := 0
	for networks.Next() {
		subnet, err := networks.Network(&record)
		subnet.IP.String()
		if err != nil {
			log.Fatal(err)
		}
		//row := domain.IP{
		//	Latitude:  record.Location.Lat,
		//	Longitude: record.Location.Lng,
		//	Address:   subnet.IP.String(),
		//}
		//db.Create(&row)
		fmt.Printf("%d: %s created\n", count, subnet.IP.String())
		count++
	}
	if networks.Err() != nil {
		log.Fatal(networks.Err())
	}
}

func countsMatch(db *gorm.DB, file *maxminddb.Reader) bool {
	type Count struct {
		Count uint
	}
	var res Count
	db.Raw("SELECT count(*) FROM ips").Scan(&res)
	if res.Count == file.Metadata.NodeCount {
		return true
	}
	return false
}
