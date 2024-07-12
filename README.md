# OlivePlay

#### Extending an Olive Branch to fun and friendly Community Play

## Overview

This project is a monorepo setup using Nx, Docker, Flutter, and Go. It contains a Flutter frontend and a Go backend using the go-chi framework. The project is structured to provide a consistent development environment and simplify deployment using Docker Compose.

## Project Structure

```
services/
├── api       # Go backend
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
└── frontend  # Flutter frontend
    ├── lib/
    ├── pubspec.yaml
    └── Dockerfile
workspace.json
docker-compose.yml
```

## Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Nx CLI](https://nx.dev/getting-started/intro)
- [Flutter SDK](https://flutter.dev/docs/get-started/install)
- [Go](https://golang.org/doc/install)

## Setup

### 1. Clone the Repository

```bash
git clone <repository-url>
cd <repository-directory>
```

### 2. Install Nx CLI

If you haven't already installed the Nx CLI, you can install it globally using npm:

```bash
npm install -g nx
```

### 3. Build and Run with Docker Compose

Build and run the services using Docker Compose:

```bash
docker-compose up --build
```

This command will build the Docker images for both the Flutter frontend and the Go backend, and start them.

### 4. Verify Setup

- **Flutter Frontend**: Open a web browser and navigate to `http://localhost`. You should see the Flutter web application.
- **Go Backend**: Open a web browser and navigate to `http://localhost:3000`. You should see "Hello, World!".

## Development

### Flutter Frontend

For local development of the Flutter frontend without Docker, you can use the following commands:

- **Serve**: Start the Flutter development server:

  ```bash
  cd services/frontend
  flutter run
  ```

- **Build**: Build the Flutter application:

  ```bash
  cd services/frontend
  flutter build web
  ```

### Go Backend

For local development of the Go backend without Docker, you can use the following commands:

- **Serve**: Start the Go server:

  ```bash
  cd services/api
  go run main.go
  ```

- **Build**: Build the Go application:

  ```bash
  cd services/api
  go build -o dist/main main.go
  ```

## Project Configuration

### Nx Configuration

The Nx configuration is located in `workspace.json` and individual `project.json` files for each service.

- **workspace.json**:

  ```json
  {
    "version": 2,
    "projects": {
      "frontend": "services/frontend",
      "api": "services/api"
    }
  }
  ```

- **services/frontend/project.json**:

  ```json
  {
    "name": "frontend",
    "projectType": "application",
    "sourceRoot": "services/frontend/lib",
    "targets": {
      "serve": {
        "executor": "nx:run-commands",
        "options": {
          "cwd": "services/frontend",
          "command": "flutter run"
        }
      },
      "build": {
        "executor": "nx:run-commands",
        "options": {
          "cwd": "services/frontend",
          "command": "flutter build"
        }
      }
    }
  }
  ```

- **services/api/project.json**:

  ```json
  {
    "name": "api",
    "projectType": "application",
    "sourceRoot": "services/api",
    "targets": {
      "serve": {
        "executor": "nx:run-commands",
        "options": {
          "cwd": "services/api",
          "command": "go run main.go"
        }
      },
      "build": {
        "executor": "nx:run-commands",
        "options": {
          "cwd": "services/api",
          "command": "go build -o dist/main main.go"
        }
      }
    }
  }
  ```

## Contributing

Please ensure you follow the guidelines below when contributing to this project:

1. Fork the repository.
2. Create a new branch with a descriptive name.
3. Make your changes.
4. Submit a pull request with a clear description of your changes.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contact

For any questions or suggestions, please contact:

- [Leo Belyi](mailto:leonid@ac93.org)
