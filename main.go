package main

import (
	"log"
	"net/http"

	"jwt/routes"
)

func main()  {
	
	r := routes.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}