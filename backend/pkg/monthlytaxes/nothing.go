package monthlytaxes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMonthlyTaxes(dni string) string {
	return dni
}

func GetDNI(w http.ResponseWriter, r *http.Request) {
	// Parse the DNI parameter from the URL
	vars := mux.Vars(r)
	dni := vars["dni"]

	// call sunat api
	response := GetMonthlyTaxes(dni)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, response)
}

func main() {
	// Create a new router instance using the Gorilla Mux router
	router := mux.NewRouter()

	// Define a route for the API endpoint that accepts a DNI parameter in the URL
	router.HandleFunc("/monthly-taxes/{dni}", GetDNI).Methods("GET")

	// Start the HTTP server on port 8080
	fmt.Println("Server is running on :8080")
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}
