package api

import (
	"fmt"
	"net/http"
)

// RestAPIStartServer - creates server with the specified routs and context path
func RestAPIStartServer(host string, port string, contextPath string) *http.Server {
	log.Info("api.RestAPIStartServer", constants.LogBeginFunc)
	router := httprouter.New()
	// all routes
	allPOSTRoutes := map[string]httprouter.Handle{
		contextPath + "/content": handleAddContent,
	}
	allGETRoutes := map[string]httprouter.Handle{
		contextPath + "/ping": heartBeat,
	}
	for endPoint, handler := range allPOSTRoutes {
		router.POST(endPoint, AddRequestID(handler))
	}
	for endPoint, handler := range allGETRoutes {
		router.GET(endPoint, AddRequestID(handler))
	}

	sock := fmt.Sprintf("%s:%s", host, port)
	srv := &http.Server{
		Addr:    sock,
		Handler: router,
	}
	go func(httpServ *http.Server, addr string) {
		httpServ.RegisterOnShutdown(func() {
			log.Info("api.RestAPIStartServer", "Shutting down http server")
		})
		stlog.Fatal(httpServ.ListenAndServe())
	}(srv, sock)
	log.Info("api.RestAPIStartServer", constants.LogFinishFunc)
	return srv
}
