package main

import (
	"bufio"
	"bytes"
	"log"
	"net/http"
)

func main() {
	http.Handle("/gaas", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		toFind := r.URL.Query().Get("q")
		if toFind == "" {
			w.WriteHeader(400)
			w.Write([]byte("need ?q= on query"))
			return
		}

		buf := bufio.NewScanner(r.Body)
		for buf.Scan() {
			if bytes.Contains(buf.Bytes(), []byte(toFind)) {
				w.Write(buf.Bytes())
				w.Write([]byte("\n"))
			}
		}
	}))

	log.Print("starting GaaS on :8000")
	http.ListenAndServe(":8000", nil)
}
