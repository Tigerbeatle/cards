package middleware

import (
	"net/http"
	"github.com/tigerbeatle/cards/models"
	"encoding/json"
	"reflect"
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
	//"github.com/dgrijalva/jwt-go"
	//"time"
	//"bytes"
	//"io/ioutil"
)

func RecoverHandler(next http.Handler) http.Handler {
	//Recover is a built-in function that regains control of a panicking goroutine.
	//Recover is only useful inside deferred functions. During normal execution,
	//a call to recover will return nil and have no other effect. If the current
	//goroutine is panicking, a call to recover will capture the value given to
	//panic and resume normal execution.
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				models.WriteError(w, models.ErrInternalServer)
			}
		}()

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func AcceptHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//fmt.Println("r.URL.Path = ", r.URL.Path)
		//fmt.Println("r.Header Accept = ", r.Header.Get("Accept"))
		//fmt.Println("r.Header Authorization = ", r.Header.Get("Authorization"))
		if r.Header.Get("Accept") != "application/vnd.api+json" {
			models.WriteError(w, models.ErrNotAcceptable)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func BodyHandler(v interface{}) func(http.Handler) http.Handler {
	t := reflect.TypeOf(v)
	m := func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			val := reflect.New(t).Interface()
			err := json.NewDecoder(r.Body).Decode(val)
			fmt.Println("Middleware BodyHandler - val:",val)
			if err != nil {
				models.WriteError(w, models.ErrBadRequest)
				return
			}

			if next != nil {
				context.Set(r, "body", val)

				params := context.Get(r, "params").(httprouter.Params)
				fmt.Println("middle params:",params)

				next.ServeHTTP(w, r)
			}
		}
		return http.HandlerFunc(fn)
	}
	return m
}

func ContentTypeHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/vnd.api+json" {
			models.WriteError(w, models.ErrUnsupportedMediaType)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
