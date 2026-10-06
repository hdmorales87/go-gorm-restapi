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

### Crear archivo .env.docker (para Docker Compose)

Para usar Docker Compose, crea un archivo `.env.docker` en la raíz del proyecto. Puedes copiar el archivo de ejemplo:

```bash
# Windows
copy .env.docker.example .env.docker

# Linux/Mac
cp .env.docker.example .env.docker
```

Este archivo contiene las variables de entorno específicas para los contenedores Docker:

```env
# Docker PostgreSQL Configuration
POSTGRES_USER=someuser
POSTGRES_PASSWORD=mysecretpassword
POSTGRES_DB=gorm

# pgAdmin Configuration
PGADMIN_DEFAULT_EMAIL=admin@admin.com
PGADMIN_DEFAULT_PASSWORD=admin
PGADMIN_PORT=5050

# Port mapping for PostgreSQL
DB_PORT=5432
```

**Importante:** El archivo `.env.docker` también está en `.gitignore` por seguridad.

Las dependencias principales son:
- `github.com/gorilla/mux` - Router HTTP
- `gorm.io/gorm` - ORM para Go
- `gorm.io/driver/postgres` - Driver de PostgreSQL para GORM
- `github.com/joho/godotenv` - Carga de variables de entorno desde archivo .env
- `github.com/swaggo/swag` - Herramienta para generar documentación Swagger
- `github.com/swaggo/files` - Servidor de archivos estáticos para Swagger UI
- `github.com/swaggo/http-swagger` - Middleware HTTP para Swagger

## Configuración de la Base de Datos

El proyecto usa PostgreSQL. La configuración de la base de datos se maneja mediante variables de entorno separadas para la aplicación Go y para Docker Compose.

### Archivos de Configuración

El proyecto utiliza dos archivos de configuración de variables de entorno:

1. **`.env`** - Para la aplicación Go (configuración de conexión a la base de datos)
2. **`.env.docker`** - Para Docker Compose (configuración de los contenedores PostgreSQL y pgAdmin)

Esta separación permite tener configuraciones diferentes para desarrollo local (con Docker) y producción (con base de datos externa).

### Configuración para la Aplicación Go (.env)

1. Copia el archivo de ejemplo:
```bash
cp .env.example .env
```

2. Edita el archivo `.env` con tus credenciales de PostgreSQL:
```env
# Database Configuration (for Go application)
DB_HOST=localhost
DB_USER=someuser
DB_PASSWORD=mysecretpassword
DB_NAME=gorm
DB_PORT=5432
```

3. El archivo `.env` está en `.gitignore` por seguridad, así que no se compartirán tus credenciales en el repositorio.

### Configuración para Docker Compose (.env.docker)

1. Copia el archivo de ejemplo:
```bash
cp .env.docker.example .env.docker
```

2. Edita el archivo `.env.docker` con las credenciales para los contenedores:
```env
# Docker PostgreSQL Configuration
POSTGRES_USER=someuser
POSTGRES_PASSWORD=mysecretpassword
POSTGRES_DB=gorm

# pgAdmin Configuration
PGADMIN_DEFAULT_EMAIL=admin@admin.com
PGADMIN_DEFAULT_PASSWORD=admin
PGADMIN_PORT=5050

# Port mapping for PostgreSQL
DB_PORT=5432
```

3. El archivo `.env.docker` también está en `.gitignore` por seguridad.

**Nota:** Las credenciales de Docker Compose están aisladas en el archivo `.env.docker` en lugar de estar harcodeadas en `docker-compose.yml`. Esto mejora la seguridad y permite cambiar las credenciales fácilmente sin modificar el archivo de configuración de Docker.

### Configuración mediante variables de entorno del sistema

Alternativamente, puedes configurar las variables de entorno directamente en tu sistema. Para la aplicación Go:

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

Para Docker Compose, se recomienda usar el archivo `.env.docker` en lugar de variables de entorno del sistema, ya que `docker-compose.yml` está configurado para leer de ese archivo específicamente.

## Docker Compose

El proyecto incluye una configuración de Docker Compose para levantar PostgreSQL y pgAdmin fácilmente.

### Requisitos Previos

Antes de levantar los contenedores, asegúrate de crear el archivo `.env.docker`:

```bash
# Windows
copy .env.docker.example .env.docker

# Linux/Mac
cp .env.docker.example .env.docker
```

Edita el archivo `.env.docker` con tus credenciales deseadas.

### Levantar los contenedores

```bash
docker-compose --env-file .env.docker up -d
```

Esto iniciará:
- **PostgreSQL** en el puerto 5432 (configurable via `DB_PORT` en `.env.docker`)
- **pgAdmin** en el puerto 5050 (configurable via `PGADMIN_PORT` en `.env.docker`)

**Nota:** El flag `--env-file .env.docker` es necesario porque docker-compose por defecto solo busca archivos llamados `.env`.

### Detener los contenedores

```bash
docker-compose --env-file .env.docker down
```

### Ver logs

```bash
docker-compose --env-file .env.docker logs -f
```

