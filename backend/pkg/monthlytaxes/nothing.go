package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

func GetMonthlyTaxes(dni string) (string, error) {
	token := "ASPJGDSP4SD783H375S"
	// Construct the URL with the provided DNI and token.
	url := fmt.Sprintf("http://api.aciertaperu.com/app4/v2/reniec?dni=%s&token=%s", dni, token)

	// Perform an HTTP GET request to the API.
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Check the response status code.
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API request failed with status code %d", resp.StatusCode)
	}

	// Read the response body.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Assuming the API returns a string response, you can return it here.
	return string(body), nil
}

func GetDNI(w http.ResponseWriter, r *http.Request) {
	// Parse the DNI parameter from the URL
	vars := mux.Vars(r)
	dni := vars["dni"]

	// call sunat api
	response, _ := GetMonthlyTaxes(dni)

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
