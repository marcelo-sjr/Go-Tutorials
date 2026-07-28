package main

import (
	_"fmt"
	"log"
	"net/http"
)

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		_, err := w.Write([]byte("hello, world!"))
		if err != nil{
			log.Fatalf("error: %v", err)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}
}