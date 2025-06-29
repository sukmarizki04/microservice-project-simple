package main

import (
	"flag"
	_ "fmt"
	"html/template"
	"log"
	"net/http"
)

var add = flag.String("addr", "17:18", "http service addres")
var templ = template.Must(template.New("qr").Parse(template.JSEscaper()))

func main() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(QR))
	err := http.ListenAndServe(*add, nil)
	if err != nil {
		log.Fatal("ListenServe:", err)
	}
}

func QR(w http.ResponseWriter, req *http.Request) {
	templ.Execute(w, req.FormValue("s"))

}

const templateStr = `<html>
<head>
<title>QR Link Generator</title>
</head>
<body>
{{if .}}
<img src="http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl={{.}}" />
<br>
{{.}}
<br>
<br>
{{end}}
<form action="/" name=f method="GET">
    <input maxLength=1024 size=70 name=s value="" title="Text to QR Encode">
    <input type=submit value="Show QR" name=qr>
</form>
</body>
</html>`
