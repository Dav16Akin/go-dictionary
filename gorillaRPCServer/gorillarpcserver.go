package gorillarpcserver

import (
	"log"
	"net/http"
	"os"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	gjson "github.com/gorilla/rpc/json"
)

type Args struct {
	Id string
}

type Book struct {
	Id     string `json:"string,omitempty"`
	Name   string `json:"name,omitempty"`
	Author string `json:"author,omitempty"`
}

type JSONServer struct{}

func (t *JSONServer) GiveBookDetail(r *http.Request, args *Args, reply *Book) error {
	var books []Book

	data, err := os.ReadFile("./book.json")
	if err != nil {
		log.Println("Error reading file: ", err)
		return err
	}

	marshalerr := json.Unmarshal(data, &books)
	if marshalerr != nil {
		log.Println("Error unmarshaling :", err)
		return err
	}

	for _, book := range books {
		if book.Id == args.Id {
			*reply = book
            log.Printf("Found book: %s", book.Name)
			break
		}
	}

	return nil
}

func Run() {
	s := rpc.NewServer()
	s.RegisterCodec(gjson.NewCodec(), "application/json")

	s.RegisterService(new(JSONServer), "")
	r := mux.NewRouter()
	r.Handle("/rpc", s)

    log.Println("JSON-RPC Server listening on :1234")
	http.ListenAndServe(":1234", r)
}
