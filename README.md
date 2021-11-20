# Internet.Protocol.Lookup.Tool/24

This standalone app is inteneded to use the databases provided by [Max Mind](https://maxmind.com/) to lookup latitude and longitude based on IP addresses

## Local development


### Requirements
* Docker / Docker-compose
* Golang 1.15
* Vue CLI
* Postgres

### Starting your local dev server
Simply run `docker-compose up` and it will take care of it.  After that you'll need to import the data as outlined below.

The API's url is localhost:8080, the UI's url is localhost:4200

#### Importing data
Within the `cmd/import/data` folder you must a file downloaded from Max Mind DB.  It can be [downloaded here](https://dev.maxmind.com/geoip/geoip2/geolite2/).  The file should be named GeoLite2-City.mmdb and placed in the `cmd/import/data` folder.


### Thoughts and Notes
tldr; I don't know how subnets work

#### A note on the persistence layer:
The only file that is actually required is the `GeoLite2-City.mmdb` file as the lookups are performed directly against that using the [oschwald/maxminddb-golang](https://github.com/oschwald/maxminddb-golang) package in golang.  For the sake of the assessment I also left in my original attempt at importing the data into a postgres DB.  If you would like to import the data into a postgres database, download it in CSV format and place it in the `dblocal` folder. While it's not required you can import the IP DB into postgres using the following command: ` CENSYS_DB_HOST=localhost CENSYS_DB_PORT=54321 go run cmd/import/main.go`  

The final app didn't end up using the postgres DB in favor of looking everything up in the `.mmdb` file but it might be interesting to see how I did that.


#### A Note on testing
While I would have liked to write some unit tests for this I was closing in on 5 hours of coding and did not have the time.  I did at least write testable code by adding interface/mocks where appropriate.

#### Error Handling
Error handling kind of sucks right now in the UI, it could definitely be improved.  Same for the API.

#### A Note on hosting SPAs
I included a handler to host the entire app using golang, but I don't think you'd want a docker compose file with just a golang server so I didn't end up using it.  If this were to be deployed ever though, it could be useful.
