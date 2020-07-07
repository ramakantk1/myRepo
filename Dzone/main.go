package main

//this example is from https://dzone.com/articles/go-microservices-part-2-building-our-first-service?fromrel=true
//https://dzone.com/articles/go-microservices-blog-series-part-1
// the second part also talk about performace foot print using tool Gatling
import (
	"fmt"
	"microservices/my-project/Dzone/dbclient"         //is the folder name where we have the package
	service "microservices/my-project/Dzone/services" //is the folder name where we have the package
)

var appName = "accountservice"

func main() {

	initializeBoltClient()
	fmt.Printf("Starting %v\n", appName)
	service.StartWebServer("6767")
	fmt.Printf("%v Web server is started\n", appName)

}

// Creates instance and calls the OpenBoltDb and Seed funcs
func initializeBoltClient() {
	service.DBClient = &dbclient.BoltClient{}
	service.DBClient.OpenBoltDb()
	service.DBClient.Seed()
}
