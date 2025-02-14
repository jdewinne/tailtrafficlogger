package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/klauspost/compress/zstd"
)

// Handler for incoming HTTP requests
func handler(w http.ResponseWriter, r *http.Request) {
	// Check for "Content-Encoding: zstd"
	if r.Header.Get("Content-Encoding") == "zstd" {
		// Read the compressed body
		var compressedBody bytes.Buffer
		_, err := io.Copy(&compressedBody, r.Body)
		if err != nil {
			http.Error(w, "Failed to read request body", http.StatusInternalServerError)
			return
		}

		// Decompress using zstd
		decoder, err := zstd.NewReader(nil)
		if err != nil {
			http.Error(w, "Failed to create zstd decoder", http.StatusInternalServerError)
			return
		}
		defer decoder.Close()

		decompressedBody, err := decoder.DecodeAll(compressedBody.Bytes(), nil)
		if err != nil {
			http.Error(w, "Failed to decompress request body", http.StatusInternalServerError)
			return
		}

		// Log the decompressed body
		fmt.Printf("Received Decompressed Body: %s\n", decompressedBody)

		// Send an HTTP 200 OK response
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	} else {
		http.Error(w, "Unsupported Content-Encoding", http.StatusUnsupportedMediaType)
	}
}

func main() {
	// Start HTTP server on port 5678
	http.HandleFunc("/", handler)
	log.Println("Server listening on :5678...")
	log.Fatal(http.ListenAndServe(":5678", nil))
}
