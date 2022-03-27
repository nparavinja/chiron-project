package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
)

type Request struct {
	ServiceName string            `json:"serviceName"`
	MethodName  string            `json:"methodName"`
	Arguments   []json.RawMessage `json:"arguments"`
}

func (s *server) routes() {
	s.router.HandleFunc("/api/usr/", s.handleUser()).Methods("POST")
	s.router.HandleFunc("/api/exm/", s.handleExamination()).Methods("POST")
	s.router.HandleFunc("/admin/", s.handleAdmin())
}

func (s *server) handleUser() http.HandlerFunc {
	invoke := func(responseWriter http.ResponseWriter, request *http.Request) (interface{}, error) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return "errorReadingRequest", err
		}
		var req Request
		err = json.Unmarshal(body, &req)
		if err != nil {
			return "errorUnmarshallingRequest", err
		}
		log.Println(req)
		return callServiceMethod(s.userService, req.MethodName, req.Arguments)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := invoke(w, r)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error(), "errorDetails": err})
		} else {
			json.NewEncoder(w).Encode(map[string]interface{}{"result": result})
		}
	})
}

func (s *server) handleExamination() http.HandlerFunc {
	invoke := func(responseWriter http.ResponseWriter, request *http.Request) (interface{}, error) {
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return "errorReadingRequest", err
		}
		var req Request
		err = json.Unmarshal(body, &req)
		if err != nil {
			return "errorUnmarshallingRequest", err
		}

		return nil, err
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var result interface{}
		var err error
		if err == nil {
			result, err = invoke(w, r)
		}
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error(), "errorDetails": err})
		} else {
			json.NewEncoder(w).Encode(map[string]interface{}{"result": result})
		}
	})
}

func (s *server) handleAdmin() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// use thing
	}
}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func callServiceMethod(service interface{}, methodName string, data ...interface{}) (interface{}, error) {
	parameters := make([]reflect.Value, len(data))
	for i := range parameters {
		parameters[i] = reflect.ValueOf(data[i])
	}
	// result is a slice of []reflect.Value
	results := reflect.ValueOf(service).MethodByName(methodName).Call(parameters)
	log.Println("Service called: ", methodName, results)
	error, ok := results[1].Interface().(error)
	if ok {
		return nil, error
	}
	return results[0].Interface(), nil
}

/*
USERS
{
	"serviceName": UserService/ExaminationService
	"methodName:" login/register/delete      // ....
	"parameters": [1, 2, 3, 4]
}

ADMIN - CRUD users, patients, doctors

*/
