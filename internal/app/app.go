// El paquete `app` contiene la lógica para inicializar y ejecutar la aplicación completa,
// incluyendo la configuración de la base de datos, el router HTTP y el servidor.
package app

import (
	"PruebaGo/logg"
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// App define la estructura principal de la aplicación, incluyendo la conexión a la base de datos, el puerto y el router HTTP.
type App struct {
	DB     *pgxpool.Pool
	port   string
	router *http.ServeMux
}

// NewApp inicializa una nueva instancia de App, cargando la configuración del archivo .env y configurando la base de datos y el router.
func NewApp() (*App, error) {

	// Cargar las variables de entorno desde el archivo .env
	err := godotenv.Load()
	if err != nil {
		// Si ocurre un error al cargar el archivo .env, se detiene el programa y se imprime el error
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Obtener la DSN (cadena de conexión de la base de datos) y el puerto desde las variables de entorno
	dsn := os.Getenv("DSN")
	port := os.Getenv("PORT")

	// Configurar la base de datos utilizando la cadena DSN
	db, err := setupDatabase(dsn)
	if err != nil {
		// Si ocurre un error al configurar la base de datos, se devuelve un error
		return nil, fmt.Errorf("failed to setup database: %w", err)
	}

	// Crear un nuevo router para manejar las rutas HTTP
	router := http.NewServeMux()

	// Devolver una instancia de la aplicación con la base de datos configurada, el puerto y el router
	return &App{
		DB:     db,     // La conexión a la base de datos
		port:   port,   // El puerto en el que correrá el servidor
		router: router, // El enrutador HTTP
	}, nil
}

// Run inicia el servidor HTTP y configura todos los middlewares y servicios necesarios.
func (app *App) Run() error {
	// Configuración del logger para que use JSON en la salida estándar
	logg.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Inicializa el middleware de logging que envuelve el enrutador
	loggingMiddleware := logg.NewLoggingMiddleware(app.router)

	// Configuración de CORS para permitir peticiones desde cualquier origen
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPatch, http.MethodOptions, http.MethodPut},
		AllowedHeaders: []string{"Origin", "Content-Type", "Authorization", "Accept", "Success", "OK"},
	})

	// Aplica el middleware de CORS sobre el middleware de logging
	handler := c.Handler(loggingMiddleware)

	// Configuración del servidor HTTP
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", app.port),
		Handler: handler,
	}

	// Inicia el servicio principal de la aplicación (base de datos, enrutador, etc.)
	if err := UserService(context.Background(), app.DB, app.router); err != nil {
		return fmt.Errorf("failed to start user service: %w", err)
	}

	// Imprime en los logs que el servidor está iniciando en el puerto especificado
	log.Printf("starting server on port %s...", app.port)

	// Inicia el servidor HTTP y escucha peticiones
	return server.ListenAndServe()
}

// setupDatabase configura y devuelve un pool de conexiones a la base de datos usando pgxpool.
func setupDatabase(dsn string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %w", err)
	}

	return dbPool, nil
}
