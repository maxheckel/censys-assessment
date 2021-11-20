package main

import (
	"flag"
	"fmt"
	"github.com/maxheckel/censys-assessment/internal/config"
	"github.com/maxheckel/censys-assessment/internal/domain"
	"github.com/maxheckel/censys-assessment/internal/server"
	"github.com/oschwald/maxminddb-golang"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
	"path/filepath"
	"time"
)

func main() {
	iteratively := flag.Bool("iteratively", false, "whether to run the import iteratively or not")
	flag.Parse()
	cfg, err := config.Load("CENSYS")
	if err != nil {
		log.Fatal(err)
	}
	db, err := server.GetDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&domain.IP{})
	if err != nil {
		log.Fatal(err)
	}

	start := time.Now()
	if *iteratively {
		err = iterativeInsert(db)
	} else {
		err = bulkInsert(db)
		if err != nil {
			err = errors.Wrap(err, "You might not have copied your GeoLite2-City-Blocks-IPv4.csv file into the dblocal folder or otherwise onto the server where the DB is hosted if not running locally!")
		}
	}
	if err != nil {
		log.Fatal(err)
	}

	elapsed := time.Since(start)
	log.Printf("Import took %s", elapsed)
}

func bulkInsert(db *gorm.DB) error {
	log.Println("Beginning import, this could take a while!")

	return db.Exec(fmt.Sprintf(`COPY ips(network,geoname_id,registered_country_geoname_id,represented_country_geoname_id,is_anonymous_proxy,is_satellite_provider,postal_code,latitude,longitude,accuracy_radius)
	FROM '%s'
	DELIMITER ','
	CSV HEADER;`, "/var/lib/postgresql/data/GeoLite2-City-Blocks-IPv4.csv")).Error

}

// Alternative way to insert the file line by line using the mmdb file type
func iterativeInsert(db *gorm.DB) error{
	fileName, err := filepath.Abs("./cmd/import/data/GeoLite2-City.mmdb")
	fmt.Println(fileName)
	file, err := maxminddb.Open(fileName)

	if err != nil {
		return err
	}

	record := domain.MaxMindRow{}
	if countsMatch(db, file) {
		return nil
	}
	err = db.Exec("TRUNCATE TABLE ips").Error
	if err != nil {
		return err
	}
	networks := file.Networks(maxminddb.SkipAliasedNetworks)
	count := 0
	var rows []domain.IP
	for networks.Next() {
		subnet, err := networks.Network(&record)
		subnet.IP.String()
		if err != nil {
			log.Fatal(err)
		}
		row := domain.IP{
			Latitude:  fmt.Sprintf("%f", record.Location.Lat),
			Longitude: fmt.Sprintf("%f", record.Location.Lng),
			Network:   subnet.IP.String(),
		}
		rows = append(rows, row)
		if len(rows) >= 1000 {
			db.Create(&rows)
			rows = []domain.IP{}
			fmt.Printf("%d: created\n", count)
		}

		count++
	}
	return networks.Err()
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
