package main

import (
	"fmt"
	"net/http"

	"github.com/restBoard/api"
)

func main() {
	fmt.Println("Serve on :8000")
	http.ListenAndServe(":8000", api.NewRouter())
}

