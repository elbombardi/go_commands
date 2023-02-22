package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

var port int
var ssl_cert_file string
var ssl_key_file string
var ssl bool = false
var directory string

func init() {
	flag.IntVar(&port, "port", random_port(), "http port")
	flag.StringVar(&ssl_cert_file, "cert_file", "", "SSL certification file")
	flag.StringVar(&ssl_key_file, "key_file", "", "SSL key file")
	flag.StringVar(&directory, "dir", ".", "Directory to serve from")
	flag.Parse()
	if ssl_cert_file != "" || ssl_key_file != "" {
		if ssl_cert_file == "" || ssl_key_file == "" {
			fmt.Fprintln(os.Stderr,
				"To launch SSL server, bother cert file and key file are mandatory :")
			flag.PrintDefaults()
			os.Exit(-1)
		}
		ssl = true
	}
}
func main() {
	wd := directory
	if wd == "." {
		wd, _ = os.Getwd()
	}
	log.Println("Current directory : ", wd)
	log.Printf("Serving on http://localhost:%v\n", port)
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		ts := time.Now()
		fs.ServeHTTP(rw, r)
		log.Printf("%v (%vms)", r.URL, time.Since(ts).Milliseconds())
	}))
	var err error
	if !ssl {
		err = http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), nil)
	} else {
		err = http.ListenAndServeTLS("0.0.0.0:"+strconv.Itoa(port), ssl_cert_file, ssl_key_file, nil)
	}
	panic(err)
}

func random_port() int {
	rand.Seed(time.Now().UnixMilli())
	port := 1000 + rand.Intn(8999)
	return port
}
