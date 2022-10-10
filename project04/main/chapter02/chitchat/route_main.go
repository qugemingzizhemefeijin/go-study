package main

import (
	"./data"
	"net/http"
)

// GET /err?msg=
// shows the error message page
func err(writer http.ResponseWriter, request *http.Request) {
	vals := request.URL.Query()
	sess, err := session(writer, request)
	if err != nil {
		generateHTML(writer, &data.Model{
			Data: vals.Get("msg"),
		}, "layout", "public.navbar", "error")
	} else {
		generateHTML(writer, &data.Model{
			Session: &sess,
			Data: vals.Get("msg"),
		}, "layout", "private.navbar", "error")
	}
}

func index(writer http.ResponseWriter, request *http.Request) {
	threads, err := data.Threads()
	if err != nil {
		error_message(writer, request, "Cannot get threads")
	} else {
		sess, err := session(writer, request)
		if err != nil {
			generateHTML(writer, &data.Model{
				Data:    threads,
			}, "layout", "public.navbar", "index")
		} else {
			generateHTML(writer, &data.Model{
				Session: &sess,
				Data:    threads,
			}, "layout", "private.navbar", "index")
		}
	}
}

