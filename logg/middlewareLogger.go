package logg

import (
	"net/http"
	"time"
)

// Logging implementa middleware para registrar solicitudes HTTP.
type Logging struct {
	Handler http.Handler
}

// NewLoggingMiddleware crea una nueva instancia de Logging.
func NewLoggingMiddleware(handler http.Handler) *Logging {
	return &Logging{Handler: handler}
}

// ServeHTTP registra la información de cada solicitud HTTP.
func (l *Logging) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()

	customWriter := &responseWriter{ResponseWriter: w, status: http.StatusOK}

	l.Handler.ServeHTTP(customWriter, r)

	duration := time.Since(start).Seconds()

	logEntry := LogEntry{
		Time:    time.Now().Format(time.RFC3339),
		Level:   "info",
		Message: "request completed",
		Request: LogInfo{
			Method:       r.Method,
			URL:          r.URL.String(),
			TimeStamp:    start.Format(time.RFC3339),
			ResponseTime: duration,
			Status:       customWriter.status,
		},
	}

	Logger.Info("request completed", "", logEntry.LogValue())
}

// responseWriter es un envoltorio alrededor de http.ResponseWriter para capturar el código de estado.
type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader sobrescribe el método WriteHeader para capturar el código de estado.
func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
