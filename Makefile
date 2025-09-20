# Makefile para gestionar el proyecto

# Variables
DOCKER_COMPOSE = docker-compose

# Comandos básicos
.PHONY: help
help:
	@echo "Comandos disponibles:"
	@echo "  build		Reconstruir los contenedores"
	@echo "  up		Iniciar los contenedores"
	@echo "  down		Detener y eliminar los contenedores"
	@echo "  restart	Reiniciar los contenedores"
	@echo "  logs		Ver los logs de los contenedores"
	@echo "  clean		Limpiar imágenes y volúmenes"
	@echo "  initdb	Inicializar la base de datos"

.PHONY: build
build:
	@echo "Reconstruyendo los contenedores..."
	$(DOCKER_COMPOSE) build --no-cache

.PHONY: up
up:
	@echo "Iniciando los contenedores..."
	$(DOCKER_COMPOSE) up -d

.PHONY: down
down:
	@echo "Deteniendo los contenedores..."
	$(DOCKER_COMPOSE) down

.PHONY: restart
restart:
	@echo "Reiniciando los contenedores..."
	$(DOCKER_COMPOSE) restart

.PHONY: logs
logs:
	@echo "Mostrando logs..."
	$(DOCKER_COMPOSE) logs -f

.PHONY: clean
clean:
	@echo "Limpiando imágenes y volúmenes..."
	$(DOCKER_COMPOSE) down -v
	docker image prune -f

.PHONY: initdb
initdb:
	@echo "Inicializando la base de datos en el contenedor..."
	$(DOCKER_COMPOSE) cp db/init.sh db:/docker-entrypoint-initdb.d/init.sh
	$(DOCKER_COMPOSE) exec db bash -c "su - postgres -c '/docker-entrypoint-initdb.d/init.sh'"
