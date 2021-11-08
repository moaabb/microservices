package handlers

import (
	"compress/gzip"
	"net/http"
	"strings"
)

type GzipHandler struct {
}

type WrappedResponseWriter struct {
	rw http.ResponseWriter
	gw *gzip.Writer
}

func (g *GzipHandler) GzipMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			wrw := NewWrappedResponseWriter(rw)
			wrw.Header().Add("Content-Encoding", "gzip")
			defer wrw.Flush()

			next.ServeHTTP(wrw, r)
		} else {
			next.ServeHTTP(rw, r)
		}
	})
}

func NewWrappedResponseWriter(rw http.ResponseWriter) *WrappedResponseWriter {
	gw := gzip.NewWriter(rw)

	return &WrappedResponseWriter{
		rw: rw,
		gw: gw,
	}
}

func (wrw *WrappedResponseWriter) Write(d []byte) (int, error) {
	return wrw.gw.Write(d)
}

func (wrw *WrappedResponseWriter) Header() http.Header {
	return wrw.rw.Header()
}

func (wrw *WrappedResponseWriter) WriteHeader(statusCode int) {
	wrw.rw.WriteHeader(statusCode)
}

func (wrw *WrappedResponseWriter) Flush() {
	wrw.gw.Flush()
	wrw.gw.Close()
}
