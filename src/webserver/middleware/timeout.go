package middleware

import (
	"net/http"
	"context"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx := r.Context()
	newCtx, _ := context.WithTimeout(ctx, 3 * time.Second)
	r.WithContext(newCtx)
	ch := make(chan struct{})
	go func() {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()
	select {
	case <-ch:
		return
	case <-newCtx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
}