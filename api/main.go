package main

import (
	_ "fmt"
	"log"
	"net/http"
)



func main(){

	Cookie := http.Cookie{
		Name: "Test-Cookie",
		Value: "Test-Value",
		Quoted: true,
		MaxAge: 3600,
		Secure: true,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Test", `test-header`)
		http.SetCookie(w, &Cookie)
		w.WriteHeader(200)
		_, err := w.Write([]byte("hello, world!"))
		//log.Printf("Request: %v\n", r)
		if err != nil{
			log.Fatalf("error: %v", err)
		}
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil{
		log.Fatal(err)
	}
}