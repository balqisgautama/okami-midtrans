package router

import (
	"github.com/balqisgautama/okami-midtrans/config"
	"github.com/balqisgautama/okami-midtrans/http/endpoint"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func ApiController(port int) {
	handler := mux.NewRouter()

	handler.HandleFunc(setPath("/health"), endpoint.HealthEndpoint.CheckingHealth).Methods("GET", "OPTIONS")

	handler.HandleFunc(setPath("/transaction"),
		endpoint.MidtransEndpoint.Transaction).Methods("POST", "OPTIONS")

	//util.Logger.Info("Hello World")
	//util.Logger.Error("Not able to reach blog.", zap.String("url", "localhost"))
	//util.Logger.Debug("logger debug", zap.String("debug", "try"))

	handler.Use(MiddlewareCORSOrigin)
	http.ListenAndServe(":"+strconv.Itoa(port), handler)
}

func setPath(path string) string {
	prefixPath := config.ApplicationConfiguration.GetServerPrefixPath()
	return "/" + prefixPath + path
}
