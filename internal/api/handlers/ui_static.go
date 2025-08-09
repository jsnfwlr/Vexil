package handlers

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/jsnfwlr/o11y"
	otelTrace "go.opentelemetry.io/otel/trace"
)

// UI inspects the URL path to locate a file within the static dir
// on the SPA handler. If a file is found, it will be served. If not, the
// file located at the index path on the SPA handler will be served. This
// is suitable behavior for serving an SPA (single page application).
func (h Handlers) UI(w http.ResponseWriter, r *http.Request) {
	ctx, span := tracer.Start(r.Context(), "doUICreateSingleFlag", otelTrace.WithSpanKind(otelTrace.SpanKindServer))
	defer span.End()

	// Join internally call path.Clean to prevent directory traversal
	path := filepath.Join(h.staticPath, r.URL.Path)

	o := o11y.Get(ctx)

	o.Debug("static ui request", span, "file_path", r.URL.String())

	// check whether a file exists or is a directory at the given path
	fi, err := os.Stat(path)
	if err != nil {
		o.Error(err, span, "file_path", path)
		if os.IsNotExist(err) {
			w.WriteHeader(http.StatusNotFound)
			// file does not exist or path is a directory, serve index.html
			http.ServeFile(w, r, filepath.Join(h.staticPath, "404.html"))

			return
		}

		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	o.Debug("requested file", span, "is_dir", fi.IsDir())

	if fi.IsDir() {
		http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))

		return
	}

	// otherwise, use http.FileServer to serve the static file
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
}
