package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"rahuljuliato.com/go/museum/data"
)

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// api/exhibitions?id=2
	// id == "2" ([]string) why? because it could be id=1,2,3,4
	id := r.URL.Query()["id"]

	if id != nil {
		finalId, err := strconv.Atoi(id[0])
		if err == nil && finalId < len(data.GetAll()) {

			json.NewEncoder(w).Encode(data.GetAll()[finalId])

		} else {
			http.Error(w, "Invalid exhibition id", http.StatusBadRequest)
		}
	} else {
		json.NewEncoder(w).Encode(data.GetAll())
	}

}
