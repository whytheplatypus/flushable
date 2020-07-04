package flushable_test

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/whytheplatypus/flushable"
)

func ExampleMultiFlusher_ServeHTTP() {
	srv := http.Server{
		Addr: ":8080",
	}
	defer srv.Shutdown(context.Background())
	m := &flushable.MultiFlusher{}
	log.SetOutput(io.MultiWriter(m, os.Stderr))
	log.SetFlags(0)
	http.Handle("/debug/log", m)

	go srv.ListenAndServe()

	go func() {
		time.Sleep(10 * time.Millisecond)
		log.Println("hello world")
	}()

	res, err := http.Get("http://localhost:8080/debug/log")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	r := bufio.NewReader(res.Body)
	result, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	// Output: hello world
}
