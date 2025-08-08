package novalagunggolang

import (
	"encoding/json"
	"net/http"
)

type student struct {
	ID   string
	Name string
	Age  int
}

var data = []student{
	student{"EOO1", "SUKMA", 22},
	student{"e002", "EDO", 45},
	student{"EOO1", "SUKMA", 22},
	student{"e002", "EDO", 45},
}

func user(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method == "GET" {
		var result, err = json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
	http.Error(w, " ", http.StatusBadRequest)
}
