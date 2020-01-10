package main

import (
	"balena-factoryreset-example/constants"
	"balena-factoryreset-example/utils"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	log.Println("TESTf1!dddfdddcdd")
	err := ioutil.WriteFile("/data/testFile", []byte("Hello World"), 0644)
	if err != nil {
		log.Println("Error writing testfile")
	}

	router := mux.NewRouter()
	router.Methods("POST").Path("/factoryReset").HandlerFunc(factoryReset)

	err = http.ListenAndServe(fmt.Sprintf("%s:%s", utils.GetEnv("HOST", constants.Host), utils.GetEnv("PORT", constants.Port)), router)
	if err != nil {
		log.Fatal("Webserver could not start: ", err)
	}
}

func factoryReset(w http.ResponseWriter, r *http.Request) {

	log.Println("Resetting wifi connections")
	err := utils.DeleteWirelessConnections()
	if err != nil {
		respondWithJSON(w, 500, err)
		return
	}

	log.Println("Resetting applications")
	err = utils.PurgeData()
	if err != nil {
		respondWithJSON(w, 500, err)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"status": "OK"})
}

func respondWithError(w http.ResponseWriter, r *http.Request, code int, err error) {
	log.Print(fmt.Sprintf("Error during request: %s", err.Error()))

	respondWithJSON(w, code, map[string]string{"error": err.Error()})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
