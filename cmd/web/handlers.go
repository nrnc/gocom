package main

import "net/http"

func (app *application) VirtualTerminal(rw http.ResponseWriter, r *http.Request) {
	app.infoLogger.Println("Hit the handler")

	rw.Write([]byte("Hello"))
}
