package handlers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/maxheckel/censys-assessment/internal/domain"
	"github.com/maxheckel/censys-assessment/internal/responses"
	"github.com/oschwald/maxminddb-golang"
	"gorm.io/gorm"
	"net"
	"net/http"
	"path/filepath"
)

func (h *Handlers) GetIPDetails(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// I don't know enough about networking to figure out how to do this with the imported DB, this seems like I could
	// optimize this but it also seems quite fast
	fileName, _ := filepath.Abs("./cmd/import/data/GeoLite2-City.mmdb")

	file, _ := maxminddb.Open(fileName)

	ip := net.ParseIP(params["address"])
	var record domain.MaxMindRow// Or any appropriate struct
	err := file.Lookup(ip, &record)
	if err != nil {
		responseError := responses.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
		w.WriteHeader(responseError.Code)
		json.NewEncoder(w).Encode(responseError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(record)
}

func handleSqlError(w http.ResponseWriter, err error) {
	var errorResponse responses.Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		errorResponse = responses.Error{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		}

	} else {
		errorResponse = responses.Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	w.WriteHeader(errorResponse.Code)
	json.NewEncoder(w).Encode(errorResponse)
}

