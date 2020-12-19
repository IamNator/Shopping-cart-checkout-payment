package main

import (
	"fmt"
	"github.com/IamNator/Shopping cart checkout payment/controller"
	"github.com/IamNator/Shopping cart checkout payment/view"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main(){
	fmt.Printf("Let's get started")
	router := mux.NewRouter()

	router.HandleFunc("/api/payment", controller.Payment)
	router.HandleFunc("/", view.Payment)
	log.Fatal(http.ListenAndServe(":8080", router))
}

