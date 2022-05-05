package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

	crypto "github.com/nparavinja/chiron-project/back-legs/encryption"
)

type Request struct {
	ServiceName string            `json:"serviceName"`
	MethodName  string            `json:"methodName"`
	Arguments   []json.RawMessage `json:"arguments"`
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
		// ok, err := crypto.ParseJWT(request.Header.Get("jwt"))
		// if !ok {
		// 	return "jwtError", err
		// }
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return "errorReadingRequest", err
		}
		var req Request
		err = json.Unmarshal(body, &req)
		if err != nil {
			return "errorUnmarshallingRequest", err
		}
		service := s.examinationService
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

func (s *server) handleAdmin() http.HandlerFunc {
	invoke := func(responseWriter http.ResponseWriter, request *http.Request) (interface{}, error) {
		// ok, err := crypto.ParseJWT(request.Header.Get("jwt"))
		// if !ok {
		// 	return "jwtError", err
		// }
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return "errorReadingRequest", err
		}
		var req Request
		err = json.Unmarshal(body, &req)
		if err != nil {
			return "errorUnmarshallingRequest", err
		}
		service := s.adminService
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

func logger(f func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Body)
		f(w, r) // original function call
	}
}

func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ok, _ := crypto.ParseJWT(r.Header.Get("jwt"))
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		next(w, r)
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
