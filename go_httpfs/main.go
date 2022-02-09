package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {
	wd, _ := os.Getwd()
	port := random_port()
	fmt.Println("Current directory : ", wd)
	fmt.Printf("Serving on http://localhost:%v\n", port)
	http.Handle("/", http.FileServer(http.Dir(".")))

	http.ListenAndServe(":"+strconv.Itoa(port), nil)
}

func random_port() int {
	rand.Seed(time.Now().UnixMilli())
	port := 1000 + rand.Intn(8999)
	return port
}
