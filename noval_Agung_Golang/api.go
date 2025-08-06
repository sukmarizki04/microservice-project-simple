package novalagunggolang

import "net/http"

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

func user(w http.ResponseWriter, R *http.Request) {

}
