package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world! "+ r.URL.EscapedPath())
}

func bitbucketauth(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("https://bitbucket.org/site/oauth2/authorize?client_id=MuzSM5B3T2aw8ZGSq4&response_type=code")
	if (err != nil){
		io.WriteString(w, "Ooops "+ err.Error())
	}else {
		io.WriteString(w, resp.Body.Read())
	}

	//io.WriteString(w, "Hello world! "+ r.URL.EscapedPath())
}

func main() {
	http.HandleFunc("/bitbucketauth", bitbucketauth)
	http.ListenAndServe(":8080", nil)
}
