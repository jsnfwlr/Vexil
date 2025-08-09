package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/jsnfwlr/o11y"
	"github.com/jsnfwlr/vexil/internal/log"
	"github.com/r3labs/sse/v2"
	otelTrace "go.opentelemetry.io/otel/trace"
)

func (h Handlers) PublishEvent(w http.ResponseWriter, r *http.Request) {
	if h.EventSrv == nil {
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

	err = h.SendEvent(r.Context(), f.Env, f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`could send event`))
	}
}

func (h Handlers) Events(w http.ResponseWriter, r *http.Request) {
	if h.EventSrv == nil {
		return
	}

	ctx, span := tracer.Start(r.Context(), "Events", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	o := o11y.Get(ctx)

	o.Debug("Client connected to stream", span)

	go func() {
		// Received Browser Disconnection
		<-r.Context().Done()
	}()

	h.EventSrv.ServeHTTP(w, r)
}

func (h Handlers) SendEvent(ctx context.Context, environment string, flag Flag) error {
	ctx, span := tracer.Start(ctx, "SendEvent", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	o := o11y.Get(ctx)

	if h.EventSrv == nil {
		return errors.New("server is not enabled for SSE")
	}

	data, err := json.Marshal(flag)
	if err != nil {
		return fmt.Errorf("could not convert flag to json string: %w", err)
	}

	event := sse.Event{
		Data: data,
	}

	ok := h.EventSrv.TryPublish(environment, &event)
	if !ok {
		o.Debug("PublishEvent failed", span, log.RequestIdKey, o11y.GetRequestID(ctx))
		return fmt.Errorf("could not publish %s flag %s to %s", flag.Type, flag.Name, environment)
	}
	o.Debug("PublishEvent succeeded", span, log.RequestIdKey, o11y.GetRequestID(ctx))

	return nil
}
