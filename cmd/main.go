package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"matching-timestamps/app/services"
	"matching-timestamps/presentation/periodic_task"
	"net/http"
	"os"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	serveOnPort := getEnv("SERVE_ON_PORT", "8000")

	r := mux.NewRouter()
	s := r.PathPrefix("").Subrouter()

	ptService := services.NewPeriodicTaskService()

	ptRoutes := periodic_task.CreateRoutes(ptService)

	for routePath, routeMethods := range ptRoutes {
		fmt.Printf("adding %s route with methods %v\n", routePath, routeMethods.Methods)
		if routeMethods.Handler != nil {
			s.Handle(routePath, routeMethods.Handler).Methods(routeMethods.Methods...)
		} else {
			s.HandleFunc(routePath, routeMethods.HandlerFunc).Methods(routeMethods.Methods...)
		}
	}

	http.Handle("/", r)
	// Bind to a port and pass our router in
	fmt.Printf("Started matching-timestamps service at port %s\n", serveOnPort)
	log.Fatal(http.ListenAndServe(":"+serveOnPort, r))
}
