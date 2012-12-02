package main

import (
	"net/http"
	"flag"
	"log"
)

func main() {
	log.Println("Starting Seli")
	var config_file = flag.String("config", "~/seli.conf", "Path to the Seli Configuration file")
	flag.Parse()
	log.Println("Config file is %s", *config_file)
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/Users/diptanuc/Src/"))))
	http.ListenAndServe(":8080", nil)
}