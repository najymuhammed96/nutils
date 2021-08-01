package nutils

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

// Middleware is a wrapper type for func(http.HandlerFunc) http.HandlerFunc
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Middleware is a wrapper type for func(http.Handler) http.Handler
type MuxMiddleware func(http.Handler) http.HandlerFunc

// FuncUse allow the use of multiple middleware functions with http.HandlerFunc
func FuncUse(h http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

// FuncUse allow the use of multiple middleware functions with http.Handler (multiplexer)
func MuxUse(h http.Handler, m ...MuxMiddleware) http.Handler {
	// loop in reverse to preserve middleware order
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

func Log(mux http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		recorder := httptest.NewRecorder()
		mux.ServeHTTP(recorder, r)

		for k, v := range recorder.HeaderMap {
			w.Header()[k] = v
		}
		code := recorder.Code
		data, _ := ioutil.ReadAll(recorder.Body)
		WriteServerLog("localhost", r.RemoteAddr, r.Method, r.URL.RequestURI(), code)
		w.WriteHeader(code)
		w.Write(data)
	})
}

func SetJSONHeader(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	}
}
