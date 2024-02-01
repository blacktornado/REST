package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*****  BASIC GENERIC ERROR HANDLING ****/
type Errorr struct {
	Cause  error  `json:"-"`
	Detail string `json:"detail"`
	Status int    `json:"status"`
}

func (e *Errorr) Error(w http.ResponseWriter, r *http.Request) string {
	if e.Cause == nil {
		return e.Detail
	}
	return e.Detail + " : " + e.Cause.Error()
}

func NewHttpErr(err error, detail string, status int) *Errorr {
	return &Errorr{Cause: err, Detail: detail, Status: status}
}

/***** Basic Logger  ****/
type Logger struct {
	handler http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Executing middlewareOne")
	start := time.Now()
	l.handler.ServeHTTP(w, r)
	fmt.Println(r.Method, r.URL.Path, time.Since(start))
	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	fmt.Println("Executing middlewareOne again")

}

/**** NewLogger constructs a new Logger middleware handler ****/
func NewLoggerMiddleware(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

// func MiddlewareOne(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Executing middlewareOne")
// 		wrappedMux := NewLoggerMiddleware(next)
// 		fmt.Println(wrappedMux)
// 		next.ServeHTTP(w, r)
// 		fmt.Println("Executing middlewareOne again")
// 	})
// }
