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
2. Instala dependencias:
    ```bash
        go mod tidy

3. Ejecuta la aplicación:
    go run cmd/main.go

## Notas
La base de datos se guarda en internal/data y se inicializa automáticamente al iniciar la app.
Para configurar el tamaño máximo de los tweets o detalles de la base de datos, edita cmd/config.