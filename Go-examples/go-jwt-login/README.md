go run main.go

Registrar usuario:
curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "user1", "password": "password123"}'

Login de usuario:
curl -X POST http://localhost:8080/login -H "Content-Type: application/json" -d '{"username": "user1", "password": "password123"}'
