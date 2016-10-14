package handlers

import (
	"net/http"
	"fmt"
)

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Todo Index!")
}
