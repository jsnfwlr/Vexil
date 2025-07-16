package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/jsnfwlr/o11y"
)

func (s Server) Events(w http.ResponseWriter, r *http.Request) {
	if s.EventSrv == nil {
		return
	}

	_, o := o11y.Get(r.Context(), nil)
	// ctx, o := o11y.Get(tracer.Start(r.Context(), "Events", trace.WithSpanKind(trace.SpanKindServer)))
	defer o.End()

	o.Debug("Client connected to stream")

	go func() {
		// Received Browser Disconnection
		<-r.Context().Done()
	}()

	s.EventSrv.ServeHTTP(w, r)
}

func (s Server) PublishEvent(w http.ResponseWriter, r *http.Request) {
	if s.EventSrv == nil {
		return
	}

	b, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`could not read body`))
	}

	var f Flag

	err = json.Unmarshal(b, &f)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`could not parse body`))
	}

	err = s.SendEvent(r.Context(), f.Env, f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`could send event`))
	}
}
