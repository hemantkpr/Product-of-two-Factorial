package main
type Input struct{
    A int `json:"a"`
    B int `json:"b"`
}
type Response struct{
    ProductOfFact int
}
type ResponseErr struct {
    Error string
}