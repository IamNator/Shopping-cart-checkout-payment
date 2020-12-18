package controller

import "net/http"

func Payment(w http.ResponseWriter, req * http.Request){
	w.Write([]byte("Hello there"))
}