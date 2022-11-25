// Copyright 2022 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Teonet fortune single-page application (SPA)
package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
)

const (
	appShort = "teofortune-onepage"
	appName  = "Teonet fortune single-page application"
	appVersion = "0.0.1"

	appPort = "8080"
)

//go:embed frontend/dist
var frontend embed.FS

func main() {

	// Logo
	fmt.Printf("%s ver. %s\n", appName, appVersion)

	// Get HTTP port from environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = appPort
	}

	// Parse parameters
	var addr string
	flag.StringVar(&addr, "addr", ":"+port, "http server local address")
	flag.Parse()

	// Static part of frontend 
	dist, err := fs.Sub(frontend, "frontend/dist")
	if err != nil {
		log.Fatalln(err)
	}
	frontendFS := http.FileServer(http.FS(dist))
	http.Handle("/", frontendFS)

	// Start HTTP server
	log.Printf("start listening for HTTP requests on %s", addr)
	log.Fatalln(http.ListenAndServe(addr, nil))
}
