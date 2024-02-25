package handlers

import (
	"fmt"
	"net/http"
)

func ClientesExtratoHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
