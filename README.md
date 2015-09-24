# httpwrapper
httpwrapper is tiny library written in Golang for serving http requests and dealing with middlewares more easier.  

## Installing

go get github.com/minhduccm/httpwrapper

## Example

	package main

	import (
		"net/http"

		"github.com/minhduccm/httpwrapper"
	)

	func TestHandlerFunc(res http.ResponseWriter, req *http.Request) (int, error) {
		res.Write([]byte("pong 123"))
		return 200, nil
	}

	func BarHandlerFunc(res http.ResponseWriter, req *http.Request) (int, error) {
		res.Write([]byte("bar"))
		return 200, nil
	}

	func AuthFunc(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {

			queryValues := r.URL.Query()
			key := queryValues.Get("key")
			if key != "minhduccm" {
				http.Error(w, http.StatusText(401), 401)
				return
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}

	func CustomFunc(key string) func(next http.Handler) http.Handler {
		return func(next http.Handler) http.Handler {
			fn := func(w http.ResponseWriter, r *http.Request) {

				queryValues := r.URL.Query()
				querykey := queryValues.Get("key")
				if querykey != key {
					http.Error(w, http.StatusText(401), 401)
					return
				}
				next.ServeHTTP(w, r)
			}
			return http.HandlerFunc(fn)
		}
	}

	func main() {
		router := httpwrapper.NewRouter()

		router.Get("/ping", httpwrapper.Perform(httpwrapper.LoggingHandler, httpwrapper.RecoveryHandler, AuthFunc, CustomFunc("minhduccm")).LeadTo(TestHandlerFunc))

		router.Get("/foo", httpwrapper.Perform().LeadTo(BarHandlerFunc))

		http.ListenAndServe(":8080", router)
	}


## More detail, see docs:
http://godoc.org/github.com/minhduccm/httpwrapper