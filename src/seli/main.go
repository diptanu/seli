package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	stats "seli/stats"
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

func upload_handler(res http.ResponseWriter, req *http.Request) {
	file, handler, err := req.FormFile("filedata")
	if err != nil {
		log.Printf("%s", err)
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Printf("%s", err)
	}
	file_name := path.Join(repo_path, handler.Filename)
	err = ioutil.WriteFile(file_name, data, 0777)
	if err != nil {
		log.Printf("%s", err)
	}
	log.Println("Written a new file :" + file_name)
	stats.SeliStats().Update()
	res.Header().Set("Content-Type", "text/plain")
	res.Write([]byte(req.Method))
}

func stats_handler(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	b, err := stats.SeliStats().Resource()
	if err != nil {
		log.Fatal(err)
	}
	res.Write(b)

}

func main() {
	flag.Parse()
	log.Println("Config file is %s", config_file)
	http.Handle(base_url, http.StripPrefix(base_url, http.FileServer(http.Dir(repo_path))))
	http.HandleFunc("/upload", upload_handler)
	http.HandleFunc("/stats/", stats_handler)
	log.Println("Starting Seli")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}
