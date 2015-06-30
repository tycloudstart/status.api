package main

import (
	"flag"
	"fmt"
	"github.com/tycloudstart.com/status.api/g"
	"github.com/tycloudstart.com/status.api/http"
	"github.com/tycloudstart.com/status.api/proc"
	"os"
)

func main() {
	cfg := flag.String("c", "cfg.json", "config file")
	version := flag.Bool("v", false, "show version")
	versionGit := flag.Bool("vg", false, "show git version")
	flag.Parse()

	if *version {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
	if *versionGit {
		fmt.Println(g.VERSION, g.COMMIT)
		os.Exit(0)
	}

	// global config
	g.ParseConfig(*cfg)
	// proc
	proc.Start()

	// http
	http.Start()

	select {}
}
