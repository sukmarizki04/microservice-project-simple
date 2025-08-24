package novalagunggolang

import (
	"encoding/json"
	"fmt"
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

type T struct {
	name  string // name of the object
	value int    // its value
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

func userq(w http.ResponseWriter, r *http.Request) {

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
func name(age int) {
	if age == 0 {
		return fmt.Append()
	}
}
