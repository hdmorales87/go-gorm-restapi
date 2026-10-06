# Go GORM REST API

REST API construida con Go, GORM, Gorilla Mux y PostgreSQL. Permite gestionar usuarios y tareas mediante operaciones CRUD.

## Requisitos Previos

- Go 1.27.1 o superior
- PostgreSQL (base de datos)
- Air (opcional, para hot reload)

## Instalación de Go

### Windows
1. Descarga el instalador desde [https://go.dev/dl/](https://go.dev/dl/)
2. Ejecuta el instalador y sigue las instrucciones
3. Verifica la instalación:
```bash
go version
```

### Linux/Mac
```bash
# Descargar e instalar
wget https://go.dev/dl/go1.27.1.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.27.1.linux-amd64.tar.gz

# Agregar al PATH (agregar a ~/.bashrc o ~/.zshrc)
export PATH=$PATH:/usr/local/go/bin

# Verificar instalación
go version
```

## Instalación de Dependencias

Clona el repositorio y navega al directorio del proyecto:

```bash
cd go-example
```

Instala las dependencias del proyecto:

```bash
go mod download
```

### Crear archivo .env

Crea un archivo `.env` en la raíz del proyecto con tus credenciales de base de datos. Puedes copiar el archivo de ejemplo:

```bash
# Windows
copy .env.example .env

# Linux/Mac
cp .env.example .env
```

Luego edita `.env` con tus credenciales reales:

```env
DB_HOST=localhost
DB_USER=someuser
DB_PASSWORD=mysecretpassword
DB_NAME=gorm
DB_PORT=5432
```

**Importante:** El archivo `.env` ya está en `.gitignore` por seguridad, así que tus credenciales no se compartirán en el repositorio.

Las dependencias principales son:
- `github.com/gorilla/mux` - Router HTTP
- `gorm.io/gorm` - ORM para Go
- `gorm.io/driver/postgres` - Driver de PostgreSQL para GORM
- `github.com/joho/godotenv` - Carga de variables de entorno desde archivo .env
- `github.com/swaggo/swag` - Herramienta para generar documentación Swagger
- `github.com/swaggo/files` - Servidor de archivos estáticos para Swagger UI
- `github.com/swaggo/http-swagger` - Middleware HTTP para Swagger

## Configuración de la Base de Datos

El proyecto usa PostgreSQL. La configuración de la base de datos se maneja mediante variables de entorno.

### Configuración mediante archivo .env

1. Copia el archivo de ejemplo:
```bash
cp .env.example .env
```

2. Edita el archivo `.env` con tus credenciales de PostgreSQL:
```env
DB_HOST=localhost
DB_USER=someuser
DB_PASSWORD=mysecretpassword
DB_NAME=gorm
DB_PORT=5432
```

3. El archivo `.env` está en `.gitignore` por seguridad, así que no se compartirán tus credenciales en el repositorio.

### Configuración mediante variables de entorno del sistema

Alternativamente, puedes configurar las variables de entorno directamente en tu sistema:

**Windows (PowerShell):**
```powershell
$env:DB_HOST="localhost"
$env:DB_USER="someuser"
$env:DB_PASSWORD="mysecretpassword"
$env:DB_NAME="gorm"
$env:DB_PORT="5432"
```

**Windows (CMD):**
```cmd
set DB_HOST=localhost
set DB_USER=someuser
set DB_PASSWORD=mysecretpassword
set DB_NAME=gorm
set DB_PORT=5432
```

**Linux/Mac (bash/zsh):**
```bash
export DB_HOST=localhost
export DB_USER=someuser
export DB_PASSWORD=mysecretpassword
export DB_NAME=gorm
export DB_PORT=5432
```

**Nota:** El archivo `.env` tiene prioridad sobre las variables de entorno del sistema.

## Documentación Swagger (API Documentation)

Este proyecto incluye Swagger/OpenAPI para documentar automáticamente la API. Swagger UI proporciona una interfaz interactiva para probar los endpoints.

### Instalación de Swagger

#### 1. Instalar el CLI de swag

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Asegúrate de que el directorio `go/bin` esté en tu PATH. En Windows, agrega `%USERPROFILE%\go\bin` a tu PATH.

#### 2. Instalar las dependencias de Swagger

```bash
go get -u github.com/swaggo/swag/cmd/swag
go get -u github.com/swaggo/files
go get -u github.com/swaggo/http-swagger
```

### Generar la Documentación

Cada vez que añadas o modifique anotaciones Swagger en el código, debes regenerar la documentación:

```bash
swag init
```

Este comando crea la carpeta `docs/` con los archivos:
- `docs/docs.go` - Código Go generado
- `docs/swagger.json` - Especificación OpenAPI en JSON
- `docs/swagger.yaml` - Especificación OpenAPI en YAML

**Nota:** La carpeta `docs/` está en `.gitignore` ya que se genera automáticamente.

### Acceder a Swagger UI

Una vez que el servidor esté corriendo, accede a la documentación Swagger en:

```
http://localhost:8080/swagger/index.html
```

Desde ahí puedes:
- Ver todos los endpoints disponibles
- Ver los modelos de datos (DTOs)
- Probar los endpoints directamente desde la interfaz
- Ver ejemplos de requests y responses

### Anotaciones Swagger

La documentación se genera a partir de anotaciones en el código:

- **main.go**: Información general de la API (título, versión, descripción, host)
- **routes/*.go**: Documentación de cada endpoint (summary, description, parámetros, responses)

Ejemplo de anotación:
```go
// CreateUserHandler godoc
// @Summary Create a new user
// @Description Create a new user with the provided data
// @Tags users
// @Accept  json
// @Produce  json
// @Param user body dto.CreateUserDTO true "User data"
// @Success 200 {object} dto.UserDTO
// @Failure 500 {object} map[string]string
// @Router /users [post]
```

### Flujo de Trabajo Recomendado

1. Añade o modifica anotaciones Swagger en los handlers
2. Ejecuta `swag init` para regenerar la documentación
3. Reinicia el servidor
4. Accede a `http://localhost:8080/swagger/index.html` para ver los cambios

## Cómo Levantar el Proyecto

### Opción 1: Ejecución Normal

Compila y ejecuta el proyecto:

```bash
go run main.go
```

El servidor estará disponible en `http://localhost:8080`

### Opción 2: Compilar y Ejecutar

```bash
go build -o main.exe .
./main.exe
```

## ¿Qué es Air?

**Air** es una herramienta de "live reload" para aplicaciones Go. Detecta cambios en los archivos y automáticamente recompila y reinicia la aplicación, lo que hace el desarrollo más rápido y eficiente sin tener que detener y reiniciar manualmente el servidor.

### Instalación de Air

```bash
go install github.com/cosmtrek/air@latest
```

Asegúrate de que el directorio `go/bin` esté en tu PATH.

### Uso de Air

Para ejecutar el proyecto con hot reload:

```bash
air
```

Air leerá la configuración del archivo `.air.toml` y:
- Observará los cambios en los archivos `.go`
- Recompilará automáticamente cuando detecte cambios
- Reiniciará el servidor con los nuevos cambios

### Configuración de Air (.air.toml)

El archivo `.air.toml` contiene la configuración de Air:

- **[build]**: Configuración de compilación
  - `bin`: Ubicación del binario compilado
  - `cmd`: Comando de compilación
  - `include_ext`: Extensiones de archivos a observar (`.go`, `.tpl`, `.html`)
  - `exclude_dir`: Directorios a ignorar (`tmp`, `vendor`, `testdata`)
  - `delay`: Tiempo de espera antes de recompilar (ms)

- **[build.windows]**: Configuración específica para Windows
  - `bin`: `tmp\main.exe`
  - `cmd`: `go build -o ./tmp/main.exe .`

- **[log]**: Configuración de logs
- **[color]**: Colores de la salida en consola

## Estructura del Proyecto

```
go-example/
├── main.go              # Punto de entrada de la aplicación
├── go.mod               # Módulos y dependencias de Go
├── go.sum               # Checksums de dependencias
├── .air.toml            # Configuración de Air (hot reload)
├── .gitignore           # Archivos ignorados por Git
├── .env.example         # Ejemplo de variables de entorno (debe copiarse a .env)
├── .env                 # Variables de entorno (creado por el usuario, en .gitignore)
├── db/
│   └── connection.go    # Conexión a la base de datos
├── models/
│   ├── User.go          # Modelo de Usuario (base de datos)
│   └── Task.go          # Modelo de Tarea (base de datos)
├── dto/
│   ├── user_dto.go      # DTOs de Usuario (API)
│   └── task_dto.go      # DTOs de Tarea (API)
├── mappers/
│   ├── user_mapper.go   # Conversión Model ↔ DTO (Usuario)
│   └── task_mapper.go   # Conversión Model ↔ DTO (Tarea)
├── services/
│   ├── user_service.go  # Lógica de negocio de usuarios
│   └── task_service.go  # Lógica de negocio de tareas
├── routes/
│   ├── index.routes.go  # Rutas principales
│   ├── users.routes.go  # Rutas de usuarios (controladores)
│   └── tasks.routes.go  # Rutas de tareas (controladores)
├── router/
│   └── router.go        # Configuración de rutas e inyección de dependencias
├── docs/                # Documentación Swagger (generada por swag init)
│   ├── docs.go          # Código Go generado
│   ├── swagger.json     # Especificación OpenAPI JSON
│   └── swagger.yaml     # Especificación OpenAPI YAML
└── tmp/                 # Directorio temporal (compilados)
```

### Arquitectura en Capas

El proyecto sigue una arquitectura en capas para mantener el código desacoplado y mantenible:

- **Models**: Definición de las estructuras de datos y esquemas de base de datos (GORM)
- **DTOs (Data Transfer Objects)**: Estructuras para entrada/salida de la API, desacopladas de los modelos de BD
- **Mappers**: Funciones para convertir entre Models y DTOs
- **Services**: Contienen la lógica de negocio y operaciones con la base de datos (CRUD), trabajan con Models y retornan DTOs
- **Routes (Controllers)**: Manejan las solicitudes HTTP, reciben DTOs del request, invocan los servicios y retornan DTOs en la respuesta
- **Router**: Configura todas las rutas e inyecta las dependencias (servicios con la base de datos)
- **DB**: Configuración y conexión a la base de datos

### Flujo de Datos y Transformación

1. **Request HTTP** → Llega al controller con JSON
2. **Controller** → Decodifica JSON a DTO (CreateDTO/UpdateDTO)
3. **Controller** → Llama al service con el DTO
4. **Service** → Usa Mapper para convertir DTO → Model
5. **Service** → Realiza operaciones en BD con el Model
6. **Service** → Usa Mapper para convertir Model → DTO
7. **Service** → Retorna el DTO al controller
8. **Controller** → Codifica DTO a JSON y lo envía en la respuesta

### Flujo de Inyección de Dependencias

La "magia" de Go permite una inyección de dependencias limpia y explícita:

1. **main.go**: Conecta a la base de datos y la pasa al router
2. **router**: Recibe la base de datos, crea los servicios (`UserService`, `TaskService`) con la DB inyectada, y configura las rutas pasando los servicios a los controladores
3. **routes (controllers)**: Reciben los servicios ya inicializados y los usan para manejar las solicitudes
4. **services**: Tienen acceso a la base de datos a través de la inyección recibida

Este enfoque hace que el código sea:
- **Testeable**: Fácil mockear dependencias en tests
- **Modular**: Cada componente tiene responsabilidades claras
- **Explícito**: Las dependencias son visibles en los parámetros de las funciones
- **Seguro**: Los modelos de BD no se exponen directamente en la API

## Endpoints de la API

### Usuarios
- `GET /users` - Obtener todos los usuarios
- `GET /users/{id}` - Obtener un usuario por ID
- `POST /users` - Crear un nuevo usuario
- `PUT /users/{id}` - Actualizar un usuario
- `DELETE /users/{id}` - Eliminar un usuario

### Tareas
- `GET /tasks` - Obtener todas las tareas
- `GET /tasks/{id}` - Obtener una tarea por ID
- `POST /tasks` - Crear una nueva tarea
- `PUT /tasks/{id}` - Actualizar una tarea
- `DELETE /tasks/{id}` - Eliminar una tarea

### Home
- `GET /` - Página de inicio

## Comandos Útiles

```bash
# Instalar dependencias
go mod download

# Actualizar dependencias
go mod tidy

# Ejecutar tests
go test ./...

# Formatear código
go fmt ./...

# Ejecutar con Air (hot reload)
air

# Ejecutar normalmente
go run main.go

# Compilar
go build -o main.exe .

# Generar documentación Swagger
swag init
```

## Licencia

Este proyecto es de ejemplo educativo.
