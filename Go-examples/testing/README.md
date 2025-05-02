go mod init math

Después de inicializar el módulo, descarga las dependencias necesarias para el proyecto ejecutando:
go mod tidy

go test
go test -v
go test -cover

go tool cover -html=coverage.out

