package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"transporte/src/controller"
	"transporte/src/library/sqlquery"
	"transporte/src/middleware"
	"transporte/src/models/tables"

	"github.com/gorilla/mux"
)

func RutasSeguridad(r *mux.Router) {

	s := r.PathPrefix("/user").Subrouter()
	s.Handle("/get/info-cls-a/data/", middleware.Autentication(http.HandlerFunc(allUser))).Methods("GET")
	s.Handle("/get/info-cla-o/data/{id}", middleware.Autentication(http.HandlerFunc(oneUser))).Methods("GET")
	s.Handle("/update/info-reg-o/data/{id}", middleware.Autentication(http.HandlerFunc(updateUser))).Methods("PUT")
	s.Handle("/create/info-reg-o/data/", middleware.Autentication(http.HandlerFunc(insertUser))).Methods("POST")
	// s.Handle("/generate/n-usr/data/", middleware.Autentication(http.HandlerFunc(generateUser))).Methods("POST")
}

func allUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "-Json")
	response := controller.NewResponseManager()

	//get allData from database
	dataUser := sqlquery.NewQuerys("Seguridad").Select("k_carg,n_docu,users,l_nomb,l_apl1,l_apl2,id").Exec().All()
	response.Data["users"] = dataUser
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func insertUser(w http.ResponseWriter, r *http.Request) {
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

	user := generateUser(data_body["l_nomb"].(string), data_body["l_apl1"].(string), data_body["l_apl2"].(string))
	data_body["users"] = user


	var data_insert []map[string]interface{}
	data_insert = append(data_insert, data_body)

	schema, table := tables.Seguridad_GetSchema()
	seguridad := sqlquery.SqlLibExec{}
	err = seguridad.New(data_insert, table).Insert(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = seguridad.Exec()
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)

	}
	//datos insertado
	returnData := seguridad.Data[0]
	delete(returnData, "id")
	delete(returnData, "l_pass")
	delete(returnData, "users")
	delete(returnData, "n_docu")
	response.Data = returnData

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

//DONE:Generate user

func generateUser(name string, apl1 string, apl2 string) string {
	existNombre2 := false
	array_nombre := strings.Split(name, " ")
	nombre1 := array_nombre[0]
	nombre2 := ""

	if len(array_nombre) == 2 {
		existNombre2 = true
		nombre2 = array_nombre[1]
	}
	var user string
	for i := 0; i <= 10; i++ {
		switch i {
		case 0:
			l := nombre1[0:3]
			user = l + apl1
		case 1:
			l := ""
			if existNombre2 {
				l = nombre2[0:1]

			} else {
				l = nombre1[0:1]

			}
			user = l + apl1
		case 2:
			l := ""
			if existNombre2 {
				l = nombre2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l + apl1
		case 3:
			l := ""
			if existNombre2 {
				l = nombre1[0:2] + apl1[0:2] + apl2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 4:
			l := ""
			if existNombre2 {
				l = nombre1[0:3] + apl1[0:2] + apl2[0:3]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 5:
			l := ""
			if existNombre2 {
				l = nombre1[0:4] + apl1[0:2] + apl2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 6:
			l := ""
			if existNombre2 {
				l = nombre1 + apl1[0:1] + apl2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 7:
			l := ""
			if existNombre2 {
				l = nombre1[0:2] + apl1 + apl2[0:1]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 8:
			l := ""
			if existNombre2 {
				l = nombre1[0:3] + apl1[0:2] + apl2
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 9:
			l := ""
			if existNombre2 {
				l = nombre1 + apl1[0:2] + apl2[0:2]
			} else {
				l = nombre1[0:1]
			}
			user = l
		case 10:
			l := ""
			if existNombre2 {
				l = nombre1 + apl1[0:3] + apl2[0:4]
			} else {
				l = nombre1[0:1]
			}
			user = l
		}
		//consulta\
		usuario := sqlquery.NewQuerys("Seguridad").Select("users").Where("users","=",user).Exec().Text("users")

		if usuario == nil {
			break
		}
	}
	return user

}

func updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
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

	data_body["where"] = map[string]interface{}{"id": id}
	var data_update []map[string]interface{}
	data_update = append(data_update, data_body)

	schema, table := tables.Seguridad_GetSchema()
	seguridad := sqlquery.SqlLibExec{}
	err = seguridad.New(data_update, table).Update(schema)
	if err != nil {
		response.Msg = err.Error()
		response.StatusCode = 300
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = seguridad.Exec()
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

func oneUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content Type", "Aplication-Json")
	response := controller.NewResponseManager()
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		response.Msg = "Error to write user"
		response.StatusCode = 400
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
		return
	}
	//get allData from database
	dataUser := sqlquery.NewQuerys("Seguridad").Select("id,users,l_nomb,l_apl1,l_apl2,k_carg,l_emai,n_celu").Where("id", "=", id).Exec().One()
	response.Data = dataUser
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
