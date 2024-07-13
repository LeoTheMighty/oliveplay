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
- [Xcode](https://apps.apple.com/us/app/xcode/id497799835?mt=12) (for iOS development)
- A physical iPhone or iOS Simulator

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

### Testing Flutter on iOS

#### Enable Developer Mode on iPhone

1. Open the **Settings** app on your iPhone.
2. Go to **Privacy & Security**.
3. Scroll down and select **Developer Mode**.
4. Toggle the **Developer Mode** switch to enable it.
5. Your iPhone will prompt you to restart. After restarting, confirm that you want to enable Developer Mode.

#### Using Xcode and a Physical Device

1. Open a terminal and run `flutter doctor` to ensure all dependencies are installed.
2. Connect your iPhone to your Mac using a USB cable.
3. Navigate to your Flutter project directory: `cd services/frontend`.
4. Open the iOS project in Xcode: `open ios/Runner.xcworkspace`.
5. In Xcode, select your connected iPhone as the target device.
6. Click the "Run" button (or use the shortcut `Cmd + R`) to build and run the app on your iPhone.

#### Using iOS Simulator

1. Open the iOS simulator from Xcode: `Xcode > Open Developer Tool > Simulator`.
2. In the terminal, navigate to your Flutter project directory: `cd services/frontend`.
3. Run the app on the iOS simulator: `flutter run`.

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
