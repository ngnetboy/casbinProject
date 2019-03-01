package main

import (
	"controller"
	"model"
	"net/http"
	"os"
	"os/signal"
	"service"
	"syscall"

	log "github.com/Sirupsen/logrus"
)

func init() {
	model.LoadConf()
}

func signalHandler(ln *http.Server) {
	chSigInt := make(chan os.Signal, 1)
	signal.Notify(chSigInt, os.Signal(syscall.SIGINT))
	chSigTerm := make(chan os.Signal, 1)
	signal.Notify(chSigTerm, os.Signal(syscall.SIGTERM))

	log.Infoln("signal handler start")

	for {
		select {
		case <-chSigInt:
			log.Infoln("signal SIGINT")
			ln.Close()
			service.DisconnectDB()

		case <-chSigTerm:
			log.Infoln("signal SIGTERM")
			ln.Close()
			service.DisconnectDB()
		}
	}
}

func main() {
	log.Infoln("api server listen address ", model.Conf.Addr)

	service.ConnecDB()

	router := controller.MapRouters()

	server := &http.Server{
		Addr:    model.Conf.Addr,
		Handler: router,
	}
	go signalHandler(server)
	server.ListenAndServe()
}
