package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"transporte/src/controller"
	"transporte/src/library/sqlquery"
	"transporte/src/middleware"
	"transporte/src/models/tables"

	"github.com/gorilla/mux"
)

func RutasClienteCars(r *mux.Router) {

	s := r.PathPrefix("/car").Subrouter()
	s.Handle("/get/info-cls-a/data/", middleware.Autentication(http.HandlerFunc(allClienteCar))).Methods("GET")
	s.Handle("/get/info-cla-o/data/{c_plac}", middleware.Autentication(http.HandlerFunc(oneCLienteCar))).Methods("GET")
	s.Handle("/update/info-reg-o/data/{c_plac}", middleware.Autentication(http.HandlerFunc(updateClienteCar))).Methods("PUT")
	s.Handle("/create/info-reg-o/data/", middleware.Autentication(http.HandlerFunc(insertClienteCar))).Methods("POST")
}

func allClienteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	//get allData from database
	dataClieCar := sqlquery.NewQuerys("ClientesCars").Select("n_docu,l_marc,l_mode,l_color,c_year,c_mode,n_seri,n_pasa,c_plac").Exec().All()
	response.Data["clienteCars"] = dataClieCar
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

//Insert Clientes to DataBase
func insertClienteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	request_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_body := make(map[string]interface{})
	json.Unmarshal(request_body, &data_body)
	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_body)

	schema, table := tables.ClientesCars_GetSchema()
	clCar := sqlquery.SqlLibExec{}
	err = clCar.New(data_insert, table).Insert(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = clCar.Exec()
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func updateClienteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	c_plac := params["c_plac"]
	if c_plac == "" {
		response.Msg = "Error to write Cliente Car"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	request_body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_body := make(map[string]interface{})
	json.Unmarshal(request_body, &data_body)
	if len(data_body) <= 0 {
		response.Msg = "No se encontraron datos para actualizar"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	data_body["where"] = map[string]interface{}{"c_plac": c_plac}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.ClientesCars_GetSchema()
	clCar := sqlquery.SqlLibExec{}
	err = clCar.New(data_update, table).Update(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = clCar.Exec()
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func oneCLienteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	c_plac := params["c_plac"]
	if c_plac == "" {
		response.Msg = "Error to write Cliente Car"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//get allData from database
	dataClieCar := sqlquery.NewQuerys("ClientesCars").Select("c_plac,n_docu,l_marc,l_mode,l_color,c_year,c_mode,n_seri,n_pasa").Where("c_plac", "=", c_plac).Exec().One()
	response.Data = dataClieCar
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
