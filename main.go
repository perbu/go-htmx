package main

import (
	"fmt"
	"github.com/perbu/go-htmx/poem"
	"github.com/perbu/go-htmx/static"
	"log"
	"net/http"
	"runtime"
	"time"
)

func main() {
	err := realMain()
	if err != nil {
		log.Fatal(err)
	}
}

func realMain() error {
	// set up a HTTP server:
	p := poem.NewPoem()

	router := http.NewServeMux()
	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/memstats", memstatsHandler)
	router.HandleFunc("/poem", p.NextHandler)
	router.HandleFunc("/createForm", createFormHandler)
	router.HandleFunc("/create", createHandler)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
	return nil
}

// indexHandler serves the index.html file, that is embedded in the binary, in the static package
func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("indexHandler")
	w.Header().Set("Content-Type", "text/html")
	w.Write(static.IndexHTML)
}

func memstatsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("memstatsHandler")
	var memstats runtime.MemStats
	runtime.ReadMemStats(&memstats)

	w.Header().Set("Content-Type", "text/html")
	// Print all the status we have collected:
	fmt.Fprintf(w, "Alloc: %d<br>", memstats.Alloc)
	fmt.Fprintf(w, "TotalAlloc: %d<br>", memstats.TotalAlloc)
	fmt.Fprintf(w, "Sys: %d<br>", memstats.Sys)
	fmt.Fprintf(w, "Lookups: %d<br>", memstats.Lookups)
	fmt.Fprintf(w, "Mallocs: %d<br>", memstats.Mallocs)
	fmt.Fprintf(w, "Frees: %d<br>", memstats.Frees)
	fmt.Fprintf(w, "NumGC: %d<br>", memstats.NumGC)
	// Number of goroutines:
	fmt.Fprintf(w, "NumGoroutine: %d<br>", runtime.NumGoroutine())
	// Print the current time:
	fmt.Fprintf(w, "Time: %s<br>", time.Now().Format(time.RFC3339))

}

func createHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("createHandler")
	r.ParseForm()
	name := r.Form.Get("name")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf(`<div>Created %s</div>`, name)))
}

func createFormHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("createFormHandler")
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(fmt.Sprintf(`<form 
id="theForm">
<input type="text" name="name">
<button value="Submit" hx-target="#theForm" hx-post="/create" hx-swap="outerHTML">Create</button>
</form>`)))
}
