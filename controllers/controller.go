package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/LucasGabrielBecker/golang-ecommerce/database"
	"github.com/LucasGabrielBecker/golang-ecommerce/models"
)

var templatesPath = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	database.DB.Find(&products)
	templatesPath.ExecuteTemplate(w, "index", &products)
}

func New(w http.ResponseWriter, r *http.Request) {

	templatesPath.ExecuteTemplate(w, "new", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	name := r.FormValue("name")
	description := r.FormValue("description")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))

	newProduct := models.CreateNewProduct(name, description, price, quantity)
	database.DB.Create(&newProduct)

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	id := r.URL.Query().Get("id")
	database.DB.Delete(&product, id)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	id := r.URL.Query().Get("id")

	database.DB.Find(&product, id)
	templatesPath.ExecuteTemplate(w, "edit", &product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}

	id := r.URL.Query().Get("id")

	var existentProduct models.Product
	database.DB.First(&existentProduct, id)

	name := r.FormValue("name")
	description := r.FormValue("description")
	price, _ := strconv.ParseFloat(r.FormValue("price"), 64)
	quantity, _ := strconv.Atoi(r.FormValue("quantity"))

	updatedProduct := models.CreateNewProduct(name, description, price, quantity)
	database.DB.Model(&existentProduct).Updates(updatedProduct)
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
