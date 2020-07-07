package service

import (
	"log"
	"net/http"
)

func StartWebServer(port string) {

	r := NewRouter()                          // NEW
	http.Handle("/", r)                       // NEW
	err := http.ListenAndServe(":"+port, nil) // Goroutine will block here
	log.Println("starting http service at port", port)

	if err != nil {
		log.Println("An error occured while starting the Http listner at port ", port)
		log.Println("Error ", err.Error())
	}
}
