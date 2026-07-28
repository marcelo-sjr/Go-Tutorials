package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
)

type ApiResponse struct {
	Sum float64 `json:"sum"`
	Average float64 `json:"avg"`
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Count int `json:"count"`
}

func main(){
	req, err := http.Get("http://localhost:8080/estatistica")
	if err != nil{
		log.Printf("Error: %v", err)
	}

	if req != nil{
		defer req.Body.Close()
		response := ApiResponse{}
		/*res, _ := io.ReadAll(req.Body)
		err = json.Unmarshal(res, &response)*/

		err = json.NewDecoder(req.Body).Decode(&response)
			if err != nil{
			log.Printf("Error %v", err)
		}
		fmt.Printf("%#v", response)
	}
}