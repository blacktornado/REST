package controller

import (
	"fmt"
	"net/http"
	controller "rest/controller/custaccount"
	"rest/middleware"
)

func Start() {
	var mux = http.NewServeMux()
	mux.HandleFunc("/api/getTopTrack", controller.GetTopTrack)
	mux.HandleFunc("/api/getTopTrackLyrics", controller.GetTopTrackL)
	mux.HandleFunc("/api/posts", controller.GetAllPosts)
	wrappedMux := middleware.NewLoggerMiddleware(mux)
	fmt.Printf("routerr initialized and listening on 3200\n")
	http.ListenAndServe(":3200", wrappedMux)

	// getPostHandler := http.HandlerFunc(controller.GetAllPosts)
	// mux.Handle("/api/posts", middlewareOne(getPostHandler))
	// fmt.Printf("routerr initialized and listening on 3200\n")
	// http.ListenAndServe(":3200", mux)

}
