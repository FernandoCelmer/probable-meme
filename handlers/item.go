package handlers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func ItemHandler(w http.ResponseWriter, r *http.Request) {
	data := Response{Message: "Item"}
	response, _ := json.Marshal(data)
	w.Write(response)
}
