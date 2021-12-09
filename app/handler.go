package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/weeee9/hexagonal-architecture/service"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (handler CustomerHandler) getAllCustomer(w http.ResponseWriter, r *http.Request) {
	customers, err := handler.service.GetAllCustomers()
	if err != nil {
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (handler CustomerHandler) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idParam := vars["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "error getting id: %s", err.Error())
		return
	}

	customer, err := handler.service.GetCustomerByID(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "error founding customer: %s", err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
