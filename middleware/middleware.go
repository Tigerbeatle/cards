package middleware

import (
	"net/http"
	"github.com/tigerbeatle/cards/models"
	"encoding/json"
	"reflect"
	"fmt"
	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
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

			//fmt.Println("Inside apiMiddleware BodyHandler r.Body:",r.Body)
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
/*
func AuthorizationHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		jwtToken := r.Header.Get("Token")
		fmt.Println("=======jwtToken:", jwtToken)

		token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(models.GetSecret()), nil
		})



		fmt.Println("err:",err)
		if err == nil && token.Valid {
		}else{
			if err.Error() == "token is expired"{
				// try to get a new token using the UUID inside this expired token
				var user models.User
				user.UUID = token.Claims["id"].(string)
				b, err := json.Marshal(user)
				if err != nil {
					fmt.Println(err)
					return
				}

				// send message to api to create user profile
				url := "http://127.0.0.1:8003/auth/1.0/accounts/generateJWT"
				client := &http.Client{Timeout: 10 * time.Second}
				req, _ := http.NewRequest("POST", url, bytes.NewBuffer(b))
				req.Header.Set("Accept", "application/vnd.api+json")
				req.Header.Set("content-type", "application/vnd.api+json")
				res, _ := client.Do(req)
				defer req.Body.Close()

				fmt.Println("res.StatusCode:",res.StatusCode)
				fmt.Println("res.Status:",res.Status)

				if res.StatusCode != 201 {
					if res.StatusCode == 805 {
						models.WriteError(w, models.ErrAccountDisabled)
						return
					}
					if res.StatusCode == 806 {
						models.WriteError(w, models.ErrLexpExpired)
						return
					} else {
						models.WriteError(w, models.ErrInternalServer)
						return
					}
				}

				//  Unpack body which is a json object to get jwtString
				bodyBytes, err := ioutil.ReadAll(res.Body)
				source := (*json.RawMessage)(&bodyBytes)
				var target models.BasicJSONReturn
				err = json.Unmarshal(*source, &target)
				if err != nil {
					// todo log this panic
					panic(err)
				}
				// replace the token in the r header with the new token
				r.Header.Set("Token", target.Payload)
				r.Header.Set("Token-renewed", "true")
			} else {
				models.WriteError(w, models.ErrUserTokenRejected)
				return
			}
		}

		next.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
*/
