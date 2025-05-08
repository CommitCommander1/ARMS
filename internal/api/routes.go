package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Harvester struct {
	ID   int    `json:"ID"`
	Name string `json:"Name"`
	X    int    `json:"X"`
	Y    int    `json:"Y"`
}

var HarvesterList []Harvester
var HarvesterMutex sync.Mutex

func HarvesterGet(w http.ResponseWriter, req *http.Request) {
	HarvesterMutex.Lock()
	defer HarvesterMutex.Unlock()

	if len(HarvesterList) == 0 {
		w.WriteHeader(http.StatusOK) // No content is a valid response
		json.NewEncoder(w).Encode([]Harvester{})
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(HarvesterList); err != nil {
		http.Error(w, "Failed to encode harvesters", http.StatusInternalServerError)
		return
	}
}

func HarvesterPost(w http.ResponseWriter, req *http.Request) {
	var newHarvester Harvester
	if err := json.NewDecoder(req.Body).Decode(&newHarvester); err != nil {
		log.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	HarvesterMutex.Lock()
	newHarvester.ID = len(HarvesterList) + 1 // Assign a new ID
	HarvesterList = append(HarvesterList, newHarvester)
	HarvesterMutex.Unlock()

	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(newHarvester); err != nil {
		http.Error(w, "Failed to encode new harvester", http.StatusInternalServerError)
		return
	}
}

func HarvesterIdGet(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))
	HarvesterMutex.Lock()
	defer HarvesterMutex.Unlock()

	for _, harvester := range HarvesterList {
		if harvester.ID == id {
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(harvester); err != nil {
				http.Error(w, "Failed to encode harvester", http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.NotFound(w, req)
}

func HarvesterIdPut(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))
	var updatedHarvester Harvester
	if err := json.NewDecoder(req.Body).Decode(&updatedHarvester); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	HarvesterMutex.Lock()
	defer HarvesterMutex.Unlock()

	for i, harvester := range HarvesterList {
		if harvester.ID == id {
			updatedHarvester.ID = id // Ensure the ID in the body matches the URL
			HarvesterList[i] = updatedHarvester
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(updatedHarvester); err != nil {
				http.Error(w, "Failed to encode updated harvester", http.StatusInternalServerError)
				return
			}
			return
		}
	}

	http.NotFound(w, req)

}

func HarvesterIdDelete(w http.ResponseWriter, req *http.Request) {
	id, _ := strconv.Atoi(req.PathValue("id"))
	for num, harvester := range HarvesterList {
		if id == harvester.ID {
			HarvesterList = append(HarvesterList[:num], HarvesterList[num+1:]...)
			return
		}
	}
	http.NotFound(w, req)

}
func mainServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Test")
}
