package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"transporte/src/controller"
	"transporte/src/library/date"
	"transporte/src/library/lib"
	"transporte/src/library/sqlquery"
	"transporte/src/middleware"
	"transporte/src/models/tables"

	"github.com/gorilla/mux"
)

func RutasServicio(r *mux.Router) {

	s := r.PathPrefix("/servicio").Subrouter()
	s.Handle("/get/info-cls-a/data/", middleware.Autentication(http.HandlerFunc(allServicio))).Methods("GET")
	s.Handle("/get/generate-fact/{id_serv}", middleware.Autentication(http.HandlerFunc(serviceFact))).Methods("GET")
	s.Handle("/get/info-cla-o/data/{id_serv}", middleware.Autentication(http.HandlerFunc(oneServicio))).Methods("GET")
	s.Handle("/get/service-cliente/data/{n_docu}", middleware.Autentication(http.HandlerFunc(servicioCliente))).Methods("GET")
	s.Handle("/update/info-reg-o/data/{id_serv}", middleware.Autentication(http.HandlerFunc(updateServicio))).Methods("PUT")
	s.Handle("/create/info-reg-o/data/", middleware.Autentication(http.HandlerFunc(insertServicio))).Methods("POST")
	s.Handle("/update/service-alta/data/{id_serv}", middleware.Autentication(http.HandlerFunc(darAlta))).Methods("PUT")
	s.Handle("/update/service-baja/data/{id_serv}", middleware.Autentication(http.HandlerFunc(darBaja))).Methods("PUT")
	s.Handle("/get/det-fact/{id_serv}", middleware.Autentication(http.HandlerFunc(detalleFactura))).Methods("GET")

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
	params := mux.Vars(r)
	id_serv := params["id_serv"]
	if id_serv == "" {
		response.Msg = "Error to write the service"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//get all Data from database
	dataServicio := sqlquery.NewQuerys("Servicios").Select("n_docu,f_fact,s_impo,k_stad,f_digi").Where("id_serv", "=", id_serv).Exec().All()
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
		response.Msg = "Error to get the service"
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

func darAlta(w http.ResponseWriter, r *http.Request) {
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

	data_body := make(map[string]interface{})
	data_body["k_stad"] = int64(0)
	// retorna fecha de formato string dd/mm/yyyy (America Bogota)
	data_body["f_fact"] = date.GetFechaLocationString()

	data_body["where"] = map[string]interface{}{"id_serv": id_serv}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Servicios_GetSchema()
	servicio := sqlquery.SqlLibExec{}
	err := servicio.New(data_update, table).Update(schema)
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

func darBaja(w http.ResponseWriter, r *http.Request) {
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

	data_body := make(map[string]interface{})
	data_body["k_stad"] = int64(0)

	data_body["where"] = map[string]interface{}{"id_serv": id_serv}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Servicios_GetSchema()
	servicio := sqlquery.SqlLibExec{}
	err := servicio.New(data_update, table).Update(schema)
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

func servicioCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	// id_serv := params["id_serv"]
	n_docu := params["n_docu"]
	if n_docu == "" {
		response.Msg = "Error to get service client"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//get all Data from database
	dataServicio := sqlquery.NewQuerys("Servicios").Select("c_year,c_mes,f_fact,s_impo,c_plac,k_stad,f_digi,id_serv").Where("n_docu", "=", n_docu).Exec().All()
	if len(dataServicio) <= 0 {
		response.Msg = "cliente no cuenta con ningun servicio"
		//Mensaje informativo
		response.StatusCode = 100
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	response.Data["services"] = dataServicio
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func detalleFactura(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()

	params := mux.Vars(r)
	id_serv := params["id_serv"]

	if id_serv == "" {
		response.Msg = "Error to get fact client"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	
	var fact []map[string]interface{}
	fact = append(fact, map[string]interface{}{
		"periodo": "2022-04",
	})
	fact = append(fact, map[string]interface{}{
		"periodo": "2022-05",
	})
	var newFact []string
	for _, v := range fact {
		newFact = append(newFact, v["periodo"].(string))

	}
	fmt.Println(newFact)
	index := lib.IndexOfStrings(newFact, "2022-08")
	fmt.Println(index)

	date_fact := date.GetDate("02/04/2020")
	date_now := date.GetDateLocation()
	//date_end := date.GetLastDateOfMonth(date_fact)
	/*if date_fact == date_end {
		is_last_of_month = true
	}*/

	month_init := int64(date_fact.Month())
	year_init := date_fact.Year()
	month_now := int64(date_now.Month())
	year_now := date_now.Year()

	//var data_facturaciones []map[string]interface{}
	var month = int64(12)

	for i := year_init; i <= year_now; i++ {
		if i == year_now {
			month = month_now
		}
		for e := month_init; e <= month; e++ {
			// fmt.Println(i, e)
			year := fmt.Sprintf("%v", i)
			month := fmt.Sprintf("%02d", e)
			if lib.IndexOfStrings(newFact, year+"-"+month) == -1 {
				fmt.Println(year, month)
			}
		}

		month_init = 1
	}

	dataServicio := sqlquery.NewQuerys("ServiciosDetalle").Select().Where("id_serv", "=", id_serv).And("k_stad", "=", 0).Exec().All()
	if len(dataServicio) <= 0 {
		response.Msg = "Servicio no existe o bien se encuentra suspendido"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	response.Data["detalleFact"] = dataServicio
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
