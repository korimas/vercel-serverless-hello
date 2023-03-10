package handler

import (
	"fmt"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, time.Now().String())
	fmt.Fprintln(w, "Hello World!")
}
