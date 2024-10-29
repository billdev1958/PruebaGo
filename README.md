# Documentación de PruebaGo

## Visión General
PruebaGo es un proyecto de servicio backend basado en Go que utiliza varias bibliotecas y herramientas clave para autenticación, gestión de entornos, registro de logs y operaciones con bases de datos. El proyecto utiliza Docker para la contenedorización, con PostgreSQL como base de datos. La versión de Go utilizada es la 1.23 e incluye un logger personalizado basado en `slog` para registros estructurados.

### Características Clave
- **Autenticación JWT** usando `github.com/golang-jwt/jwt/v5` para autenticación segura basada en tokens.
- **Gestión de UUID** mediante `github.com/google/uuid` para generar identificadores únicos.
- **Integración con PostgreSQL** a través de `github.com/jackc/pgx/v5` para la conectividad y operaciones con la base de datos.
- **Configuración de Entorno** utilizando `github.com/joho/godotenv` para gestionar variables de entorno.
- **Manejo de CORS** proporcionado por `github.com/rs/cors` para controlar solicitudes de origen cruzado.
- **Utilidades Criptográficas** de `golang.org/x/crypto` para encriptación y hashing seguros.
- **Logger Personalizado** utilizando `slog` de Go para proporcionar registros estructurados detallados y formateados.
- **PostgreSQL 16.4** contenedorizado usando Docker.

## Dependencias
Los siguientes módulos de Go se utilizan en este proyecto:

- `github.com/golang-jwt/jwt/v5 v5.2.1` - Proporciona soporte para JSON Web Token (JWT) para la gestión de tokens de autenticación.
- `github.com/google/uuid v1.6.0` - Genera identificadores únicos universales (UUID) para varias entidades.
- `github.com/jackc/pgx/v5 v5.7.1` - Controlador de PostgreSQL de alto rendimiento y conjunto de herramientas para operaciones con bases de datos.
- `github.com/joho/godotenv v1.5.1` - Carga variables de entorno desde un archivo `.env` para la gestión de configuraciones.
- `github.com/rs/cors v1.11.1` - Proporciona middleware para manejar el intercambio de recursos de origen cruzado (CORS) en solicitudes HTTP.
- `golang.org/x/crypto v0.27.0` - Paquetes criptográficos utilizados para hashing de contraseñas y otras funciones de seguridad.

## Logger Personalizado
PruebaGo incluye un logger personalizado basado en `slog`, que se utiliza para registrar información estructurada a lo largo de la aplicación. El logger captura detalles importantes como métodos de solicitud, tiempos de respuesta y códigos de estado, proporcionando una manera informativa de depurar y monitorear el servicio.

### Ejemplo de Registro
El logger personalizado puede capturar información como:
- Método HTTP utilizado (`GET`, `POST`, etc.).
- URL accedida.
- Marca de tiempo de la solicitud.
- Tiempo de respuesta en segundos.
- Código de estado de la respuesta HTTP.

Esto ayuda a identificar rápidamente cuellos de botella de rendimiento o problemas dentro de la aplicación.

## Entorno de Desarrollo
- **Versión de Go**: 1.23
- **Docker**: Docker se utiliza para contenedorización de la aplicación y los servicios, asegurando consistencia entre los entornos de desarrollo y producción.
- **PostgreSQL 16.4**: El proyecto utiliza un contenedor de PostgreSQL 16.4 para gestionar datos persistentes.

## Comandos del Makefile
El `Makefile` simplifica las tareas comunes de desarrollo. A continuación se detallan los comandos incluidos:

- **showdb**: Conecta al contenedor de PostgreSQL para inspeccionar registros de la base de datos.
  ```sh
  make showdb
  ```
  Este comando ejecuta una sesión interactiva dentro del contenedor de PostgreSQL usando `psql`.

- **godoc**: Inicia un servidor local para ver la documentación de Go del proyecto.
  ```sh
  make godoc
  ```
  Este comando inicia un servidor `godoc` en `http://localhost:6060`, donde puedes navegar a `/pkg/PruebaGo/` para ver la documentación del proyecto.

## Ejecución del Proyecto
### Requisitos Previos
Asegúrate de tener instalados los siguientes elementos en tu sistema:
- **Docker**: Para ejecutar el contenedor de PostgreSQL.
- **Go 1.23**: Para compilar y ejecutar la aplicación.

### Pasos para Ejecutar
1. **Iniciar el Contenedor de PostgreSQL**
   ```sh
   docker-compose up -d
   ```
   Este comando inicia el contenedor de la base de datos en segundo plano.

2. **Ejecutar la Aplicación**
   Puedes usar los comandos proporcionados en el `Makefile` para gestionar la aplicación.

3. **Acceder a la Documentación**
   Ejecuta `make godoc` para iniciar el servidor local `godoc` y ver la documentación del proyecto.

## Ejemplos de Solicitudes HTTP
### Registro de Usuario
- **URL**: `http://localhost:8080/v1/user`
- **Método**: `POST`
- **Cuerpo de la Solicitud**:
  ```json
  {
    "username": "john_doe",
    "phone": "1234567811",
    "email": "johndoe@example.com",
    "password": "StrongPa@123"
  }
  ```

### Inicio de Sesión
- **URL**: `http://localhost:8080/v1/login`
- **Método**: `POST`
- **Cuerpo de la Solicitud**:
  ```json
  {
    "identifier": "john_doe",
    "password": ""
  }
  ```

