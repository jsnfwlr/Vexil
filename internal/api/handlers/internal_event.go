package handlers

import (
	"encoding/json"
	"io"
	"net/http"
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
