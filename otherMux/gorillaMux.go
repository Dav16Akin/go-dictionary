package othermux

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category is: %v\n", vars["category"])
	fmt.Fprintf(w, "ID is: %v\n", vars["id"])
}

func RunGorilla(port string) {

	r := mux.NewRouter()

	r.HandleFunc("/articles/{category}/{id:[0-9]+}", ArticleHandler)

	r.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from the server")
	})


	address := port
    if string(port[0]) != ":" {
        address = ":" + port
    }

	srv := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Server is listening on PORT:8000")
	log.Fatal(srv.ListenAndServe())
}