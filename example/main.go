package main

import "github.com/haoxins/more"
import "net/http"
import "log"
import "io"

func main() {
	opts := map[string]interface{}{
		"dir":   "template/include",
		"ext":   "html",
		"cache": true,
	}

	r := more.New(opts)

	http.HandleFunc("/include", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		io.WriteString(w, r.Render("index", ""))
	})

	log.Println("port 3000 ...")
	http.ListenAndServe(":3000", nil)
}
