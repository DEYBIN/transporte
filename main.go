package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transporte/src/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	enableCORS(r)
	r.HandleFunc("/", HomeHandler)
	//rutas de autentificacion
	routes.RutasAuth(r)
	routes.RutasSeguridad(r)
	routes.RutasCliente(r)
	routes.RutasClienteCars(r)
	routes.RutasServicio(r)
	fmt.Println("Server on port 5000")
	log.Fatal(http.ListenAndServe(":5000", r))

}

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	data := map[string]interface{}{"api": "apiexample", "version": 1.1}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Auth-Date, Auth-Periodo, Access-Token")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}
