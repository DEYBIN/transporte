package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"transporte/src/controller"

	"github.com/gorilla/mux"
)

func RutasAuth(r *mux.Router) {

	s := r.PathPrefix("/auth").Subrouter()
	s.HandleFunc("/", auth).Methods("GET")
	s.HandleFunc("/login", login).Methods("PUT")

}

func auth(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func login(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "Aplication-Json")
	response := controller.NewResponseManager()

	req_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		response.Status = "Error"
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return 
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
