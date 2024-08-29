## ABOUT

* Implementation of a minimal HTTP server that returns the request headers in the body of the HTTP response.
* Dependencies: Go language, or can also compiled into a static native executable.
* Plain HTTP and HTTPS version.
* Achieves thousands of req/sec with minimal CPU/memory usage.

### TO RUN

```
# HTTP version
go run server.go

# HTTPS version
go run server-tls.go
```
