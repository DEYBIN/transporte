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

func RutasServicio(r *mux.Router) {

	s := r.PathPrefix("/servicio").Subrouter()
	s.Handle("/get/info-cls-a/data/", middleware.Autentication(http.HandlerFunc(allServicio))).Methods("GET")
	s.Handle("/get/generate-fact/", middleware.Autentication(http.HandlerFunc(serviceFact))).Methods("GET")
	s.Handle("/get/info-cla-o/data/{id_serv}", middleware.Autentication(http.HandlerFunc(oneServicio))).Methods("GET")
	s.Handle("/update/info-reg-o/data/{id_serv}", middleware.Autentication(http.HandlerFunc(updateServicio))).Methods("PUT")
	s.Handle("/create/info-reg-o/data/", middleware.Autentication(http.HandlerFunc(insertServicio))).Methods("POST")
}

func allServicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	//get allData from database
	dataServicio := sqlquery.NewQuerys("Servicios").Select("c_year,c_mes,n_docu,f_fact,s_impo,c_plac,k_stad,f_digi,id_serv").Exec().All()
	response.Data["servicios"] = dataServicio
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}


func serviceFact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	//get all Data from database
	dataServicio := sqlquery.NewQuerys("Servicios").Select("n_docu,f_fact,s_impo,k_stad,f_digi,id_serv").Exec().All()
	response.Data["fact"] = dataServicio
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertServicio(w http.ResponseWriter, r *http.Request) {
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

	schema, table := tables.Servicios_GetSchema()
	servicio := sqlquery.SqlLibExec{}
	err = servicio.New(data_insert, table).Insert(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = servicio.Exec()
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

func updateServicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id_serv := params["id_serv"]
	if id_serv == "" {
		response.Msg = "Error to write service"
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

	data_body["where"] = map[string]interface{}{"id_serv": id_serv}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Servicios_GetSchema()
	servicio := sqlquery.SqlLibExec{}
	err = servicio.New(data_update, table).Update(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = servicio.Exec()
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

func oneServicio(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id_serv := params["id_serv"]
	if id_serv == "" {
		response.Msg = "Error to write the service"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//get allData from database
	dataServicio := sqlquery.NewQuerys("Servicios").Select("id_serv,c_year,c_mes,n_docu,f_fact,s_impo,c_plac,k_stad,f_digi").Where("id_serv", "=", id_serv).Exec().One()
	response.Data = dataServicio
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
