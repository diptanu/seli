package main

import (
	"net/http"
	"flag"
	"log"
	"io/ioutil"
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

func upload_handler(w http.ResponseWriter, req *http.Request){
	file, handler , err := req.FormFile("filedata")
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(handler.Filename, data, 0777) 
        if err != nil { 
                log.Println(err) 
        } 
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(req.Method))
}

func main() {
	log.Println("Starting Seli")
	flag.Parse()
	log.Println("Config file is %s", config_file)
	http.Handle(base_url, http.StripPrefix(base_url, http.FileServer(http.Dir(repo_path))))
	http.HandleFunc("/upload", upload_handler)
	err := http.ListenAndServe(port, nil)
        if err != nil {
		log.Fatal(err)
        }
}
