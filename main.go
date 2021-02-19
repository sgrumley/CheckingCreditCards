package main

import(
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	controllers "github.com/sgrumley/CheckingCreditCards/controllers"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/VerifyCreditCard", controllers.VerifyCreditCard).Methods("POST")

	port := "8080"
	fmt.Println("served on " + port)

	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		fmt.Print(err)
	}
}