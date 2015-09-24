package httpwrapper

import (
	"net/http"

	"github.com/gorilla/context"
	"github.com/julienschmidt/httprouter"
)

type Router struct {
	*httprouter.Router
}

func NewRouter() *Router {
	return &Router{httprouter.New()}
}

func wrapHandler(handler http.Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		context.Set(r, "params", ps)
		handler.ServeHTTP(w, r)
	}
}

func (r *Router) Get(path string, handler http.Handler) {
	r.GET(path, wrapHandler(handler))
}

func (r *Router) Post(path string, handler http.Handler) {
	r.POST(path, wrapHandler(handler))
}

func (r *Router) Put(path string, handler http.Handler) {
	r.PUT(path, wrapHandler(handler))
}

func (r *Router) Delete(path string, handler http.Handler) {
	r.DELETE(path, wrapHandler(handler))
}

func (r *Router) Head(path string, handler http.Handler) {
	r.HEAD(path, wrapHandler(handler))
}

func (r *Router) Options(path string, handler http.Handler) {
	r.OPTIONS(path, wrapHandler(handler))
}

func (r *Router) Patch(path string, handler http.Handler) {
	r.PATCH(path, wrapHandler(handler))
}
