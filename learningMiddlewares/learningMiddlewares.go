package learningmiddlewares

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

type City struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Area uint64 `json:"area"`
}

var (
	cities = make(map[int]City)
	index  = 1
	mutex  = &sync.Mutex{}
)

func ContentTypeMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		contentType := r.Header.Get("Content-Type")

		if contentType == "" || contentType != "application/json" {
			http.Error(w, "Content-Type must be application/json", http.StatusUnsupportedMediaType)
			return
		}

		h.ServeHTTP(w, r)
	})
}

func ServerTimeMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		cookie := http.Cookie{Name: "Server_Time_UTC",
			Value: strconv.FormatInt(time.Now().Unix(), 10), Path: "/"}

		http.SetCookie(w, &cookie)

		h.ServeHTTP(w, r)
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	switch r.Method {
	case "POST":
		var citiesData City

		if err := json.NewDecoder(r.Body).Decode(&citiesData); err != nil {
			http.Error(w, "Invalid Json", http.StatusBadRequest)
			return
		}

		mutex.Lock()

		citiesData.ID = index
		cities[index] = citiesData
		index++

		mutex.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(citiesData)
	default:
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
	}
}

func learningMiddlewares() {
	mainHandleLogic := http.HandlerFunc(mainLogic)

	// here we are chaining middlewares together
	http.Handle("/city", ContentTypeMiddleware(ServerTimeMiddleware(mainHandleLogic)))

	fmt.Println("Server running on PORT 8080....")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
