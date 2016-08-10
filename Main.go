package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world! "+ r.URL.EscapedPath())
}

func bitbucketauth(w http.ResponseWriter, r *http.Request) {
	//resp, err := http.Post("https://bitbucket.org/site/oauth2/authorize?client_id=MuzSM5B3T2aw8ZGSq4&response_type=code")
	resp, err := http.Post("https://bitbucket.org/site/oauth2/access_token", "grant_type=password&username=ohl@ciklum.com&password=ct798wLas9", nil)
	if (err != nil){
		io.WriteString(w, "Ooops "+ err.Error())
	}else {
		io.WriteString(w, "No errors. Status "+resp.Status)
	}

	//io.WriteString(w, "Hello world! "+ r.URL.EscapedPath())
}

func main() {
	http.HandleFunc("/bitbucketauth", bitbucketauth)
	http.ListenAndServe(":8080", nil)
}
