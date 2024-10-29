// Paquete para configurar un logger con slog
package logg

import "log/slog"

// Logger es la instancia del logger.
var Logger *slog.Logger

// LogInfo contiene información sobre la solicitud HTTP para ser registrada.
type LogInfo struct {
	Method       string
	URL          string
	TimeStamp    string
	ResponseTime float64
	Status       int
}

// LogEntry representa una entrada de log que incluye detalles de la solicitud y el mensaje.
type LogEntry struct {
	Time    string
	Level   string
	Message string
	Request LogInfo
}

// LogValue convierte LogEntry en un formato adecuado para ser registrado por slog.
func (le LogEntry) LogValue() slog.Value {
	return slog.GroupValue(
		slog.String("time", le.Time),
		slog.String("level", le.Level),
		slog.String("message", le.Message),
		slog.Group("request",
			slog.String("method", le.Request.Method),
			slog.String("url", le.Request.URL),
			slog.String("timestamp", le.Request.TimeStamp),
			slog.Float64("responseTime", le.Request.ResponseTime),
			slog.Int("status", le.Request.Status),
		))
}
