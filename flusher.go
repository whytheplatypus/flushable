// Package flushable provides a way to write to many http connections.
//
// Motivated by a desire to passivly observe a programs logs.
//
// Of course it could be used for other things too.
//
// Typcally, observing logs requires significant effort.
// Small simple projects require observability but the overhead
// of integrating with the usual tools is disproportionate and the
// tool itself is usually overkill.
//
// Often the two systems need to have prior knowledge of eachother,
// or at least the developer does.
//
// Inspired by net/http/pprof/
//
// 	import _ "github.com/whytheplatypus/flushable/log"
//
// Installs handlers for /debug/log that send log messages
// while the connection is open.
// e.g.
// 	curl example.test/debug/log
//
// Will show any logs created until the command is exited.
package flushable

import "net/http"

type writeFlusher interface {
	http.ResponseWriter
	http.Flusher
}

// A MultiFlusher provides the ability to write to multiple
// http.ResponseWriter instances at once, flushing to their
// clients on each write.
type MultiFlusher struct {
	writers map[writeFlusher]bool
}

func (m *MultiFlusher) Write(p []byte) (n int, err error) {
	for w := range m.writers {
		w.Write(p)
		w.Flush()
	}
	// this is, of course, non-sense
	return len(p), nil
}

// Include a standard http.ResponseWriter to be written to.
func (m *MultiFlusher) Include(w http.ResponseWriter) {
	if m.writers == nil {
		m.writers = map[writeFlusher]bool{}
	}
	m.writers[w.(writeFlusher)] = true
}

// Remove a standard http.ResponseWriter from the
func (m *MultiFlusher) Remove(w http.ResponseWriter) {
	if m.writers == nil {
		m.writers = map[writeFlusher]bool{}
	}
	delete(m.writers, w.(writeFlusher))
}

// ServeHTTP allows MultiFlusher to be used as an http.Handler.
// The http.ResponseWriter is added to the list of writers, and removed
// when the client closes the request.
func (m *MultiFlusher) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Include(w)
	defer m.Remove(w)
	<-r.Context().Done()
}
