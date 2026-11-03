// Go Path can be hosted statically or served locally by this dependency-free Go server.
package main

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"
)

//go:embed dist/*
var assets embed.FS

func handler() http.Handler {
	public, err := fs.Sub(assets, "dist")
	if err != nil {
		panic(err)
	}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})
	mux.Handle("GET /", http.FileServer(http.FS(public)))
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
		w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self'; style-src 'self' 'unsafe-inline' https://fonts.googleapis.com; font-src https://fonts.gstatic.com; img-src 'self' data:; connect-src 'self'; object-src 'none'; base-uri 'self'; frame-ancestors 'none'")
		mux.ServeHTTP(w, r)
	})
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == "--sample-certificate" {
		if err := sampleCredential(os.Stdout); err != nil {
			log.Fatal(err)
		}
		return
	}
	addr := os.Getenv("GO_PATH_ADDR")
	if addr == "" {
		addr = "127.0.0.1:8080"
	}
	server := &http.Server{Addr: addr, Handler: academy(handler()), ReadHeaderTimeout: 5 * time.Second, ReadTimeout: 15 * time.Second, WriteTimeout: 15 * time.Second, IdleTimeout: 60 * time.Second}
	log.Printf("Go Path listening on %s", addr)
	log.Fatal(server.ListenAndServe())
}
