package main

import (
	"net/http"
	"flag"
	"log"
)

var config_file string
var port string
var base_url string 
var repo_path string

func init() {
	flag.StringVar(&config_file, "config", "~/seli.conf", "Path to the Seli Configuration file")
	flag.StringVar(&port, "port", ":80", "Port on which seli would run")
	flag.StringVar(&base_url, "baseurl", "/Custom/", "Base URL of the repo")
	flag.StringVar(&repo_path, "repopath", "/mnt/", "Path of the filesystem")
}

func main() {
	log.Println("Starting Seli")
	flag.Parse()
	log.Println("Config file is %s", config_file)
	http.Handle(base_url, http.StripPrefix(base_url, http.FileServer(http.Dir(repo_path))))
	err := http.ListenAndServe(port, nil)
        if err != nil {
         log.Fatal(err)
        }
}
