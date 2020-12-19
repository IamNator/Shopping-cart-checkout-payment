package view

import (
	"html/template"
	"net/http"
)

func Payment(w http.ResponseWriter, req * http.Request){
	tpl := template.Must(template.ParseFiles("templates/checkout-page.html"))
	tpl.Execute(w, nil)
	//w.Write([]byte("payment checkout page"))
}

