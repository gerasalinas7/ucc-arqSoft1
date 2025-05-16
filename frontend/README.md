# React + Vite

This template provides a minimal setup to get React working in Vite with HMR and some ESLint rules.

Currently, two official plugins are available:

- [@vitejs/plugin-react](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react) uses [Babel](https://babeljs.io/) for Fast Refresh
- [@vitejs/plugin-react-swc](https://github.com/vitejs/vite-plugin-react/blob/main/packages/plugin-react-swc) uses [SWC](https://swc.rs/) for Fast Refresh

## Expanding the ESLint configuration

If you are developing a production application, we recommend using TypeScript with type-aware lint rules enabled. Check out the [TS template](https://github.com/vitejs/vite/tree/main/packages/create-vite/template-react-ts) for information on how to integrate TypeScript and [`typescript-eslint`](https://typescript-eslint.io) in your project.

---------PASOS

npm create vite@latest frontend --template react
cd frontend
npm install
(React y Javascript)


npm install react-router-dom
npm install axios


---- 
npm install

levantar API 

go get github.com/gin-contrib/cors

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

	// Habilitar CORS
	r.Use(cors.Default())

Para probar:

Inicia el backend ejecutando go run main.go (si no lo has hecho ya).

curl -X POST http://localhost:8080/register \
  -H "Content-Type: application/json" \
  -d '{"username": "admin", "password": "admin"}'


Inicia el front-end ejecutando npm run dev en el directorio del front-end.

Ve a http://localhost:3000 y prueba el formulario de login.

Usa las credenciales admin para iniciar sesión correctamente y ser redirigido a la página de "Hello World".

Si las credenciales son incorrectas, verás el mensaje "Login Incorrecto".

nvm use 18