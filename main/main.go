package main

import (
	"flag"
	"fmt"
	"github.com/bradrydzewski/go.auth"
	"net/http"

)

var homepage = `
<html>
	<head>
		<title>Login</title>
	</head>
	<body>
		<div>Welcome to the go.auth Bitbucket demo</div>
		<div><a href="/auth/bitbucket">Authenticate with your Bitbucket Id</a><div>
	</body>
</html>
`

var privatepage = `
<html>
	<head>
		<title>Login</title>
	</head>
	<body>
		<div>oauth url: <a href="%s" target="_blank">%s</a></div>
		<div><a href="/auth/logout">Logout</a><div>
	</body>
</html>
`

 //private webpage, authentication required
func Private(w http.ResponseWriter, r *http.Request) {
	user := r.URL.User.Username()
	fmt.Fprintf(w, fmt.Sprintf(privatepage, user, user))
}

// public webpage, no authentication required
func Public(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, homepage)
}

// logout handler
func Logout(w http.ResponseWriter, r *http.Request) {
	auth.DeleteUserCookie(w, r)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func OnSuccess(w http.ResponseWriter, r *http.Request,  u auth.User, t auth.Token) {

	println("ok! "+ u.Name()+" "+t.Token())
}

func main() {

	// You should pass in your access key and secret key as args.
	// Or you can set your access key and secret key by replacing the defauаlt values below (2nd input param in flag.String)
	consumerKey := flag.String("consumer_key", "MuzSM5B3T2aw8ZGSq4", "your oauth consumer key")
	secretKey := flag.String("secret_key", "TFURz5wvpPCeKcEk9uzysyPKWBuDusXE", "your oauth secret key")
	flag.Parse()

	//url that google should re-direct to
	redirect := "http://localhost:8080/auth/bitbucket"

	// set the auth parameters
	auth.Config.CookieSecret = []byte("7H9xiimk2QdTdYI7rDddfJeV")
	auth.Config.LoginSuccessRedirect = "/private"
	auth.Config.CookieSecure = false


	// login handler
	//p := auth.NewBitbucketProvider(*consumerKey, *secretKey, redirect)
	bitbucketHandler := auth.Bitbucket(*consumerKey, *secretKey, redirect)

	bitbucketHandler.Success = OnSuccess
	http.Handle("/auth/bitbucket", bitbucketHandler)
	//http.Handle("/auth/bitbucket", auth.New(p))

	// logout handler
	http.HandleFunc("/auth/logout", Logout)

	// public urls
	http.HandleFunc("/", Public)

	// private, secured urls
	http.HandleFunc("/private", auth.SecureFunc(Private))

	println("bitbucket demo starting on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

/*import (
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
*/
