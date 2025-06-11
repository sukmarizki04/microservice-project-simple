package main

import (
	"database/sql"
	"net/http"
	"product-service/internal/application"
	"product-service/internal/infrastructure/repository"
	web "product-service/internal/interfaces"

	_ "github.com/mattn/go-sqlite3" // Required SQLite driver
)

func main() {
	// repo := &repository.MySQLProductRepo{}
	db, err := sql.Open("sqlite3", "file:product.db?cache=shared&mode=rwc")
	if err != nil {
		panic(err)
	}
	// ✅ One-time schema injection
	if err := repository.InitSQLiteSchema(db); err != nil {
		panic(err)
	}

	repo := &repository.SQLiteProductRepo{DB: db}
	appService := &application.ProductService{Repo: repo}
	handler := &web.ProductHandler{Service: appService}
	router := web.NewRouter(handler)

	http.ListenAndServe(":8080", router)
}
