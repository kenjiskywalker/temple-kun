package temple-kun

import (
	"fmt"
	"net/http"
	"strings"
)

func init() {
	http.HandleFunc("/", handler)
	//	http.ListenAndServe("localhost:9999", nil)
}

func handler(rw http.ResponseWriter, req *http.Request) {

	pos := strings.LastIndex(req.RemoteAddr, ":")
	if pos > 0 {
		addr := strings.Split(req.RemoteAddr, ":")
		// fmt.Fprintf(rw, "YOUR IP ADDRESS %s", addr[0])
		fmt.Fprintf(rw, addr[0])
	}

}
