package main

import (
    "fmt"
    "net/http"
    "log"
    "io/ioutil"
    "encoding/json"
    "github.com/julienschmidt/httprouter"
)

func Calculate(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    input := Input{}
    reserr := ResponseErr{}
    res := Response{}
    c := make(chan int)
    reqBody, err := ioutil.ReadAll(r.Body)
    if err!= nil {
        fmt.Fprintf(w, "kindly enter correct data")
    }
    if err = json.Unmarshal(reqBody,&input); err != nil || input.A <0 || input.B <0{
		log.Printf("Incorrect Input, %v", err)
        w.WriteHeader(400) // Return 400 Bad Request.
        reserr = ResponseErr{Error: "Incorrect input"}
        json.NewEncoder(w).Encode(reserr)
        return
    } else {
    go Factorial(input.A,c)
	go Factorial(input.B,c)
	x,y := <-c,<-c
	Product := x*y
    log.Printf("Product of a! and b!: %v", Product)
    w.WriteHeader(http.StatusCreated)
    res = Response{ProductOfFact: Product}
    json.NewEncoder(w).Encode(res)
    
    }
}