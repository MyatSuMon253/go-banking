package app

import (
	"log"
	"net/http"

	"github.com/MyatSuMon253/go-banking/domain"
	"github.com/MyatSuMon253/go-banking/service"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	ch := CustomerHandlers{service: service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	router.HandleFunc("/customers", ch.getAllCustomers).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
