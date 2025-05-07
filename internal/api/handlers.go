package api

import (
	"fmt"
	"net/http"
)

func HarvesterGet(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Test")
}

func HarvesterPost(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Test")
}

func HarvesterIdGet(w http.ResponseWriter, req *http.Request) {
	id := req.PathValue("id")
	fmt.Println(id)
}
func mainServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Test")
}

func StartAPIServer() {
	http.HandleFunc("GET /harvesters", HarvesterGet) // POST new harv, GET list of active
	http.HandleFunc("POST /harvesters", HarvesterPost)
	http.HandleFunc("GET /harvesters/{id}", HarvesterIdGet) //GET/PUT/PATCH/DELETE detailed info of specific harv
	http.HandleFunc("/", mainServer)

	err := http.ListenAndServe(":8090", nil)
	if err != nil {
		fmt.Print(err)
	}
}
