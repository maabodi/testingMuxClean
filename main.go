package main

import (
	"cobaclean/internal/config"
	"cobaclean/internal/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	conf := config.InitConfiguration()
	router := handler.ConfigureRouter(conf)

	fmt.Print("\nRunning on port : 1323\n")
	log.Fatal(http.ListenAndServe("0.0.0.0:1323", router))
}
