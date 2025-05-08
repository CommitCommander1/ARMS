package api

import (
	"fmt"
	"net/http"
)

func StartAPIServer() {
	http.HandleFunc("GET /harvesters", HarvesterGet) // POST new harv, GET list of active
	http.HandleFunc("POST /harvesters", HarvesterPost)
	http.HandleFunc("GET /harvesters/{id}", HarvesterIdGet) //GET/PUT/PATCH/DELETE detailed info of specific harv
	http.HandleFunc("PUT /harvesters/{id}", HarvesterIdPut)
	http.HandleFunc("DELETE /harvesters/{id}", HarvesterIdDelete)
	http.HandleFunc("/", mainServer)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Print(err)
	}
}
