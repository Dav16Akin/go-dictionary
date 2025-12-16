package gorestfulfundamemtals

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func pingTime(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, time.Now().String())
}

func Run() {
	webservice := new(restful.WebService)

	webservice.Route(webservice.GET("/ping").To(pingTime))

	restful.Add(webservice)

	fmt.Println("Server is running on Port 8000")

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatalf("Error : %v", err)
	}
}
