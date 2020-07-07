package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	defaultHttpPort = uint16(8080)
	defaultRootPath  = "../../www/"
)

func main() {

	var httpPort uint16 = defaultHttpPort

	var srvRootPath string = defaultRootPath
	if 2 <= len(os.Args) {
		srvRootPath = os.Args[1]
	}

	fmt.Printf("%s — this program serves cyber80 over HTTP on port %d", programName, httpPort)
	fmt.Print("\n\n")
	fmt.Printf("\t%s should now be available at http://localhost:%d/", programName, httpPort)
	fmt.Print("\n\n")
	fmt.Printf("\t%s is serving files from the directory: %q", programName, srvRootPath)
	fmt.Print("\n\n")
	fmt.Printf("\tYou can override the directory %s serves files from with a command similar to:", programName)
	fmt.Print("\n")
	fmt.Printf("\t\t%s /path/to/directory", programName)
	fmt.Print("\n\n")
	fmt.Printf("\tPress [CTRL]+[C] to exit.")

	var handler http.Handler
	{
		dir := http.Dir(srvRootPath)

		handler = http.FileServer(dir)
	}

	{
		addr := fmt.Sprintf(":%d", httpPort)
		http.ListenAndServe(addr, handler)
	}
}
