package temple

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("localhost:9999", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(rw, "%s", req.RemoteAddr)
}
