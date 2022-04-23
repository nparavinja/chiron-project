package main

import (
	"encoding/json"
	"io/ioutil"
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
		service := s.patientService
		method := reflect.ValueOf(service).MethodByName(req.MethodName)
		parameters := make([]reflect.Value, len(req.Arguments))
		for i, parameter := range req.Arguments {
			p := reflect.New(method.Type().In(i))
			instance := p.Interface()
			err := json.Unmarshal([]byte(parameter), &instance)
			if err != nil {
				return "errorUnmarshalingParameter", err
			}
			parameters[i] = p.Elem()
		}
		result := method.Call(parameters)
		// result is a slice of []reflect.Value
		error, ok := result[1].Interface().(error)
		if ok {
			return nil, error
		}
		return result[0].Interface(), nil
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		result, err := invoke(w, r)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error(), "errorDetails": err})
		} else {
			//
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

/*
USERS
{
	"serviceName": UserService/ExaminationService
	"methodName:" login/register/delete      // ....
	"parameters": [1, 2, 3, 4]
}

ADMIN - CRUD users, patients, doctors

*/
