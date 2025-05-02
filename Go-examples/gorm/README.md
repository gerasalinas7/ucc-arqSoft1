Instalar MySQL:
brew install mysql

Iniciar MySQL:
brew services start mysql

Acceder a MySQL:
mysql -u root

1. CREATE DATABASE db_vinilos;

2. Seleccionar la base de datos db_vinilos
Después de crearla o si ya existe, selecciona la base de datos con el siguiente comando:

USE db_vinilos;

3. Crear la tabla albums
Ahora, si no existe la tabla albums, crea la tabla con el siguiente comando:

CREATE TABLE albums (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    year INT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);

4. Verificar que la tabla albums fue creada correctamente
Puedes verificar que la tabla albums haya sido creada correctamente con el siguiente comando:

SHOW TABLES;
Esto debería mostrarte una lista de tablas, incluida la tabla albums.

5. Insertar los registros
Ahora que la tabla albums existe, puedes insertar los tres registros que mencionaste con este comando:

INSERT INTO albums (title, artist, year, price)
VALUES
    ('The Number of the Beast', 'Iron Maiden', 1982, 25.19),
    ('Youthanasia', 'Megadeth', 1994, 13.65),
    ('Master of Puppets', 'Metallica', 1986, 20.97);

6. Verificar los registros
Finalmente, puedes verificar que los registros se insertaron correctamente con:

SELECT * FROM albums;

7. Salir de MySQL
Cuando hayas terminado, puedes salir de MySQL con:

EXIT;

--------------

Inicializar el módulo Go:
go mod init gorm

go get github.com/gin-gonic/gin
go get github.com/go-sql-driver/mysql
go get gorm.io/gorm
go get gorm.io/driver/mysql

**Go recomienda que el nombre del directorio esté en minúsculas.


-----
go run main.go

4. Probar los endpoints con Postman o curl
a. Probar GET /albums (Obtener todos los álbumes)
Usa Postman o curl para hacer una solicitud GET a http://localhost/albums.

Con curl:

curl -X GET http://localhost/albums
Deberías obtener una respuesta JSON con todos los álbumes en la base de datos.

b. Probar GET /albums/:id (Obtener un álbum por ID)
Para obtener un álbum específico, por ejemplo, el álbum con ID 1, puedes hacer una solicitud GET a http://localhost/albums/1.

Con curl:

curl -X GET http://localhost/albums/1
Esto debería devolverte el álbum con ese ID si existe.

c. Probar POST /albums (Agregar un nuevo álbum)
Para agregar un nuevo álbum, realiza una solicitud POST con un cuerpo JSON. Por ejemplo:

{
    "title": "The Number of the Beast",
    "artist": "Iron Maiden",
    "year": 1982,
    "price": 25.19
}
Con curl:

curl -X POST http://localhost/albums \
     -H "Content-Type: application/json" \
     -d '{"title": "The Number of the Beast", "artist": "Iron Maiden", "year": 1982, "price": 25.19}'
Esto debería agregar el nuevo álbum a la base de datos.

d. Probar PUT /albums/:id (Actualizar un álbum)
Para actualizar un álbum existente, realiza una solicitud PUT con el ID del álbum y los nuevos datos. Por ejemplo, para actualizar el álbum con ID 1:

{
    "title": "The Number of the Beast (Reissue)",
    "artist": "Iron Maiden",
    "year": 1982,
    "price": 29.99
}
Con curl:

curl -X PUT http://localhost/albums/1 \
     -H "Content-Type: application/json" \
     -d '{"title": "The Number of the Beast (Reissue)", "artist": "Iron Maiden", "year": 1982, "price": 29.99}'
Esto debería actualizar los datos del álbum con ID 1.

e. Probar DELETE /albums/:id (Eliminar un álbum)
Para eliminar un álbum, realiza una solicitud DELETE con el ID del álbum que deseas eliminar:

Con curl:
curl -X DELETE http://localhost/albums/1