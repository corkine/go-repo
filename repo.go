package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var port = flag.Int("port", 9001, "Http port")

const LOG = `0.0.1 2021.07.15 实现了基本的 go repo 逻辑(TODO:使用数据库增加自定义指向，比如 gitee 和 github)`

var VERSION = func() string {
	split := strings.Split(LOG, "\n")
	lastLine := split[len(split)-1]
	if lastLine == "" && len(split) > 1 {
		lastLine = split[len(split)-2]
	}
	lineSplit := strings.Split(lastLine, " ")
	return lineSplit[0]
}()

func main() {
	flag.Parse()
	help := func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{
			"version": VERSION,
			"log":     LOG,
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
		uri = strings.ReplaceAll(uri, "?go-get=0", "")
		uri = strings.ReplaceAll(uri, "?go-get=1", "")
		uri = strings.ReplaceAll(uri, ".git", "")
		repoHost := "repo.mazhangjing.com"
		goImport := fmt.Sprintf(
			`<meta content="%s%s git https://gitee.com/corkine%s.git" name="go-import">`,
			repoHost, uri, uri)
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
