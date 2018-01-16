package main

import (
	"os/exec"
	"./backend"
	"log"
	"flag"
	"runtime"
)

// openBrowser tries to open the URL in a browser,
// and returns whether it succeed in doing so.
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"google-chrome"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}

func main() {
	localPort := flag.String("l", "8066", "the local port of front-end")
	cacheFile := flag.String("c", ".cache", "the cache file")
	staticPath := flag.String("s", "static", "the static path of the front-end")
	flag.Parse()

	server := http.New(*cacheFile)
	server.StaticDir = *staticPath
	openBrowser("http://127.0.0.1:"+*localPort)
	log.Printf("Serving at http://127.0.0.1:"+*localPort)
	if err := server.ListenAndServe("0.0.0.0:"+*localPort); err != nil {
		log.Fatal(err)
	}

}