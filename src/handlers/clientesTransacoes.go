package handlers

import (
	"fmt"
	"net/http"
)

func ClientTransacoesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}