### Credenciales en Docker Compose

Las credenciales de PostgreSQL y pgAdmin están aisladas en el archivo `.env.docker` en lugar de estar harcodeadas en `docker-compose.yml`. El archivo `docker-compose.yml` referencia estas variables usando la sintaxis `${VARIABLE}`.

Variables requeridas en `.env.docker` para Docker Compose:
- `POSTGRES_USER`: Usuario de PostgreSQL
- `POSTGRES_PASSWORD`: Contraseña de PostgreSQL
- `POSTGRES_DB`: Nombre de la base de datos
- `PGADMIN_DEFAULT_EMAIL`: Email para pgAdmin
- `PGADMIN_DEFAULT_PASSWORD`: Contraseña para pgAdmin
- `PGADMIN_PORT`: Puerto para pgAdmin (por defecto 5050)
- `DB_PORT`: Puerto para PostgreSQL (por defecto 5432)

### Acceder a pgAdmin

Una vez que los contenedores estén corriendo, accede a pgAdmin en:
```
http://localhost:5050
```

Credenciales por defecto (configurables en `.env`):
- Email: `admin@admin.com`
- Contraseña: `admin`

Para conectar a PostgreSQL desde pgAdmin:
- Host: `postgres` (nombre del servicio en docker-compose)
- Port: `5432`
- Database: `gorm` (o el valor de `POSTGRES_DB`)
- Username: `someuser` (o el valor de `POSTGRES_USER`)
- Password: `mysecretpassword` (o el valor de `POSTGRES_PASSWORD`)

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
├── main.go                  # Punto de entrada de la aplicación
├── go.mod                   # Módulos y dependencias de Go
├── go.sum                   # Checksums de dependencias
├── .air.toml                # Configuración de Air (hot reload)
├── .gitignore               # Archivos ignorados por Git
├── .env.example             # Ejemplo de variables de entorno para la app Go (debe copiarse a .env)
├── .env                     # Variables de entorno para la app Go (creado por el usuario, en .gitignore)
├── .env.docker.example      # Ejemplo de variables de entorno para Docker (debe copiarse a .env.docker)
├── .env.docker              # Variables de entorno para Docker (creado por el usuario, en .gitignore)
├── docker-compose.yml       # Configuración de Docker Compose (PostgreSQL + pgAdmin)
├── db/
│   └── connection.go        # Conexión a la base de datos
├── models/
│   ├── User.go              # Modelo de Usuario (base de datos)
│   └── Task.go              # Modelo de Tarea (base de datos)
├── dto/
│   ├── user_dto.go          # DTOs de Usuario (API)
│   └── task_dto.go          # DTOs de Tarea (API)
├── mappers/
│   ├── user_mapper.go       # Conversión Model ↔ DTO (Usuario)
│   └── task_mapper.go       # Conversión Model ↔ DTO (Tarea)
├── repositories/
│   ├── interfaces.go        # Interfaces de repositorios (ITaskRepository, IUserRepository)
│   ├── task_repository.go   # Implementación de repositorio de tareas con GORM
│   └── user_repository.go   # Implementación de repositorio de usuarios con GORM
├── services/
│   ├── interfaces.go        # Interfaces de servicios (ITaskService, IUserService)
│   ├── user_service.go      # Lógica de negocio de usuarios
│   └── task_service.go      # Lógica de negocio de tareas
├── routes/
│   ├── index.routes.go      # Rutas principales
│   ├── users.routes.go      # Rutas de usuarios (controladores)
│   └── tasks.routes.go      # Rutas de tareas (controladores)
├── router/
│   └── router.go            # Configuración de rutas e inyección de dependencias
├── docs/                    # Documentación Swagger (generada por swag init)
│   ├── docs.go              # Código Go generado
│   ├── swagger.json         # Especificación OpenAPI JSON
│   └── swagger.yaml         # Especificación OpenAPI YAML
└── tmp/                     # Directorio temporal (compilados)
```

### Arquitectura en Capas

El proyecto sigue una arquitectura en capas (Layered Architecture) con principios SOLID para mantener el código desacoplado y mantenible:

- **Models**: Definición de las estructuras de datos y esquemas de base de datos (GORM)
- **DTOs (Data Transfer Objects)**: Estructuras para entrada/salida de la API, desacopladas de los modelos de BD
- **Mappers**: Funciones para convertir entre Models y DTOs
- **Repositories**: Capa de acceso a datos, encapsula la lógica de interacción con la base de datos. Se abstraen mediante interfaces para permitir cambio de tecnología de BD sin afectar las capas superiores.
- **Services**: Contienen la lógica de negocio, trabajan con repositorios y retornan DTOs. Se abstraen mediante interfaces para facilitar testing y desacoplamiento.
- **Routes (Controllers)**: Manejan las solicitudes HTTP, reciben DTOs del request, invocan los servicios (vía interfaces) y retornan DTOs en la respuesta
- **Router**: Configura todas las rutas e inyecta las dependencias (repos → servicios → controladores)
- **DB**: Configuración y conexión a la base de datos

### Flujo de Datos y Transformación

1. **Request HTTP** → Llega al controller con JSON
2. **Controller** → Decodifica JSON a DTO (CreateDTO/UpdateDTO)
3. **Controller** → Llama al service (vía interfaz) con el DTO
4. **Service** → Usa Mapper para convertir DTO → Model
5. **Service** → Llama al repositorio (vía interfaz) con el Model
6. **Repository** → Realiza operaciones en BD con GORM
7. **Repository** → Retorna el Model al service
8. **Service** → Usa Mapper para convertir Model → DTO
9. **Service** → Retorna el DTO al controller
10. **Controller** → Codifica DTO a JSON y lo envía en la respuesta

### Flujo de Inyección de Dependencias

La inyección de dependencias sigue una cadena explícita de arriba hacia abajo:

1. **main.go**: Conecta a la base de datos y la pasa al router
2. **router**:
   - Crea los repositorios (`TaskRepository`, `UserRepository`) con la DB inyectada
   - Crea los servicios (`TaskService`, `UserService`) inyectando los repositorios correspondientes
   - Configura las rutas pasando los servicios (como interfaces) a los controladores
3. **routes (controllers)**: Reciben los servicios ya inicializados (como interfaces) y los usan para manejar las solicitudes
4. **services**: Tienen acceso a los repositorios (como interfaces) para realizar operaciones de datos
5. **repositories**: Tienen acceso directo a la base de datos a través de GORM

Este enfoque hace que el código sea:
- **Testeable**: Fácil mockear repositorios para testear servicios sin base de datos, y mockear servicios para testear controladores
- **Modular**: Cada componente tiene responsabilidades claras y está desacoplado
- **Explícito**: Las dependencias son visibles en los parámetros de las funciones
- **Seguro**: Los modelos de BD no se exponen directamente en la API
- **Flexible**: Es posible cambiar la tecnología de base de datos o ORM modificando solo la capa de repositorios

### Cambios Técnicos Recientes

#### Implementación de Patrón Repository e Interfaces

Se ha refactorizado la arquitectura para seguir mejores prácticas de diseño en Go, implementando:

**1. Capa de Repositorios (Nueva)**
- Se creó el directorio `repositories/` con:
  - `interfaces.go`: Define las interfaces `ITaskRepository` y `IUserRepository` con métodos CRUD
  - `task_repository.go`: Implementación concreta usando GORM para operaciones de tareas
  - `user_repository.go`: Implementación concreta usando GORM para operaciones de usuarios

**2. Interfaces de Servicios**
- Se creó `services/interfaces.go` con:
  - `ITaskService`: Define todos los métodos del servicio de tareas
  - `IUserService`: Define todos los métodos del servicio de usuarios

**3. Refactorización de Servicios**
- `services/task_service.go`: Ya no depende directamente de GORM. Ahora recibe `ITaskRepository` por inyección de dependencias.
- `services/user_service.go`: Ya no depende directamente de GORM. Ahora recibe `IUserRepository` por inyección de dependencias.

**4. Actualización de Controladores**
- `routes/tasks.routes.go`: Todos los handlers ahora reciben `ITaskService` en lugar de `*TaskService`
- `routes/users.routes.go`: Todos los handlers ahora reciben `IUserService` en lugar de `*UserService`

**5. Actualización del Router**
- `router/router.go`: Se modificó el flujo de inyección de dependencias:
  ```go
  // Antes: Servicios recibían la base de datos directamente
  userService := services.NewUserService(database)
  taskService := services.NewTaskService(database)

  // Ahora: Repositorios reciben la DB, servicios reciben repositorios
  taskRepository := repositories.NewTaskRepository(database)
  userRepository := repositories.NewUserRepository(database)
  var taskService services.ITaskService = services.NewTaskService(taskRepository)
  var userService services.IUserService = services.NewUserService(userRepository)
  ```

**Beneficios de los Cambios:**

- **Principio de Inversión de Dependencias (DIP)**: Las capas superiores no dependen de implementaciones concretas, sino de abstracciones (interfaces)
- **Testing mejorado**: Posibilidad de crear mocks de repositorios para testear servicios sin necesidad de una base de datos real
- **Cambio de tecnología**: Si se desea cambiar GORM por otro ORM (como SQLBoiler, sqlx, etc.), solo se modifican las implementaciones de los repositorios sin afectar servicios o controladores
- **Separación de responsabilidades**: Los repositorios manejan acceso a datos, los servicios manejan lógica de negocio, los controladores manejan HTTP
- **Desacoplamiento**: Cada capa es independiente y puede evolucionar por separado

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

# Docker Compose - Levantar contenedores
docker-compose --env-file .env.docker up -d

# Docker Compose - Detener contenedores
docker-compose --env-file .env.docker down

# Docker Compose - Ver logs
docker-compose --env-file .env.docker logs -f

# Docker Compose - Ver estado de contenedores
docker-compose --env-file .env.docker ps
```

## Licencia

Este proyecto es de ejemplo educativo.
