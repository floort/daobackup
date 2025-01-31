package main

import (
	"io"
	"log"
	"net/http"
)

var StorageBackend = FileSystemStorage{
	Root:          "/Users/floort/daobackup/",
	VerifyOnRead:  true,
	VerifyOnWrite: true,
}

func chunkHandler(w http.ResponseWriter, r *http.Request) {
	chunkhash := r.PathValue("hash")
	switch r.Method {
	case "PUT":
		blob, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		err = StorageBackend.Put(chunkhash, blob)
		if err != nil {
			panic(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	case "HEAD":
		if StorageBackend.Exists(chunkhash) {
			w.WriteHeader(http.StatusOK)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	case "GET":
		if StorageBackend.Exists(chunkhash) {
			blob, err := StorageBackend.Get(chunkhash)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
			w.Write(blob)
			return
		} else {
			w.WriteHeader(http.StatusNotFound)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only PUT, HEAD and GET methods are allowed."))
		return
	}
}

func main() {
	http.HandleFunc("/chunk/{hash}", chunkHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
