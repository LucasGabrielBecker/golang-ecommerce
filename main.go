package main

import (
	"net/http"

	"github.com/LucasGabrielBecker/golang-ecommerce/database"
	"github.com/LucasGabrielBecker/golang-ecommerce/routes"
)

func main() {
	database.ConnecToDB()
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
