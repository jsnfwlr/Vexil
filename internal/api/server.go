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
	var statusCode int
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

	o.Error(errors.New("error handling request"), log.RequestBodyKey, body)
	w.WriteHeader(e.Status())
	http.Error(w, string(body), statusCode)
}

func New(ctx context.Context, cfg Config) Server {
	r := mux.NewRouter()

	mw := []mux.MiddlewareFunc{
		o11y.SetRequestID,
		// o11y.LogRequest,
	}

	r.Use(mw...)

	_, o := o11y.Get(ctx, nil)
	o.Debug("creating base router")

	h := handlers.New(cfg.DBClient, cfg.EnableSSE)

	a := oapi.NewStrictHandlerWithOptions(h, nil, oapi.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  errorHandler,
		ResponseErrorHandlerFunc: errorHandler,
	})

	oapi.HandlerFromMux(a, r)

	m := http.NewServeMux()
	m.Handle("/", r)

	if cfg.EnableSSE {
		o.Debug("enabling server-sent events")
		m.HandleFunc("/api/events", h.Events)
		m.HandleFunc("/api/event", h.PublishEvent)
	}

	s := &http.Server{
		Addr:                         fmt.Sprintf("%s:%d", Address, Port),
		DisableGeneralOptionsHandler: true,
		Handler:                      m,
	}

	return Server{
		server: s,
	}
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
