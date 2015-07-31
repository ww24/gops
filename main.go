package main

import (
	"encoding/json"
	"flag"
	"net/http"

	"github.com/ww24/gops/libs"
	"github.com/ww24/gops/routes"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func api(c web.C, w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(libs.JSON{"status": "ok"})
	if err != nil {
		// status 500 返したい
		panic(err)
	}
}

func main() {
	libs.Configure()

	// goji.Get("/", func(c web.C, w http.ResponseWriter, r *http.Request) {
	// 	header := w.Header()
	// 	header.Add("Content-Type", "text/html; charset=utf-8")
	// 	view := mustache.Render("{{neko}}わーくす", context{"neko": "ねこねこ"})
	// 	fmt.Fprintln(w, view)
	// })
	goji.Get("/", routes.Top)
	goji.Get("/auth", routes.Auth)
	// static route
	goji.Get("/*", http.FileServer(http.Dir("public")))
	flag.Set("bind", ":8000")
	goji.Serve()
}
