package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/jsnfwlr/o11y"
	"github.com/jsnfwlr/vexil/internal/db"
	"github.com/jsnfwlr/vexil/internal/log"
	"github.com/r3labs/sse/v2"
)

func (s Server) SendEvent(ctx context.Context, environment string, flag Flag) error {
	_, o := o11y.Get(ctx, nil)

	if s.EventSrv == nil {
		return errors.New("server is not enabled for SSE")
	}

	data, err := json.Marshal(flag)
	if err != nil {
		return fmt.Errorf("could not convert flag to json string: %w", err)
	}

	event := sse.Event{
		Data: data,
	}

	ok := s.EventSrv.TryPublish(environment, &event)
	if !ok {
		o.Debug("PublishEvent failed", log.RequestIdKey, o11y.GetRequestID(ctx))
		return fmt.Errorf("could not publish %s flag %s to %s", flag.Type, flag.Name, environment)
	}
	o.Debug("PublishEvent succeeded", log.RequestIdKey, o11y.GetRequestID(ctx))

	return nil
}

type Flag struct {
	Name  string
	Value string
	Type  db.FlagType
	Env   string
}
