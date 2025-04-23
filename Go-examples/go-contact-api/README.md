Paso 5: Probar los Endpoints
Obtener todos los contactos:
curl http://localhost:80/contacts

Obtener un contacto por ID:
curl http://localhost:80/contacts/1

Agregar un nuevo contacto:
curl -X POST http://localhost:80/contacts -d '{"id": "3", "name": "Alice", "email": "alice@example.com", "phone": "321-654-9870"}' -H "Content-Type: application/json"

Editar un contacto:
curl -X PUT http://localhost:80/contacts/1 -d '{"id": "1", "name": "John Doe Updated", "email": "john.updated@example.com", "phone": "123-456-7899"}' -H "Content-Type: application/json"

Eliminar un contacto:
curl -X DELETE http://localhost:80/contacts/2