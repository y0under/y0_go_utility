package main

import (
	"fmt"
	"http/server_base"
	"net/http"
	//"testing"
)

func main() {
	rt := server_base.Router{}
	rt.GET(`^/$`, func(parameter server_base.RouterParameter, writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(writer, "Hello, World!!!!11!")
	})

	rt.GET(`^/(?P<name>\w+)$`, func(parameter server_base.RouterParameter, writer http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(writer, "Hello, %v!!!!11!\n", parameter["name"])
	})

	rt.POST(`^/api$`, func(parameter server_base.RouterParameter, writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("content-type", "text/json")
		fmt.Fprintln(writer, `{"status: "OK"}"`)
	})

	http.ListenAndServe(":8080", &rt)
}
