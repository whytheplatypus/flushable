# flushable [![GoDoc](https://godoc.org/github.com/whytheplatypus/flushable?status.svg)](http://godoc.org/github.com/whytheplatypus/flushable) [![Report card](https://goreportcard.com/badge/github.com/whytheplatypus/flushable)](https://goreportcard.com/report/github.com/whytheplatypus/flushable) [![Sourcegraph](https://sourcegraph.com/github.com/whytheplatypus/flushable/-/badge.svg)](https://sourcegraph.com/github.com/whytheplatypus/flushable?badge)

Package flushable provides a way to write to many http connections.

Motivated by a desire to passivly observe a programs logs.

Of course it could be used for other things too.

Typcally, observing logs requires significant effort.
Small simple projects require observability but the overhead
of integrating with the usual tools is disproportionate and the
tool itself is usually overkill.

Often the two systems need to have prior knowledge of eachother,
or at least the developer does.

Inspired by net/http/pprof/

```
import _ "github.com/whytheplatypus/flushable/log
```

Installs handlers for /debug/log that send log messages
while the connection is open.
e.g.

```
curl example.test/debug/log
```

Will show any logs created until the command is exited.
