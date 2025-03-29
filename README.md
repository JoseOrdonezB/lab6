# La Liga Tracker ⚽

Aplicación backend escrita en Go para gestionar partidos de fútbol. Permite crear, consultar, actualizar y eliminar partidos, así como registrar eventos clave como goles, tarjetas y tiempo extra.

## 🛠️ Tecnologías

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [SQLite](https://www.sqlite.org/)
- [Swagger](https://swagger.io/) para documentación
- Docker

## 🚀 Endpoints principales

| Método | Endpoint                        | Descripción                         |
|--------|----------------------------------|-------------------------------------|
| GET    | `/api/matches`                  | Obtener todos los partidos          |
| GET    | `/api/matches/:id`              | Obtener un partido por ID           |
| POST   | `/api/matches`                  | Crear un nuevo partido              |
| PUT    | `/api/matches/:id`              | Actualizar un partido existente     |
| DELETE | `/api/matches/:id`              | Eliminar un partido                 |
| PATCH  | `/api/matches/:id/goals`        | Registrar un gol                    |
| PATCH  | `/api/matches/:id/yellowcards`  | Registrar una tarjeta amarilla      |
| PATCH  | `/api/matches/:id/redcards`     | Registrar una tarjeta roja          |
| PATCH  | `/api/matches/:id/extratime`    | Registrar tiempo extra              |

## 📄 Documentación de la API
La documentación en formato JSON de la API se puede encontrar en el sigueinte directorio:
lab6/docs/swagger.json

## Vistas de frontend
Frontend funcionando:
![Actualizar Partido](./assets/actualizar-partido.png)
![Buscar Partido](./assets/buscar-partido.png)
![Crear Partido](./assets/crear-partido.png)
![Listado Partidos](./assets/listado-partidos.png)

## 🐳 Docker

Para ejecutar la app con Docker:

```bash
docker build -t laliga-tracker .
docker run -p 8080:8080 laliga-tracker