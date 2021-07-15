package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var port = flag.Int("port", 9001, "Http port")

func main() {
	flag.Parse()
	help := func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"version": "",
			"log":     "",
			"usage":   "import repo.mazhangjing.com/xxx to use https://gitee.com/corkine/xxx go mod repo",
		}
		bytes, err := json.Marshal(data)
		_, err = w.Write(bytes)
		if err != nil {
			log.Println(err)
		}
	}
	http.HandleFunc("/about", help)
	http.HandleFunc("/usage", help)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		uri := r.RequestURI
		host := r.Host
		goImport := fmt.Sprintf(
			`<meta content="%s%s git https://gitee.com/corkine%s.git" name="go-import">`,
			host, uri, uri)
		newURL := fmt.Sprintf("https://gitee.com/corkine%s.git", uri)
		htmlBody := fmt.Sprintf(`<!DOCTYPE html>
			<html lang='zh-CN'>
			<head>
			<title>Go Simple Repo by CM</title>
			<meta http-equiv="refresh" content="5;url=%s">
			%s
			</head>
			<body>
			<h1>Waiting for 5 seconds redirect to %s...</h1>
			</body>
			`, newURL, goImport, newURL)
		_, err := w.Write([]byte(htmlBody))
		if err != nil {
			log.Println(err)
		}
	})
	log.Printf("Server run on port %d", *port)
	log.Fatalln(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
