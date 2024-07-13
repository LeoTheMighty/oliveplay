# OlivePlay

#### Extending an Olive Branch to fun and friendly Community Play

## Overview

This project is a monorepo setup using Nx, Docker, Flutter, and Go. It contains a Flutter frontend (named `app`) and a Go backend using the go-chi framework. The project is structured to provide a consistent development environment and simplify deployment using Docker Compose.

## Project Structure

```
services/
├── api       # Go backend
│   ├── main.go
│   ├── go.mod
│   ├── go.sum
│   └── Dockerfile
└── app       # Flutter frontend
    ├── lib/
    ├── pubspec.yaml
    ├── .env.dev
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

### 3. Find Your Local IP Address

To connect your iOS device to your backend service running on your local machine, you need to know the local IP address of your development machine. Use the following command to get the IP address that starts with `192`:

#### On macOS and Linux:

```bash
ifconfig | grep "inet " | grep -Fv 127.0.0.1 | awk '{print $2}' | grep '^192'
```

#### On Windows:

1. Open **Command Prompt**.
2. Run the following command:

```bash
ipconfig
```

Look for the `IPv4 Address` under the section corresponding to your active network connection. Ensure it starts with `192`.

### 4. Create a `.env.dev` File

Copy over the template and fill out the necessary environment variables.

```bash
cp .env.dev.tmpl .env.dev
```

Looks like this:

```plaintext
API_URL=http://<your-local-ip>:3000
```

Replace `<your-local-ip>` with the IP address of your development machine.

### 5. Build and Run with Docker Compose

Build and run the services using Docker Compose:

```bash
docker-compose up
```

This command will build the Docker images for both the Flutter frontend and the Go backend, and start them.

### 6. Verify Setup

- **Flutter Frontend**: Open a web browser and navigate to `http://localhost`. You should see the Flutter web application.
- **Go Backend**: Open a web browser and navigate to `http://<your-local-ip>:3000`. You should see "Hello, World!".

## Development

### Flutter Frontend

For local development of the Flutter frontend without Docker, you can use the following commands:

- **Serve**: Start the Flutter development server:

  ```bash
  npx nx run app:serve
  ```

- **Build**: Build the Flutter application:

  ```bash
  npx nx run app:build
  ```
  
- **Install**: Install the flutter assets (Helpful for IDE setup)

    ```bash
fvm flutter upgrade
npx nx run app:install
    ```

### Go Backend

For local development of the Go backend without Docker, you can use the following commands:

- **Serve**: Start the Go server:

  ```bash
  npx nx run api:serve
  ```

- **Build**: Build the Go application:

  ```bash
  npx nx run api:build
  ```

### Testing Flutter on iOS

#### Enable Developer Mode on iPhone

1. Open the **Settings** app on your iPhone.
2. Go to **Privacy & Security**.
3. Scroll down and select **Developer Mode**.
4. Toggle the **Developer Mode** switch to enable it.
5. Your iPhone will prompt you to restart. After restarting, confirm that you want to enable Developer Mode.

#### Using Xcode and a Physical Device

1. **Find Local IP Address**:
  - Follow the steps in the "Find Your Local IP Address" section to get your local IP address.

2. **Update `.env.dev` File**:
  - Ensure the `.env.dev` file in your `services/app` directory contains the correct API URL pointing to your local IP address:
    ```plaintext
    API_URL=http://<your-local-ip>:3000
    ```

3. **Run Docker Compose**:
   ```bash
   docker-compose up --build
   ```
   
4. **Update Xcode Environment Variable**
   - Open your Flutter project in Xcode:
     ```bash
     cd services/app
     open ios/Runner.xcworkspace
     ```
  
   Then on the `Product > Scheme > Edit Scheme > Run > Arguments > Environment Variables` menu option, add this
   environment variable:

    ```plaintext
   API_URL=http://<your-local-ip>:3000
    ```

4. **Run the Flutter App in Xcode**:
  - Ensure your connected iPhone is selected as the target device.
  - Click the "Run" button (or use the shortcut `Cmd + R`) to build and run the app on your iPhone.

## Project Configuration

### Nx Configuration

The Nx configuration is located in `project.json` and individual `project.json` files for each service.

- **project.json**:

  ```json
  {
    "version": 2,
    "projects": {
      "app": "services/app",
      "api": "services/api"
    }
  }
  ```

- **services/app/project.json**:

  ```json
  {
    "name": "app",
    "projectType": "application",
    "sourceRoot": "services/app/lib",
    "targets": {
      "serve": {
        "executor": "nx:run-commands",
        "options": {
          "cwd": "services/app",
          "command": "flutter run"
        }
      },
      "build": {
        "executor": "nx:run-commands",
        "options": {
          "cwd": "services/app",
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
