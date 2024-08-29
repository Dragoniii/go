package main

import "net/http"

func middlewareHandler(handler http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        handler.ServeHTTP(w, r)
    })
}

func intendedFunction(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Ok..."))
}

func middleware() {
    intendedHandler := http.HandlerFunc(intendedFunction)
    http.Handle("/foo", middlewareHandler(intendedHandler))
    http.ListenAndServe(":5000", nil)
}