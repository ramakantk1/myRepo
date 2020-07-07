package model

// this is part 3 example
//https://callistaenterprise.se/blogg/teknik/2017/02/27/go-blog-series-part3/
//https://dzone.com/articles/go-microservices-part-3-embedded-database-and-json

//The case of the first letter denotes scoping (Upper-case == public, lower-case package-scoped).
//We also use the built-in support for declaring how each field should be serialized by the json.Marshal function in Go.

//it also uses BoltDB key-value store
//go get github.com/boltdb/bolt
type Account struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
