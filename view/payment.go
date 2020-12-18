package view

import "net/http"

func Payment(w http.ResponseWriter, req * http.Request){
	w.Write([]byte("payment checkout page"))
}