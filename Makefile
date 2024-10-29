# Makefile
CONTAINER_NAME=pruebago-db-1  
USER=root
DATABASE=prueba_go

# Función para obtener el ID del contenedor basado en el nombre
CONTAINER_ID=$(shell docker ps -qf "name=$(CONTAINER_NAME)")

# Entra a la base de datos para revisar registros
showdb:
	@docker exec -it $(CONTAINER_ID) psql -U $(USER) -d $(DATABASE)

godoc:
	@echo "Iniciando el servidor godoc. Puedes acceder a la documentación de PruebaGo en http://localhost:6060/pkg/PruebaGo/"
	@godoc -http=:6060