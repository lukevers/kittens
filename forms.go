package main

import (
	_ "github.com/gorilla/mux"
	"net/http"
	_ "io/ioutil"
)

func UpdateServer(w http.ResponseWriter, req *http.Request) {

	/*_, err := ioutil.ReadAll(req.Body)
	if err != nil {
		warn(err)
	}

	nick := req.FormValue("nick")
	warn(nick) */

	warn(req.URL.RawQuery)

}
