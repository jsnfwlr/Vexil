package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jsnfwlr/o11y"

	"github.com/jsnfwlr/vexil/internal/api/handlers"
	"github.com/jsnfwlr/vexil/internal/api/oapi"
	"github.com/jsnfwlr/vexil/internal/log"
)

const (
	Address = "0.0.0.0"
	Port    = 9765
)

type Server struct {
	router *mux.Router
	server *http.Server
}

func errorHandler(w http.ResponseWriter, r *http.Request, err error) {
	var body []byte

	ctx := r.Context()
	ctx, o := o11y.Get(ctx, nil)

	e := StatusError{}

	if !errors.As(err, &e) {
		e = NewStatusError(ctx, http.StatusInternalServerError, err)
	} else {
		e = err.(StatusError)
	}

	body, _ = json.Marshal(e)

	o.Error(errors.New("error handling request"), log.RequestBodyKey, string(body))

	http.Error(w, string(body), e.Code)
}

func New(ctx context.Context, cfg Config) (server Server, fault error) {
	r := mux.NewRouter()

	mw := []mux.MiddlewareFunc{
		o11y.SetRequestID,
		// o11y.LogRequest,
	}

	r.Use(mw...)

	_, o := o11y.Get(ctx, nil)
	o.Debug("creating base router")

	h, err := handlers.New(ctx, cfg.DBClient, cfg.EnableSSE, "static", "index.html")
	if err != nil {
		return Server{}, err
	}

	api := oapi.NewStrictHandlerWithOptions(h, nil, oapi.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  errorHandler,
		ResponseErrorHandlerFunc: errorHandler,
	})

	if cfg.EnableSSE {
		o.Debug("enabling server-sent events")
		r.HandleFunc("/api/events", h.Events)
		r.HandleFunc("/api/event", h.PublishEvent)
	}
	r.Path("/").Name("ui-index").Methods(http.MethodGet).HandlerFunc(h.UI)
	r.PathPrefix("/js").Name("ui-js").Methods(http.MethodGet).HandlerFunc(h.UI)
	r.PathPrefix("/css").Name("ui-css").Methods(http.MethodGet).HandlerFunc(h.UI)
	r.Path("/favicon.ico").Name("ui-favicon").Methods(http.MethodGet).HandlerFunc(h.UI)

	oh := oapi.HandlerFromMux(api, r)

	core := &http.Server{
		Addr:                         fmt.Sprintf("%s:%d", Address, Port),
		DisableGeneralOptionsHandler: true,
		Handler:                      oh,
	}

	s := Server{
		server: core,
	}

	return s, nil
}

func (srvr *Server) Start(ctx context.Context) error {
	_, o := o11y.Get(ctx, nil)
	o.Info("starting server", "address", Address, "port", Port)
	err := srvr.server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
