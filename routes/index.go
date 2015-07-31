/**
 * Index router
 */

package routes

import (
	"fmt"
	"net/http"

	"github.com/hoisie/mustache"
	"github.com/zenazn/goji/web"
)

// Top page router
func Top(c web.C, w http.ResponseWriter, r *http.Request) {
	header := w.Header()
	header.Add("Content-Type", "text/html; charset=utf-8")
	view := mustache.RenderFile("views/index.html")
	fmt.Fprintln(w, view)
}
