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
        cd Twitter-Clone/

2. Instala dependencias:
    ```bash
        go mod tidy

3. Ejecuta la aplicación:
    ```bash
        go run cmd/api/main.go

## Notas

- **Inicio de base de datos:** La base de datos se inicializa automáticamente al iniciar la app.
- **Cambiar directorio de base de datos:** La base de datos se guarda en internal/data y para cambiar
su directorio se deben editar la constante "DBInitialize" y "DBPath" en la ruta "cmd/config/constants.go"