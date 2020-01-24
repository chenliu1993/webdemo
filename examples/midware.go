package examples

import (
	"log"
	"net/http"
)

// MiddleWareOne example for midware.
func MiddleWareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("midWareOne executing...")
		next.ServeHTTP(w, r)
		log.Println("midWareOne done.")
	})
}

// MiddleWareTwo example for midware.
func MiddleWareTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("midWareTwo executing...")
		if r.URL.Path != "/" {
			return
		}
		next.ServeHTTP(w, r)
		log.Println("middleWareTwo done.")
	})
}

// FinalFunc exmaple for midware.
func FinalFunc(w http.ResponseWriter, r *http.Request) {
	log.Println("Final Excutor.")
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Hello MiddleWare"))
}
