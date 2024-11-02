# Twitter Clone API

Una API simple de clon de Twitter en Go, que permite a los usuarios registrarse, publicar tweets y seguir a otros usuarios.

## Funcionalidades

- **Usuarios:** Registro de nuevos usuarios.
- **Tweets:** Publicación de tweets y visualización de la línea de tiempo de usuarios seguidos.
- **Seguidores:** Seguir a otros usuarios y ver a quiénes se sigue.

## Instalación

1. Clona el repositorio:
   ```bash
   git clone https://github.com/tu_usuario/twitter-clone.git
   cd twitter-clone
Instala dependencias:
bash
Copiar código
go mod tidy
Uso
Ejecuta la aplicación:

bash
Copiar código
go run cmd/main.go
Accede a los endpoints con herramientas como curl o Postman.

Endpoints
Registrar usuario: POST /users

json
Copiar código
{ "username": "nombre_usuario" }
Publicar tweet: POST /tweets/:username

json
Copiar código
{ "content": "Contenido del tweet" }
Ver línea de tiempo: GET /timeline/:username

Seguir usuario: POST /follow/:username

json
Copiar código
{ "FollowingUsername": "nombre_usuario_a_seguir" }
Ver usuarios seguidos: GET /following/:username

Notas
La base de datos se guarda en internal/data y se inicializa automáticamente al iniciar la app.
Para configurar el tamaño máximo de los tweets o detalles de la base de datos, edita cmd/config.