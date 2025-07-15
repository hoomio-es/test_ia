# Property Manager

Una aplicación web para gestionar propiedades inmobiliarias construida con Go y TailwindCSS.

## Características

- Gestión completa de propiedades
- Interfaz moderna y responsive
- Base de datos PostgreSQL
- Formulario detallado para cada propiedad
- Gestión de información financiera y contratos

## Requisitos

- Docker y Docker Compose
- Go 1.21+
- PostgreSQL 15+

## Instalación

1. Clona el repositorio
2. Ejecuta los siguientes comandos:

```bash
docker-compose up -d --build
```

3. La aplicación estará disponible en http://localhost:8080

## Estructura de la Base de Datos

La base de datos PostgreSQL incluye una tabla `properties` con todos los campos necesarios para gestionar las propiedades, incluyendo:
- Información básica de la propiedad
- Datos de contacto
- Información financiera
- Contratos y seguros
- Notas adicionales

## Tecnologías Utilizadas

- Backend: Go (Fiber Framework)
- Frontend: HTML5, TailwindCSS
- Base de datos: PostgreSQL
- Contenedores: Docker/Docker Compose
