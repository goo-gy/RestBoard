package main

import (
	"fmt"
	"github.com/restBoard/api"
	"net/http"
)

func main() {
	fmt.Println("Serve on :8000")
	http.ListenAndServe(":8000", api.NewRouter())
}
