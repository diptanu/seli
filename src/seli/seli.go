package main

import (
	"net/http"
	"flag"
	"log"
)

var config_file string

func init() {
	flag.StringVar(&config_file, "config", "~/seli.conf", "Path to the Seli Configuration file")
}

func main() {
	log.Println("Starting Seli")
	flag.Parse()
	log.Println("Config file is %s", config_file)
	http.Handle("/tmpfiles/", http.StripPrefix("/tmpfiles/", http.FileServer(http.Dir("/Users/diptanuc/Src/"))))
	http.ListenAndServe(":8080", nil)
}