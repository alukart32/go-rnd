package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var assetsDir string

func init() {
	path, e := os.Getwd()
	if e != nil {
		log.Fatal(e)
	}
	assetsDir = filepath.FromSlash(path + "/assets/")
}

type CounterHandler struct {
	counter int
}

func (c *CounterHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Println(c.counter)
	c.counter++
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Counter:", c.counter)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - not found\n"))
		return
	}

	fmt.Fprintln(w, "Hello World!!!")
}

func ua(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User agent: %s\n", r.Header.Get("User-Agent"))
}

func urlPassParam(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!\n", r.URL.Path[5:])
}

// URL
// scheme:[//[user:password@]host[:port]][/]path[?query][#fragment]
// https://example.com/path/page?name=John&occupation=teacher
func queryNameHandler(w http.ResponseWriter, r *http.Request) {
	keys, err := r.URL.Query()["name"]
	ret := "guest"
	if err {
		ret = keys[0]
	}
	fmt.Fprintf(w, "Hello, %s\n", ret)
}

func process(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, filepath.Join(assetsDir, "form.html"))
	case "POST":
		if e := r.ParseForm(); e != nil {
			http.Error(w, "ParseForm() err : "+e.Error(), http.StatusInternalServerError)
			return
		}

		name := r.FormValue("name")
		occupation := r.FormValue("occupation")
		fmt.Fprintf(w, "%s is a %s\n", name, occupation)
	default:
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
	}
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	f, e := os.ReadFile(filepath.Join(assetsDir, "conccurency_visibility_synch_guarantees.PNG"))
	if e != nil {
		log.Fatal(e)
	}
	w.Header().Set("Content-Type", "image/png")
	w.Write(f)
}

func main() {
	// add a new handler
	http.Handle("/count", new(CounterHandler))
	// add new handlefuncs
	http.HandleFunc("/home", home)
	http.HandleFunc("/query", queryNameHandler)
	http.HandleFunc("/ua", ua)
	http.HandleFunc("/url/", urlPassParam)

	// file serve
	// fileServer := http.FileServer(http.Dir("./assets"))
	// http.Handle("/", fileServer)

	// process form input
	http.HandleFunc("/", process)

	// return the image
	http.HandleFunc("/image", imageHandler)

	log.Println("http server listening 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
