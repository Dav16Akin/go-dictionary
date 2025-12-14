package othermux

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/julienschmidt/httprouter"
)

func getCommandOutput(command string, arguments ...string) (string, error) {
	// args... unpacks arguments array into elements
	cmd := exec.Command(command, arguments...)

	out, err := cmd.Output()

	if err != nil {
		return "", err
	}

	return string(out), nil
}

func goVersion(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	output, err := getCommandOutput("go", "version")
	if err != nil {
		http.Error(w, "Command Failed: "+err.Error(), 500)
		return
	}
	fmt.Fprintln(w, output)
}
func getFileContent(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	output, err := os.ReadFile("./latin.txt")
	if err != nil {
		http.Error(w, "File not found or unreadable", 404)
	}
	fmt.Fprintln(w, string(output))
}

func RunHttpRouter(port string) {
	router := httprouter.New()
	router.ServeFiles("/static/*filepath", http.Dir("/Users/apple/Documents/static"))
	// Mapping to methods is possible with HttpRouter
	router.GET("/api/v1/go-version", goVersion)
	// Path variable called name used here
	router.GET("/api/v1/show-file", getFileContent)

	fmt.Printf("Server running on port %s", port)
	log.Fatal(http.ListenAndServe(port, router))
}
