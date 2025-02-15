package httpserver

import (
	httpresponse "gobackend/utils/httpResponse"
	"net/http"

	coreV1 "gobackend/internal/core"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type HttpServer struct {
	Router *mux.Router
	Logger *logrus.Entry
	Core   coreV1.ICore
}

func (h HttpServer) Init() error {
	h.registerRoutes()
	return nil
}
func (h HttpServer) registerRoutes() {
	h.Router.HandleFunc("/ping", h.Ping()).Methods("GET")
}
func (h HttpServer) Ping() func(rw http.ResponseWriter, rq *http.Request) {
	return func(rw http.ResponseWriter, rq *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				h.Logger.WithField("err", rec).Error("Panic recoverrd in Chat() server function")
				httpresponse.HTTPFailWith5xx(http.StatusText(500), http.StatusInternalServerError, rw, h.Logger)
			}
		}()
		httpresponse.HTTPSuccessWith200("/ping", rw, h.Logger)
	}
}

// func (h HttpServer) Chat() func(rw http.ResponseWriter, rq *http.Request) {
// 	return func(rw http.ResponseWriter, rq *http.Request) {
// 		defer func() {
// 			if rec := recover(); rec != nil {
// 				h.Logger.WithField("err", rec).Error("Panic recoverrd in Chat() server function")
// 				httpresponse.HTTPFailWith5xx(http.StatusText(500), http.StatusInternalServerError, rw, h.Logger)
// 			}
// 		}()
// 		h.Logger.Info("server Chat() enter")
// 		ctx := context.Background()
// 		chat := model.ChatRequest{}
// 		reqBody, err := io.ReadAll(rq.Body)
// 		if err != nil {
// 			h.Logger.Error(err)
// 			httpresponse.HTTPFailWith4xx(err.Error(), http.StatusBadRequest, rw, h.Logger)
// 			return
// 		}
// 		err = json.Unmarshal(reqBody, &chat)
// 		if err != nil {
// 			h.Logger.WithField("err", err).Error(err.Error())
// 			httpresponse.HTTPFailWith4xx(err.Error(), http.StatusBadRequest, rw, h.Logger)
// 			return
// 		}
// 		res, err := h.Core.Chat(ctx, chat)
// 		if err != nil {
// 			h.Logger.Error(err)
// 			httpresponse.HTTPFailWith5xx(err.Error(), http.StatusInternalServerError, rw, h.Logger)
// 			return
// 		}
// 		h.Logger.Info("server Chat() exit")
// 		httpresponse.HTTPSuccessWith200(res, rw, h.Logger)
// 	}
// }
