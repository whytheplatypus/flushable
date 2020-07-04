// Package log serves logs over http
//
// Inspired by net/http/pprof/
//
// 	import _ "github.com/whytheplatypus/flushable/log
//
// Installs handlers for /debug/log that send log messages
// while the connection is open.
// e.g.
// 	curl example.test/debug/log
//
// Will show any logs created until the command is exited.
package log

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/whytheplatypus/flushable"
)

func init() {
	m := &flushable.MultiFlusher{}
	log.SetOutput(io.MultiWriter(m, os.Stderr))
	http.Handle("/debug/log", m)
}
