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

func RutasClienteCar(r *mux.Router) {

	s := r.PathPrefix("/clienteCar").Subrouter()
	s.Handle("/get/info-cls-a/data/", middleware.Autentication(http.HandlerFunc(allClienteCar))).Methods("GET")
	s.Handle("/get/info-cla-o/data/{n_docu}", middleware.Autentication(http.HandlerFunc(oneCLiente))).Methods("GET")
	s.Handle("/update/info-reg-o/data/{n_docu}", middleware.Autentication(http.HandlerFunc(updateCliente))).Methods("PUT")
	s.Handle("/create/info-reg-o/data/", middleware.Autentication(http.HandlerFunc(insertClienteCar))).Methods("POST")
}

func allClienteCar(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	//get allData from database
	dataClientesCars := sqlquery.NewQuerys("ClientesCars").Select("l_marc,l_mode,l_color,c_year,c_mode,n_seri,n_pasa").Exec().All()
	response.Data["clienteCars"] = dataClientesCars
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
	clienteCar := sqlquery.SqlLibExec{}
	err = clienteCar.New(data_insert, table).Insert(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = clienteCar.Exec()
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
	n_docu := params["n_docu"]
	if n_docu == "" {
		response.Msg = "Error to write user"
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

	data_body["where"] = map[string]interface{}{"n_docu": n_docu}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Clientes_GetSchema()
	cliente := sqlquery.SqlLibExec{}
	err = cliente.New(data_update, table).Update(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = cliente.Exec()
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
	n_docu := params["n_docu"]
	if n_docu == "" {
		response.Msg = "Error to write Cliente"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//get allData from database
	dataCliente := sqlquery.NewQuerys("Clientes").Select("n_docu,c_docu,c_clie,k_gene,f_naci,l_dire,l_refe,c_ubig,n_tele,n_celu,l_obse").Where("n_docu", "=", n_docu).Exec().One()
	response.Data = dataCliente
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
