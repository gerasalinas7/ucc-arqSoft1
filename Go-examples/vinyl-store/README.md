Estructura del Proyecto:

/vinyl-store
    /albums
        albums.go
    /handlers
        handlers.go
    /models
        models.go
    main.go
    go.mod
    go.sum

Paso 1: Inicialización del Proyecto
Crear el directorio del proyecto:

mkdir vinyl-store
cd vinyl-store

Inicializar el módulo Go:
go mod init vinyl-store

Instalar Gin-Gonic:
go get github.com/gin-gonic/gin

Instalar el Paquete github.com/google/uuid
go get github.com/google/uuid


GET /albums: Devuelve todos los álbumes.
curl http://localhost:80/albums

GET /albums/:id: Devuelve un álbum por su ID.
curl http://localhost:80/albums/1

POST /albums: Agrega un nuevo álbum. (Envía datos JSON en el cuerpo de la solicitud).
curl -X POST http://localhost:80/albums -d '{"title": "Back in Black", "artist": "AC/DC", "year": 1980, "price": 19.99}' -H "Content-Type: application/json"

PUT /albums/:id: Actualiza un álbum existente.
curl -X PUT http://localhost:80/albums/1 -d '{"title": "The Dark Side of the Moon", "artist": "Pink Floyd", "year": 1973, "price": 20.50}' -H "Content-Type: application/json"

DELETE /albums/:id: Elimina un álbum por su ID.
curl -X DELETE http://localhost:80/albums/2