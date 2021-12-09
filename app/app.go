package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/weeee9/hexagonal-architecture/domain"
	"github.com/weeee9/hexagonal-architecture/service"
)

func Start() {
	router := mux.NewRouter()

	// wiring
	customerHander := CustomerHandler{
		service: service.NewCustomerService(domain.NewCustomerRepository("postgres")),
	}

	router.HandleFunc("/customers", customerHander.getAllCustomer).Methods(http.MethodGet)
	router.HandleFunc("/customer/{id:[0-9]+}", customerHander.getCustomerByID).Methods(http.MethodGet)

	log.Println("service start at :8080 ...")
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
